package main

import (
	"fmt"

	"github.com/getlantern/systray"
)

//go:generate rsrc -ico icon.ico

func main() {
	systray.Run(OnReady, OnExit) // Initialises system tray icon and starts a goroutine
	fmt.Print("Automation processes")
}
