package read_event

import (
	"fmt"
	"os"
	"os/exec"
)

func ClearTerminal() {
	// TODO: implement other platforms support 
	// https://stackoverflow.com/questions/22891644/how-can-i-clear-the-terminal-screen-in-go
	cmd := exec.Command("clear") // Linux
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Say(line1 string, line2 string) {
	ClearTerminal()
	if len(line1) >= 16 {
		line1 = line1[:15]
	}
	if len(line2) >= 16 {
		line2 = line2[:15]
	}
	fmt.Printf("%s\n%s\n", line1, line2)
	fmt.Printf("\033[0;0H")
}