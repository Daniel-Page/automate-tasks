package main

import (
	_ "embed"

	"github.com/getlantern/systray"
)

//go:generate rsrc -ico icon.ico

//go:embed icon.ico
var icon []byte

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(icon)
	systray.SetTitle("Awesome App")
	systray.SetTooltip("Pretty awesome")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	// Sets the icon of a menu item. Only available on Mac and Windows.
	mQuit.SetIcon(icon)
}

func onExit() {
	// clean up here
}
