package main

import "github.com/getlantern/systray"

func main() {
	downloadsDir := getDownloadsDirectory()
	process := Recycler(downloadsDir)

	process.run()
	process.print()

	systray.Run(OnReady, OnExit) // Initialises system tray icon and starts a goroutine
}
