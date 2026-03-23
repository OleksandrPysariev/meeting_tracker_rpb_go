package set_event

import (
	"fmt"
	"time"

	"github.com/OleksandrPysariev/meeting_tracker_rpb_go/set_event/utils"
)

func setEvent() {
	bodyBytes, err := utils.GetCurrentEvent()
	if err != nil {
		fmt.Printf("[set_event] error getting event: %s\n", err)
		return
	}
	utils.WriteBodyToFile(bodyBytes)
}

func RunSetEvent() {
	fmt.Print("[set_event] Starting setEvent...\n")
	if err := utils.InitCalendarService(); err != nil {
		fmt.Printf("[set_event] Failed to initialize Google Calendar service: %s\n", err)
		return
	}
	fmt.Print("[set_event] Google Calendar service initialized.\n")
	for {
		setEvent()
		time.Sleep(60 * time.Second)
	}
}
