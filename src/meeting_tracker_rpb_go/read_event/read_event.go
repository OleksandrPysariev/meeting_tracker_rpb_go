package read_event

import (
	"fmt"
	"time"
	"strings"

	device "github.com/d2r2/go-hd44780"
	"github.com/d2r2/go-i2c"
	"github.com/d2r2/go-logger"
	"github.com/stianeikeland/go-rpio/v4"

	"github.com/OleksandrPysariev/meeting_tracker_rpb_go/buzzer"
	utils "github.com/OleksandrPysariev/meeting_tracker_rpb_go/read_event/utils"
)

// You can manage verbosity of log output
// in the package by changing last parameter value
// (comment/uncomment corresponding lines).
var lg = logger.NewPackageLogger("read_event",
	// logger.DebugLevel,
	logger.InfoLevel,
)

var new_line1 = ""
var new_line2 = ""

var BACKLIGHT = true
var ON = true

func switch_backlight(lcd *device.Lcd) {
	BACKLIGHT = !BACKLIGHT
	if BACKLIGHT {
		err := lcd.BacklightOn()
		utils.CheckError(err)
	} else {
		err := lcd.BacklightOff()
		utils.CheckError(err)
	}
}

func switch_on_off(lcd *device.Lcd) {
	ON = !ON
	lcd.Clear()
	switch_backlight(lcd)
}

func RunReadEvent() {
	fmt.Print("[read_event] Running RunReadEvent...\n")

	// Open and map memory to access gpio, check for errors
	err := rpio.Open()
	utils.CheckError(err)
	// Unmap gpio memory when done
	defer rpio.Close()

	// Use mcu pin 40, corresponds to GPIO 21 on the pi
	pin := rpio.Pin(21)
	pin.Input()
	pin.PullUp()
	pin.Detect(rpio.RiseEdge) // enable falling edge event detection

	// Connect to i2c bus
	i2c, err := i2c.NewI2C(0x27, 1)
	utils.CheckError(err)
	defer i2c.Close()

	// Instantiate lcd screen
	lcd, err := device.NewLcd(i2c, device.LCD_16x2)
	utils.CheckError(err)
	lcd.Clear()
	// Enable screen backlight
	err = lcd.BacklightOn()
	utils.CheckError(err)

	currentlyInTheMeeting := false
	event := utils.ParseEvent()
	lastCalled := utils.TimeNow()
	for {
		if pin.EdgeDetected() { // check if event occured
			fmt.Printf("[%s] Button press activated!\n", utils.TimeNow().Format("15:04:05"))
			time.Sleep(time.Second)
			switch_on_off(lcd)
		}
		if !ON {
			time.Sleep(1 * time.Second)
			continue
		}
		now := utils.TimeNow()
		// Monitor when the meeting starts to get the new meeting when it ends
		if !currentlyInTheMeeting && event.Start.Before(now) && event.End.After(now) {
			currentlyInTheMeeting = true
			go buzzer.Play()
		}
		// Get new meeting to show if you just finished a meeting
		if currentlyInTheMeeting && event.End.Before(now) {
			currentlyInTheMeeting = false
			event = utils.ParseEvent()
			lastCalled = utils.TimeNow()
		}
		// Refresh current meeting every 2 minutes to track new meetings throughout the day
		if now.Sub(lastCalled).Seconds() > 120 {
			event = utils.ParseEvent()
			lastCalled = utils.TimeNow()
		}
		
		current_line1 := new_line1
		current_line2 := new_line2

		if !utils.DateEqual(event.Start, utils.TimeNow()) {
			new_line1 = strings.Repeat(" ", 4)
			new_line1 += utils.TimeNow().Format("15:04:05")
			new_line1 += strings.Repeat(" ", 4)
			// lcd can't show empty message so " " is showed
			new_line2 = strings.Repeat(" ", 16) // 16 empty strings to clear the line
		} else {
			new_line1 = utils.TimeNow().Format("15:04:05")
			new_line1 += " "
			new_line1 += event.Time
			new_line2 = event.Description
		}
		if current_line1 != new_line1 || current_line2 != new_line2 {
			// clear display if text has changed
			err = lcd.Clear()
			utils.CheckError(err)
		}
		err = lcd.ShowMessage(new_line1, device.SHOW_LINE_1)
		utils.CheckError(err)
		err = lcd.ShowMessage(new_line2, device.SHOW_LINE_2)
		utils.CheckError(err)
		time.Sleep(10 * time.Millisecond)
	}
}
