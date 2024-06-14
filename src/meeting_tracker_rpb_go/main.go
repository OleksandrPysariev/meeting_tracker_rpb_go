package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"sync"

	"github.com/d2r2/go-logger"

	"github.com/OleksandrPysariev/meeting_tracker_rpb_go/set_event"
	"github.com/OleksandrPysariev/meeting_tracker_rpb_go/read_event"
)

var wg sync.WaitGroup

func main() {
	logger.ChangePackageLogLevel("i2c", logger.InfoLevel)
	fmt.Print("Starting main\n")
	wg.Add(1)
	go set_event.RunSetEvent()
	fmt.Print("Ran RunSetEvent\n")
	go read_event.RunReadEvent()
	fmt.Print("Ran RunReadEvent\n")
	wg.Wait()
}

func init() {
	fmt.Print("Application start...\n")
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-c
        wg.Done()
		fmt.Print("\nApplication shut down.\n")
        os.Exit(1)
    }()
}