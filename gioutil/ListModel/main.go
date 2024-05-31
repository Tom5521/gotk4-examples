package main

import (
	"fmt"
	"os"

	"github.com/diamondburned/gotk4/pkg/core/gioutil"
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

type Book struct {
	Name  string
	Pages int
}

func activate(app *gtk.Application) {
	w := gtk.NewApplicationWindow(app)

	books := []Book{
		{"El ingenioso hidalgo don Quijote de la Mancha", 462}, // I LOVE this book :D
		{"Le petit prince", 92},
		{"Romeo and Juliet", 480},
	}

	// ListModel is a wrapper around an internal ListModel that allows any Go value
	// to be used as a list item. Internally, it uses core/gbox to store the values
	// in a global registry for later retrieval.
	model := gioutil.NewListModel[Book]()
	model.Splice(0, 0, books...) // Fill the model

	factory := gtk.NewSignalListItemFactory()
	factory.ConnectSetup(func(listitem *gtk.ListItem) {
		listitem.SetChild(gtk.NewLabel(""))
	})
	factory.ConnectBind(func(listitem *gtk.ListItem) {
		label := listitem.Child().(*gtk.Label)
		book := model.Item(int(listitem.Position()))
		label.SetText(fmt.Sprintf("Name: %s | Pages: %v", book.Name, book.Pages))
	})

	selectionModel := gtk.NewNoSelection(model.ListModel)
	lv := gtk.NewListView(selectionModel, &factory.ListItemFactory)

	w.SetChild(lv)
	w.Show()
}
