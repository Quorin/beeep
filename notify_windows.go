// +build windows,!linux,!freebsd,!netbsd,!openbsd,!darwin,!js

package beeep

import (
	"errors"
	"os/exec"
	"syscall"
	"time"

	toast "github.com/go-toast/toast"
	"github.com/tadvi/systray"
	"golang.org/x/sys/windows/registry"
)

var isWindows10 bool

func init() {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		return
	}
	defer k.Close()

	maj, _, err := k.GetIntegerValue("CurrentMajorVersionNumber")
	if err != nil {
		return
	}

	isWindows10 = maj == 10
}

// Notify sends desktop notification.
func Notify(appId, title, message, appIcon string) error {
	if isWindows10 {
		return toastNotify(appId, title, message, appIcon)
	}

	err := baloonNotify(title, message, appIcon, false)
	if err != nil {
		e := msgNotify(title, message)
		if e != nil {
			return errors.New("beeep: " + err.Error() + "; " + e.Error())
		}
	}

	return nil

}

func msgNotify(title, message string) error {
	msg, err := exec.LookPath("msg")
	if err != nil {
		return err
	}
	cmd := exec.Command(msg, "*", "/TIME:3", title+"\n\n"+message)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return cmd.Run()
}

func baloonNotify(title, message, appIcon string, bigIcon bool) error {
	tray, err := systray.New()
	if err != nil {
		return err
	}

	err = tray.ShowCustom(pathAbs(appIcon), title)
	if err != nil {
		return err
	}

	go func() {
		tray.Run()
		time.Sleep(3 * time.Second)
		tray.Stop()
	}()

	return tray.ShowMessage(title, message, bigIcon)
}

func toastNotify(appId, title, message, appIcon string) error {
	notification := toastNotification(appId, title, message, pathAbs(appIcon))
	return notification.Push()
}

func toastNotification(appId, title, message, appIcon string) toast.Notification {
	return toast.Notification{
		AppID:   appId,
		Title:   title,
		Message: message,
		Icon:    appIcon,
	}
}
