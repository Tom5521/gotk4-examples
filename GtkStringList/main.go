package main

import (
	"fmt"
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

func activate(app *gtk.Application) {
	w := gtk.NewApplicationWindow(app)

	// gtk.StringList is a structure that,
	// as you would expect, creates a list model based on a slice of strings.

	items := func() []string {
		var items []string
		for range 10 {
			items = append(items, gofakeit.Sentence(5))
		}
		return items
	}()
	model := gtk.NewStringList(items)

	// In order to implement it, first we have to create a factory,
	// which will be what builds each element of the list.

	factory := gtk.NewSignalListItemFactory()

	// When creating the factory instance, we connect to the setup signal,
	// That will assign a type to each row in the list,
	// it can be anything that enters the gtk.Widgetter interface.
	factory.ConnectSetup(func(listitem *gtk.ListItem) {
		box := gtk.NewBox(gtk.OrientationHorizontal, 6)
		box.SetHomogeneous(true)
		label := gtk.NewLabel("")
		button := gtk.NewButton()
		box.Append(label)
		box.Append(button)
		listitem.SetChild(box)
	})
	// After that we connect the bind signal,
	// which will be the one that assigns the data to the widget that we specified previously
	factory.ConnectBind(func(listitem *gtk.ListItem) {
		box := listitem.Child().(*gtk.Box) // Determine the interface type.
		// We will have to do the same with the data, and as this is a StringList,
		// the data type will obviously be a string.
		obj := listitem.Item().Cast().(*gtk.StringObject)

		widgets := box.ObserveChildren()

		label := widgets.Item(0).Cast().(*gtk.Label)
		// Apply the data.
		label.SetLabel(obj.String())
		// We configure the preferences for each item.
		label.SetHAlign(1)

		button := widgets.Item(1).Cast().(*gtk.Button)
		button.SetLabel("Print data")
		button.ConnectClicked(func() {
			fmt.Println(obj.String())
		})
	})

	// Now we create a selection model, which as its name indicates,
	// determines how the selection will be, and to that model you pass the base model,
	// that is, StringList.
	smodel := gtk.NewMultiSelection(model)
	// Now we create the widget gtk.ListView,
	// And we specify the factory and the selection model,
	// which already brings with it the StringList model
	list := gtk.NewListView(smodel, &factory.ListItemFactory)

	box := gtk.NewBox(gtk.OrientationVertical, 6)

	entry := gtk.NewEntry()
	button := gtk.NewButtonWithLabel("Add Text")
	button.ConnectClicked(func() {
		model.Append(entry.Text())
	})

	widgets := []gtk.Widgetter{entry, button, list}
	for _, w := range widgets {
		box.Append(w)
	}

	sbox := gtk.NewScrolledWindow()
	sbox.SetChild(box)
	w.SetChild(sbox)
	w.Show()
}
