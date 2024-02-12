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

func PrintHi() {
	fmt.Println("Hi")
}

func activate(app *gtk.Application) {
	w := gtk.NewApplicationWindow(app)

	box := gtk.NewNotebook()

	// Button1
	button1 := gtk.NewButton()
	button1.SetLabel("Button 1 in tab1 :D")
	button1.ConnectClicked(PrintHi)
	box.AppendPage(button1, tabNameBox("tab1", box))

	button2 := gtk.NewButton()
	button2.SetLabel("Button 2 in tab2 :v")
	button2.ConnectClicked(PrintHi)
	box.AppendPage(button2, tabNameBox("tab2", box))

	t1 := box.Page(button1)
	t1.SetObjectProperty("tab-expand", true)

	t2 := box.Page(button2)
	t2.SetObjectProperty("tab-expand", true)

	w.SetChild(box)
	w.Show()
}

func tabNameBox(tabName string, notebook *gtk.Notebook) gtk.Widgetter {
	box := gtk.NewBox(gtk.OrientationHorizontal, 6)
	lb := gtk.NewLabel(tabName)
	lb.SetHExpand(true)

	click := gtk.NewGestureClick()
	click.ConnectReleased(func(_ int, _, _ float64) {
		click.SetState(gtk.EventSequenceClaimed)
		p := notebook.PageNum(box)
		notebook.RemovePage(p)
	})

	img := gtk.NewImageFromIconName("emblem-unreadable")
	img.SetFocusOnClick(true)
	img.AddController(click)

	box.Append(lb)
	box.Append(img)

	return box
}
