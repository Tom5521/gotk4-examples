package main

import (
	"fmt"
	"os"
	"time"

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

	gtkswitch := gtk.NewSwitch()
	gtkswitch.ConnectStateSet(func(state bool) bool {
		fmt.Println("Switch state:", state)
		return state
	})

	hbox := gtk.NewBox(gtk.OrientationHorizontal, 6)
	hbox.Append(gtkswitch)

	vbox := gtk.NewBox(gtk.OrientationVertical, 6)
	vbox.Append(hbox)

	a := gtk.NewRevealer()
	a.SetChild(vbox)
	go func() {
		time.Sleep(4 * time.Second)
		a.SetRevealChild(true)
	}()

	w.SetChild(a)
	w.Show()
}
