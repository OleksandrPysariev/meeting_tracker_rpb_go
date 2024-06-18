/*

A PWM example by @Drahoslav7, using the go-rpio library

Toggles a LED on physical pin 35 (GPIO pin 19)
Connect a LED with resistor from pin 35 to ground.

*/

package buzzer

import (
        "os"
        "github.com/a-h/beeper"
        "github.com/stianeikeland/go-rpio"
)

func Play() {
        err := rpio.Open()
        if err != nil {
                os.Exit(1)
        }
        defer rpio.Close()

        pin := rpio.Pin(19)
        playTune(pin)
}

func middleVoice(pin rpio.Pin) {
        bpm := 140
        m := beeper.NewMusic(pin, bpm)

        m.Rest(beeper.Quaver)
        m.Note("C5", beeper.Quaver)
        m.Note("C5", beeper.Quaver)
        m.Note("C5", beeper.Quaver)

        m.Rest(beeper.Quaver)
        m.Note("C5", beeper.Quaver)
        m.Note("C5", beeper.Quaver)
        m.Note("C5", beeper.Quaver)

        m.Rest(beeper.Quaver)
        m.Note("C5", beeper.Quaver)
        m.Note("C5", beeper.Quaver)
        m.Note("C5", beeper.Quaver)

        m.Rest(beeper.Quaver)
        m.Note("C5", beeper.Quaver)
        m.Note("C5", beeper.Quaver)
        m.Note("C5", beeper.Quaver)

        m.Rest(beeper.Quaver)
        m.Note("G4", beeper.Quaver)
        m.Note("G4", beeper.Quaver)
        m.Note("G4", beeper.Quaver)

        m.Rest(beeper.Quaver)
        m.Note("G4", beeper.Quaver)
        m.Note("G4", beeper.Quaver)
        m.Note("G4", beeper.Quaver)

        m.Rest(beeper.Quaver)
        m.Note("G4", beeper.Quaver)
        m.Note("G4", beeper.Quaver)
        m.Note("G4", beeper.Quaver)

        m.Rest(beeper.Quaver)
        m.Note("G4", beeper.Quaver)
        m.Note("G4", beeper.Quaver)
        m.Rest(beeper.Quaver)

}

func BassVoice(pin rpio.Pin) {
        bpm := 140
        m := beeper.NewMusic(pin, bpm)
        m.Note("F4", beeper.Quaver)
        m.Note("A4", beeper.Quaver)
        m.Note("A4", beeper.Quaver)
        m.Note("A4", beeper.Quaver)

        m.Note("F4", beeper.Quaver)
        m.Note("A4", beeper.Quaver)
        m.Note("A4", beeper.Quaver)
        m.Note("A4", beeper.Quaver)

        m.Note("F4", beeper.Quaver)
        m.Note("A4", beeper.Quaver)
        m.Note("A4", beeper.Quaver)
        m.Note("A4", beeper.Quaver)

        m.Note("F4", beeper.Quaver)
        m.Note("A4", beeper.Quaver)
        m.Note("A4", beeper.Quaver)
        m.Note("A4", beeper.Quaver)

        m.Note("C4", beeper.Quaver)
        m.Note("E4", beeper.Quaver)
        m.Note("E4", beeper.Quaver)
        m.Note("E4", beeper.Quaver)

        m.Note("C4", beeper.Quaver)
        m.Note("E4", beeper.Quaver)
        m.Note("E4", beeper.Quaver)
        m.Note("E4", beeper.Quaver)

        m.Note("C4", beeper.Quaver)
        m.Note("E4", beeper.Quaver)
        m.Note("E4", beeper.Quaver)
        m.Note("E4", beeper.Quaver)

        m.Note("C4", beeper.Quaver)
        m.Note("E4", beeper.Quaver)
        m.Note("E4", beeper.Quaver)
        m.Note("C4", beeper.Quaver)
}

func playTune(pin rpio.Pin) {
        bpm := 140
        m := beeper.NewMusic(pin, bpm)
        m.Note("F5", beeper.Quaver)
        m.Note("A5", beeper.Quaver)
        m.Note("B5", beeper.Crotchet)

        m.Note("F5", beeper.Quaver)
        m.Note("A5", beeper.Quaver)
        m.Note("B5", beeper.Crotchet)

        m.Note("F5", beeper.Quaver)
        m.Note("A5", beeper.Quaver)
        m.Note("B5", beeper.Quaver)
        m.Note("E6", beeper.Quaver)

        m.Note("D6", beeper.Crotchet)
        m.Note("B5", beeper.Quaver)
        m.Note("C6", beeper.Quaver)

        m.Note("B5", beeper.Quaver)
        m.Note("G5", beeper.Quaver)
        m.Note("E5", beeper.Crotchet*2)
        m.Rest(beeper.Quaver)
        m.Note("D5", beeper.Quaver)

        m.Note("E5", beeper.Quaver)
        m.Note("G5", beeper.Quaver)
        m.Note("E5", beeper.Crotchet*2)
}