// +build windows,!linux,!freebsd,!netbsd,!openbsd,!darwin,!js

package beeep

import (
	toast "github.com/go-toast/toast"
)

// Alert displays a desktop notification and plays a default system sound.
func Alert(appId, title, message, appIcon string) error {
	if isWindows10 {
		note := toastNotification(appId, title, message, pathAbs(appIcon))
		note.Audio = toast.Default
		return note.Push()
	}

	if err := Notify(appId, title, message, appIcon); err != nil {
		return err
	}
	return Beep(DefaultFreq, DefaultDuration)
}
