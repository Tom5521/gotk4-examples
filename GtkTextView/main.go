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

func PrintHelloWorld() {
	fmt.Println("Hello World")
}

func activate(app *gtk.Application) {
	w := gtk.NewApplicationWindow(app)

	box := gtk.NewBox(gtk.OrientationVertical, 6)

	tv := gtk.NewTextView()
	tv.SetVExpand(true)
	tb := tv.Buffer()

	tb.SetText("Hello World!")
	tv.SetWrapMode(gtk.WrapWordChar)

	button := gtk.NewButtonWithLabel("Save data")
	button.ConnectClicked(func() {
		err := os.WriteFile(
			"buffer.log",
			[]byte(
				tb.Text(
					tb.StartIter(),
					tb.EndIter(),
					false,
				),
			),
			os.ModePerm,
		)
		if err != nil {
			fmt.Println(err)
		}
	})

	sw := gtk.NewScrolledWindow()
	sw.SetChild(tv)
	box.Append(sw)
	box.Append(button)

	w.SetChild(box)
	w.Show()
}
