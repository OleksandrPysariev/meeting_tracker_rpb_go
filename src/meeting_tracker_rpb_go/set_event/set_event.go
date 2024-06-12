package set_event

import (
	"time"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/OleksandrPysariev/meeting_tracker_rpb_go/set_event/utils"
)

func setEvent() {
	// Yield response body
	bodyBytes := utils.ReadResponseBody()
	// Write body to event.json
	utils.WriteBodyToFile(bodyBytes)
}

func RunSetEvent() {
	for {
		setEvent()
		time.Sleep(60 * time.Second)
	}
}

func init() {
	fmt.Print("Application start...\n")
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-c
        // Run Cleanup
		fmt.Print("\nApplication shut down.\n")
        os.Exit(1)
    }()
}