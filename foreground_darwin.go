// foreground_darwin.go
//go:build darwin
// +build darwin

package main

import (
	"github.com/progrium/darwinkit/objc"
)

func bringAppToForeground() error {
	app := objc.Call[objc.Object](objc.GetClass("NSApplication"), objc.Sel("sharedApplication"))
	objc.Call[objc.Void](app, objc.Sel("activateIgnoringOtherApps:"), true)
	return nil
}
