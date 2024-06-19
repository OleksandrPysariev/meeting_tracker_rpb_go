package read_event

import (
	"fmt"
	"time"

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
	pin.Detect(rpio.FallEdge) // enable falling edge event detection

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
			fmt.Print("Button press activated!")
			time.Sleep(time.Second / 4)
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
		if !utils.DateEqual(event.Start, utils.TimeNow()) {
			line1 := "    "
			line1 += utils.TimeNow().Format("15:04:05")
			line1 += "    "
			// lcd can't show empty message so " " is showed
			line2 := "                " // 16 empty strings to clear the line
			err = lcd.ShowMessage(line1, device.SHOW_LINE_1)
			utils.CheckError(err)
			err = lcd.ShowMessage(line2, device.SHOW_LINE_2)
			utils.CheckError(err)
		} else {
			line1 := utils.TimeNow().Format("15:04:05")
			line1 += " "
			line1 += event.Time
			line2 := event.Description
			err = lcd.ShowMessage(line1, device.SHOW_LINE_1)
			utils.CheckError(err)
			err = lcd.ShowMessage(line2, device.SHOW_LINE_2)
			utils.CheckError(err)

		}
		time.Sleep(10 * time.Millisecond)
	}
}
