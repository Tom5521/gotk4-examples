//go:generate youtube-dl https://www.youtube.com/watch?v=Elj4zDLqJvw -o video

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

func activate(app *gtk.Application) {
	w := gtk.NewApplicationWindow(app)
	w.SetDefaultSize(500, 500)
	w.SetTitle("Example video")

	video := gtk.NewVideoForFilename("Video/video.mkv")

	w.SetChild(video)
	w.Show()
}
