package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

const (
	serviceAccountFile = "service.json"
	workCalendarID     = "oleksandr.pysariev@vimmi.net"
	personalCalendarID = "apalexlife@gmail.com"
)

var calendarService *calendar.Service

type Event struct {
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
	Time        string    `json:"time"`
	Description string    `json:"description"`
}

func InitCalendarService() error {
	ctx := context.Background()
	srv, err := calendar.NewService(ctx,
		option.WithCredentialsFile(serviceAccountFile),
		option.WithScopes(calendar.CalendarScope),
	)
	if err != nil {
		return fmt.Errorf("unable to create calendar service: %v", err)
	}
	calendarService = srv
	return nil
}

func fetchEvents(calendarID string, origin string) ([]*calendar.Event, error) {
	now := time.Now().UTC()
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 0, 0, time.UTC)

	events, err := calendarService.Events.List(calendarID).
		TimeMin(now.Format(time.RFC3339)).
		TimeMax(endOfDay.Format(time.RFC3339)).
		SingleEvents(true).
		OrderBy("startTime").
		Do()
	if err != nil {
		return nil, fmt.Errorf("unable to fetch events from %s: %v", calendarID, err)
	}

	var filtered []*calendar.Event
	for _, event := range events.Items {
		// Filter out whole-day events (they have Date instead of DateTime)
		if event.Start.DateTime == "" || event.End.DateTime == "" {
			continue
		}
		if event.Summary == "" {
			event.Summary = capitalize(origin) + " event"
		}
		filtered = append(filtered, event)
	}
	return filtered, nil
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func getEventTime(dateTimeStr string) (time.Time, error) {
	return time.Parse(time.RFC3339, dateTimeStr)
}

func getEarliestEvent(workEvents, personalEvents []*calendar.Event) *calendar.Event {
	if len(workEvents) == 0 && len(personalEvents) == 0 {
		return nil
	}
	if len(workEvents) == 0 {
		return personalEvents[0]
	}
	if len(personalEvents) == 0 {
		return workEvents[0]
	}

	workStart, _ := getEventTime(workEvents[0].Start.DateTime)
	personalStart, _ := getEventTime(personalEvents[0].Start.DateTime)

	if workStart.Before(personalStart) || workStart.Equal(personalStart) {
		return workEvents[0]
	}
	return personalEvents[0]
}

func shortenedSummary(input string) string {
	input = strings.ReplaceAll(input, "Palyanytsya", "")
	input = strings.ReplaceAll(input, "Palyanysta", "")
	input = strings.ReplaceAll(input, "Connect", "")
	input = strings.ReplaceAll(input, "Product", "")
	input = strings.ReplaceAll(input, "-", "")
	input = strings.ReplaceAll(input, "/", " ")
	input = strings.TrimSpace(input)

	if len(input) <= 13 {
		return input
	}

	words := strings.Fields(input)
	shortened := make([]string, len(words))
	for i, word := range words {
		if len(word) <= 3 {
			shortened[i] = word
		} else {
			shortened[i] = word[:3]
		}
	}
	return strings.Join(shortened, " ")
}

func GetCurrentEvent() ([]byte, error) {
	if calendarService == nil {
		return nil, fmt.Errorf("calendar service not initialized")
	}

	workEvents, err := fetchEvents(workCalendarID, "work")
	if err != nil {
		fmt.Printf("[google_calendar] error fetching work events: %s\n", err)
		workEvents = nil
	}

	personalEvents, err := fetchEvents(personalCalendarID, "personal")
	if err != nil {
		fmt.Printf("[google_calendar] error fetching personal events: %s\n", err)
		personalEvents = nil
	}

	event := getEarliestEvent(workEvents, personalEvents)

	if event == nil {
		fmt.Println("[google_calendar] No upcoming events found.")
		now := time.Now()
		emptyEvent := Event{
			Start:       now,
			End:         now,
			Time:        "",
			Description: "",
		}
		return json.Marshal(emptyEvent)
	}

	start, _ := getEventTime(event.Start.DateTime)
	end, _ := getEventTime(event.End.DateTime)
	duration := int(end.Sub(start).Minutes())

	result := Event{
		Start:       start,
		End:         end,
		Time:        start.Format("15:04"),
		Description: fmt.Sprintf("%d %s", duration, shortenedSummary(event.Summary)),
	}

	return json.Marshal(result)
}
