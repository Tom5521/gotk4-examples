package main

import (
	"fmt"
	"os"

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

	box := gtk.NewBox(gtk.OrientationVertical, 6)

	// two ways to declare a button
	button1, button2 := gtk.NewButton(), gtk.NewButtonWithLabel("Hello World")

	// You can set the button label with SetLabel method
	button1.SetLabel("Button 1")

	// You can connect the button clicks with the ConnectClick method.
	button1.ConnectClicked(func() {
		fmt.Println(button1.Label())
	})

	button2.ConnectClicked(func() {
		fmt.Println(button2.Label())
	})

	box.Append(button1)
	box.Append(button2)

	w.SetChild(box)
	w.Show()
}
