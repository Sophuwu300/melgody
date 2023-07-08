package main

import (
	g "github.com/AllenDang/giu"
	"time"
)

var s string

func gui() {
	g.SingleWindow().Layout(
		g.Label("Hello world from giu. Press the button:"),
		g.Button("Press me!").OnClick(func() {
			go func() {
				for _, x := range "hello from giu" {
					s = s + string(x)
					g.Update()
					time.Sleep(time.Millisecond * 100)
				}
			}()
		}),
		g.Label(s),
	)
}

func main() {
	wnd := g.NewMasterWindow("Hello world", 600, 400, 0)
	wnd.Run(gui)
}