package main

import (
	"log"
	"os/exec"

	"github.com/getlantern/systray"
)

func main() {

	cmd := exec.Command("cmd", "/C", "start", "C:\\Windows\\System32\\cmd.exe")

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	downloadsDir := getDownloadsDirectory()
	process := Recycler(downloadsDir)

	process.run()
	process.print()

	systray.Run(OnReady, OnExit) // Initialises system tray icon and starts a goroutine
}
