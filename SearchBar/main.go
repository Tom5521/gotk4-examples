package main

import (
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

// Translated from: https://gitlab.gnome.org/GNOME/gtk/tree/main/examples/search-bar.c

func activate(app *gtk.Application) {
	w := gtk.NewApplicationWindow(app)
	w.Present()

	searchBar := gtk.NewSearchBar()
	searchBar.SetVAlign(gtk.AlignStart)
	w.SetChild(searchBar)

	box := gtk.NewBox(gtk.OrientationHorizontal, 6)
	searchBar.SetChild(box)

	entry := gtk.NewEntry()
	entry.SetHExpand(true)
	box.Append(entry)

	menuButton := gtk.NewButton()
	box.Append(menuButton)

	searchBar.ConnectEntry(entry)
	searchBar.SetKeyCaptureWidget(w)
}
