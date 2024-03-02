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

const loremIpsum string = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`

func activate(app *gtk.Application) {
	w := gtk.NewApplicationWindow(app)
	w.SetDefaultSize(400, 500)

	box := gtk.NewBox(gtk.OrientationVertical, 6)

	tv := gtk.NewTextView()
	tv.SetVExpand(true)
	tb := tv.Buffer()
	tb.SetText(loremIpsum)
	tv.SetWrapMode(gtk.WrapWordChar)

	saveButton := gtk.NewButtonFromIconName("document-save")
	saveButton.ConnectClicked(func() {
		data := tb.Text(tb.StartIter(), tb.EndIter(), false)

		err := os.WriteFile("buffer.log", []byte(data), os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	})
	cancelButton := gtk.NewButtonFromIconName("window-close")
	cancelButton.ConnectClicked(func() {
		w.Close()
	})

	// It is expected to be displayed below the content and expand horizontally to fill the area.
	bar := gtk.NewActionBar()
	// It allows placing children at the start or the end. In addition,
	// it contains an internal centered box which is centered with respect to the full width of the box,
	// even if the children at either side take up different amounts of space.
	bar.PackStart(saveButton)
	bar.PackStart(cancelButton)

	box.Append(bar)
	box.Append(tv)

	w.SetChild(box)
	w.Show()
}
