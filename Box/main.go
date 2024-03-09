package main

import (
	"os"
	"slices"
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

	// The GtkBox widget arranges child widgets into a single row or column.
	// Whether it is a row or column depends on the value of its GtkOrientable:orientation property.
	// Within the other dimension, all children are allocated the same size.
	// Of course, the GtkWidget:halign and GtkWidget:valign properties can be used on the children to
	// influence their allocation.
	labelsBox := gtk.NewBox(gtk.OrientationVertical, 6)
	labelsBox.SetVExpand(true)

	var labels []*gtk.Label

	for i := range 30 {
		newLabel := gtk.NewLabel("Label " + strconv.Itoa(i))
		labels = append(labels, newLabel)
		labelsBox.Append(newLabel)
	}

	button1 := gtk.NewButtonWithLabel("Change orientation")
	button1.ConnectClicked(func() {
		switch labelsBox.Orientation() {
		case gtk.OrientationVertical:
			labelsBox.SetOrientation(gtk.OrientationHorizontal)
		default:
			labelsBox.SetOrientation(gtk.OrientationVertical)
		}
	})

	button2 := gtk.NewButtonWithLabel("Remove last item")
	button2.ConnectClicked(func() {
		var index int
		if len(labels) <= 0 {
			return
		}
		index = len(labels) - 1
		labelsBox.Remove(labels[index])
		labels = slices.Delete(labels, index, index+1)
	})

	buttonsBox := gtk.NewBox(gtk.OrientationVertical, 6)
	buttonsBox.Append(button1)
	buttonsBox.Append(button2)

	mainBox := gtk.NewBox(gtk.OrientationVertical, 6)
	mainBox.Append(labelsBox)
	mainBox.Append(buttonsBox)

	sbox := gtk.NewScrolledWindow()
	sbox.SetChild(mainBox)
	w.SetChild(sbox)
	w.Show()
}
