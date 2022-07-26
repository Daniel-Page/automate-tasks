package main

import (
	_ "embed"
	"fmt"

	"github.com/getlantern/systray"
)

//go:generate rsrc -ico icon.ico

//go:embed icon.ico
var icon []byte

func main() {
	systray.Run(onReady, onExit) // Starts system tray icon
	fmt.Print("Automation processes")
}

func onReady() {
	systray.SetIcon(icon)
	systray.SetTitle("Automate Tasks")
	systray.SetTooltip("Automate Tasks")

	mSettings := systray.AddMenuItem("Settings", "Q")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	// Create goroutine
	go func() {
		for {
			select {
			case <-mQuit.ClickedCh:
				systray.Quit()
				fmt.Println("Quit2 now...")
			case <-mSettings.ClickedCh:
				fmt.Println("Quit2 now...")
			}
		}
	}()
}

func onExit() {
	// clean up here
}
