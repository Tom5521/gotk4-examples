package main

import (
	"math/rand/v2"
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
	w.SetDefaultSize(600, 600)

	// GtkFixed places its child widgets at fixed positions and with fixed sizes.
	// GtkFixed performs no automatic layout management.
	//
	// For most applications, you should not use this container!
	// It keeps you from having to learn about the other GTK containers,
	// but it results in broken applications. With GtkFixed,
	// the following things will result in truncated text, overlapping widgets,
	// and other display bugs.
	fixed := gtk.NewFixed()

	// I make a function to avoid repeating many times the same code.
	addController := func(widget AddControllerer) {
		MakeRandomMoveController(widget, w, fixed)
	}

	button := gtk.NewButtonWithLabel("Click Me!")
	button.ConnectClicked(func() {
		button.SetLabel("How do you clicked me?")
	})
	addController(button)

	label1 := gtk.NewLabel("Hi world!")
	addController(label1)
	label2 := gtk.NewLabel("Hi world! (2)")
	addController(label2)

	fixed.Put(button, 200, 200)
	fixed.Put(label1, 200, 300)
	fixed.Put(label2, 500, 400)

	w.SetChild(fixed)
	w.Show()
}

type AddControllerer interface {
	gtk.Widgetter
	AddController(gtk.EventControllerer)
}

func MakeRandomMoveController(widget AddControllerer, w *gtk.ApplicationWindow, fixed *gtk.Fixed) {
	// We create a new controller.
	controller := gtk.NewEventControllerMotion()
	controller.ConnectEnter(func(_, _ float64) {
		// Create the new coordinates for the widget based on the window size.
		// And I subtract 30 pixels to prevent the widget from being
		// able to protrude too much from the edge of the window.
		newX := float64(rand.N(w.Width() - 30))
		newY := float64(rand.N(w.Height() - 30))
		// Apply the new coordinates.
		fixed.Move(widget, newX, newY)
	})
	// Now we add the controller to the widget
	widget.AddController(controller)
}
