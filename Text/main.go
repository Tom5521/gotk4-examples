package main

import (
	"os"

	"github.com/brianvoe/gofakeit"
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

const loremIpsum string = "Lorem ipsum"

func activate(app *gtk.Application) {
	w := gtk.NewApplicationWindow(app)

	// Unlike a label, this text can be selected and/or edited.
	// The GtkText widget is a single-line text entry widget.
	text := gtk.NewText()
	text.SetEditable(false)
	text.SetText(loremIpsum)

	button := gtk.NewButtonWithLabel("Change text")
	button.ConnectClicked(func() {
		text.SetText(gofakeit.Word())
	})

	box := gtk.NewBox(gtk.OrientationVertical, 6)
	box.Append(text)
	box.Append(button)

	w.SetChild(box)
	w.Show()
}
