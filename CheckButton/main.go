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

	// two ways to declare a CheckButton
	check1, check2 := gtk.NewCheckButton(), gtk.NewCheckButtonWithLabel("Hi world")

	check1.ConnectToggled(func() {
		mode := check1.Active()
		fmt.Println("check1:", mode)
	})
	check2.ConnectToggled(func() {
		mode := check2.Active()
		fmt.Println("check2:", mode)
	})

	box.Append(check1)
	box.Append(check2)

	w.SetChild(box)
	w.Show()
}
