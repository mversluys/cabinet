// foreground_windows.go
//go:build windows
// +build windows

package main

import (
	"syscall"
)

var (
	user32              = syscall.MustLoadDLL("user32.dll")
	setForegroundWindow = user32.MustFindProc("SetForegroundWindow")
)

// bringAppToForeground brings the app to the foreground on Windows.
func bringAppToForeground() error {
	hwnd := getCurrentWindowHandle()
	_, _, err := setForegroundWindow.Call(hwnd)
	if err != nil && err.Error() != "The operation completed successfully." {
		return err
	}
	return nil
}

// getCurrentWindowHandle is a placeholder; you may need to implement this
// depending on how you get the HWND of the Wails app.
func getCurrentWindowHandle() uintptr {
	// Implement logic to get HWND for the Wails window
	return 0 // placeholder
}
