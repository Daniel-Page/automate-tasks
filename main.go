package main

import (
	"unsafe"
	"workspace/winapi"
	"golang.org/x/sys/windows"
)

func main() {
	//fmt.Print(winapi.Hello())
	hwnd, err := winapi.CreateMainWindow()
	if err != nil {
		panic(err)
	}

	var data winapi.NOTIFYICONDATA

	data.CbSize = uint32(unsafe.Sizeof(data))
	data.UFlags = winapi.NIF_ICON | winapi.NIF_MESSAGE | winapi.NIF_TIP | winapi.NIF_INFO
	data.HWnd = hwnd
	data.UCallbackMessage = winapi.NotifyIconMsg
	copy(data.SzTip[:], windows.StringToUTF16("Tray Icon"))
	copy(data.SzInfo[:], windows.StringToUTF16("Hello from Tay Icon!"))

	icon, err := winapi.LoadImage(
		0,
		windows.StringToUTF16Ptr("icon.ico"),
		winapi.IMAGE_ICON,
		0,
		0,
		winapi.LR_DEFAULTSIZE|winapi.LR_LOADFROMFILE)
	if err != nil {
		panic(err)
	}
	data.HIcon = icon

	if _, err := winapi.Shell_NotifyIcon(winapi.NIM_ADD, &data); err != nil {
		panic(err)
	}

	//ShowWindow(hwnd, SW_SHOW)

	var msg winapi.MSG
	for {
		r, err := winapi.GetMessage(&msg, 0, 0, 0)
		if err != nil {
			panic(err)
		}
		if r == 0 {
			break
		}

		winapi.TranslateMessage(&msg)
		winapi.DispatchMessage(&msg)
	}
}
