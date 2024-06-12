package read_event

import (
	"fmt"
	"time"

	"github.com/OleksandrPysariev/meeting_tracker_rpb_go/read_event/utils"
)

var BACKLIGHT = true
var ON = true

func RunReadEvent() {
	fmt.Print("Running RunReadEvent...")
	currentlyInTheMeeting := false
	event := read_event.ParseEvent()
	lastCalled := read_event.TimeNow()
	for {
		if !ON {
			time.Sleep(1 * time.Second)
			continue
		}
		now := read_event.TimeNow()
		// Monitor when the meeting starts to get the new meeting when it ends
		if !currentlyInTheMeeting && event.Start.Before(now) && event.End.After(now) {
			currentlyInTheMeeting = true
			// todo: play melody
		}
		// Get new meeting to show if you just finished a meeting
		if currentlyInTheMeeting && event.End.Before(now) {
			currentlyInTheMeeting = false
			event = read_event.ParseEvent()
			lastCalled = read_event.TimeNow()
		}
		// Refresh current meeting every 2 minutes to track new meetings throughout the day
		if now.Sub(lastCalled).Seconds() > 120 {
			event = read_event.ParseEvent()
			lastCalled = read_event.TimeNow()
		}
		if !read_event.DateEqual(event.Start, read_event.TimeNow()) {
			line1 := "    "
			line1 += read_event.TimeNow().Format("15:04:05")
			line1 += "    "

			line2 := ""
			read_event.Say(line1, line2)
		} else {
			line1 := read_event.TimeNow().Format("15:04:05")
			line1 += " "
			line1 += event.Time
			line2 := event.Description
			read_event.Say(line1, line2)
		}
		time.Sleep(10 * time.Millisecond)
	}
	// TODO: implement LCD commands
	// https://github.com/d2r2/go-hd44780
	// fmt.Printf("Received a message: %#v", event)
}