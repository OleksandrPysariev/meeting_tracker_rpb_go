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
	fmt.Print("[main] Starting main\n")
	wg.Add(1)
	go set_event.RunSetEvent()
	fmt.Print("[main] Ran RunSetEvent\n")
	go read_event.RunReadEvent()
	fmt.Print("[main] Ran RunReadEvent\n")
	wg.Wait()
}

func init() {
	fmt.Print("[init] Application start...\n")
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-c
        wg.Done()
		fmt.Print("\n[init] Application shut down.\n")
        os.Exit(1)
    }()
}