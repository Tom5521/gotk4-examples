package main

import (
	"os"

	"github.com/brianvoe/gofakeit"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func AppendWidgets(parent *gtk.Box, widgets ...gtk.Widgetter) {
	for _, w := range widgets {
		parent.Append(w)
	}
}

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
	w.SetDefaultSize(650, 500)

	box := gtk.NewBox(gtk.OrientationVertical, 6)

	// The GtkPaned type is a container that allows the user to freely adjust its size.
	// It is implemented by providing two interfaces of type Gtk.Widget, here is an example.
	paned := gtk.NewPaned(gtk.OrientationHorizontal)
	paned.SetVExpand(true)

	r := Right()
	l := Left()

	paned.SetStartChild(r)
	paned.SetEndChild(l)

	button := gtk.NewButtonWithLabel("Change orientation")
	button.ConnectClicked(func() {
		switch paned.Orientation() {
		case gtk.OrientationVertical:
			paned.SetOrientation(gtk.OrientationHorizontal)
		default:
			paned.SetOrientation(gtk.OrientationVertical)
		}
	})

	AppendWidgets(box,
		paned,
		button,
	)

	w.SetChild(box)
	w.Show()
}

func Right() *gtk.ScrolledWindow {
	sbox := gtk.NewScrolledWindow()
	box := gtk.NewBox(gtk.OrientationVertical, 2)
	l := gtk.NewLabel("Hello world!")

	for range 10 {
		b := gtk.NewButton()
		b.SetLabel(gofakeit.Word())
		box.Append(b)
	}

	AppendWidgets(box, l)
	sbox.SetChild(box)
	return sbox
}

func Left() *gtk.ScrolledWindow {
	sbox := gtk.NewScrolledWindow()
	box := gtk.NewBox(gtk.OrientationVertical, 6)
	l := gtk.NewLabel("hi 2")
	l2 := gtk.NewLabel("I like cheese")

	for range 10 {
		box.Append(gtk.NewLabel(gofakeit.Word()))
	}

	AppendWidgets(box, l, l2)
	sbox.SetChild(box)
	return sbox
}
