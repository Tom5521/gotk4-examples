package main

import (
	"os"
	"time"

	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func main() {
	app := gtk.NewApplication("com.test.window", gio.ApplicationFlagsNone)
	app.ConnectActivate(func() {
		activate(app)
	})
	if code := app.Run(os.Args); code > 0 {
		os.Exit(code)
	}
}

func activate(app *gtk.Application) {
	w := gtk.NewApplicationWindow(app)
	w.SetDefaultSize(300, 300)

	box := gtk.NewBox(gtk.OrientationVertical, 6)

	label1 := gtk.NewLabel("Hi")

	go func() {
		// You can set the label text with SetLabel or SetText method
		label1.SetLabel("Button 1")
		time.Sleep(time.Second * 2)
		label1.SetText("Button 1 (modified)")
	}()

	box.Append(label1)

	w.SetChild(box)
	w.Show()
}
