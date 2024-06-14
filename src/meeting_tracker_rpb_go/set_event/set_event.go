package set_event

import (
	"time"
	"fmt"

	"github.com/OleksandrPysariev/meeting_tracker_rpb_go/set_event/utils"
)

func setEvent() {
	// Yield response body
	bodyBytes := utils.ReadResponseBody()
	// Write body to event.json
	utils.WriteBodyToFile(bodyBytes)
}

func RunSetEvent() {
	fmt.Print("[set_event] Starting setEvent...\n")
	for {
		setEvent()
		fmt.Print("[set_event] event.json stored!\n")
		time.Sleep(60 * time.Second)
	}
}