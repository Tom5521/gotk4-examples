package main

import (
	"os"
	"strconv"

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
	w.SetDefaultSize(400, 500)

	box := gtk.NewBox(gtk.OrientationVertical, 1)
	for i := range 10 {
		box.Append(gtk.NewButtonWithLabel("Button " + strconv.Itoa(i)))
	}

	// GtkFrame is a widget that surrounds its child with a decorative frame and an optional label.
	frame := gtk.NewFrame("Buttons frame!")
	frame.SetLabelWidget(gtk.NewButtonWithLabel("Buttons!"))
	frame.SetLabelAlign(0.5)
	frame.SetChild(box)

	w.SetChild(frame)
	w.Show()
}
