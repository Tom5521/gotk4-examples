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

func PrintHelloWorld() {
	fmt.Println("Hello World")
}

func activate(app *gtk.Application) {
	w := gtk.NewApplicationWindow(app)

	// In the GtkGrid class,
	// widgets are placed by coordinates.
	box := gtk.NewGrid()

	// Button1
	button1 := gtk.NewButton()
	button1.SetLabel("Click me!")
	button1.ConnectClicked(PrintHelloWorld)
	box.Attach(button1, 0, 0, 1, 1)

	// Button2
	button2 := gtk.NewButton()
	button2.SetLabel("Button 2")
	button2.ConnectClicked(PrintHelloWorld)
	box.Attach(button2, 1, 0, 1, 1)

	// Button3
	button3 := gtk.NewButton()
	button3.SetLabel("Exit")
	button3.ConnectClicked(func() {
		w.Destroy()
	})
	box.Attach(button3, 1, 1, 1, 1)

	w.SetChild(box)
	w.Show()
}
