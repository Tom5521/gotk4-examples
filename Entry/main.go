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

func activate(app *gtk.Application) {
	w := gtk.NewApplicationWindow(app)

	// A fairly large set of key bindings are supported by default.
	// If the entered text is longer than the allocation of the widget,
	// the widget will scroll so that the cursor position is visible.
	entry := gtk.NewEntry()
	entry.ConnectChanged(func() {
		// Alter the text through the buffer, to avoid possible warnings.
		//
		// If we alter it directly with entry.SetText a warning like this will appear:
		//
		// 2024/03/01 09:44:07 WARN Cannot end irreversible action while in user action priority=4
		// code_file=../gtk/gtk/gtktexthistory.c code_line=1042
		// code_func=gtk_text_history_end_irreversible_action glib_domain=Gtk
		switch entry.Text() {
		case "reset":
			entry.Buffer().SetText("", 0)
		case "lock":
			entry.SetEditable(false)
		case "random":
			newWord := gofakeit.Word()
			entry.Buffer().SetText(newWord, len(newWord))
		}
	})

	button := gtk.NewButtonWithLabel("Unlock entry")
	button.ConnectClicked(func() {
		entry.SetEditable(true)
		entry.Buffer().SetText("", 0)
	})

	vbox := gtk.NewBox(gtk.OrientationVertical, 6)
	vbox.Append(entry)
	vbox.Append(button)

	w.SetChild(vbox)
	w.Show()
}
