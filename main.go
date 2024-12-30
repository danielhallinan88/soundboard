package main

import (
//	"image/color"
	"fmt"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/unit"
	"gioui.org/op"
	"gioui.org/layout"
//	"gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// https://jonegil.github.io/gui-with-gio/egg_timer/
// https://jonegil.github.io/gui-with-gio/egg_timer/03_button.html

func draw(w *app.Window) error {
	// ops are the operations from the UI
	var ops op.Ops

	// startButton is a clickable widget
	var startButton widget.Clickable

	// th defines the material design style
	th := material.NewTheme()

	// listen for events in the window.
	for {
		// first grab the event
		//evt := w.Event()

		// then detect the type
		//switch typ := evt.(type) {
		switch e := w.Event().(type) {

		// this is sent when the application should re-render.
		case app.FrameEvent:

			//gtx := app.NewContext(&ops, typ)
			gtx := app.NewContext(&ops, e)

			layout.Flex{
				Axis: layout.Vertical,
				Spacing: layout.SpaceStart,
			}.Layout(gtx,
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(th, &startButton, "Start")
						return btn.Layout(gtx)
					},
				),
				layout.Rigid(
					layout.Spacer{Height: unit.Dp(25)}.Layout,
				),
			)
			e.Frame(gtx.Ops)
			//typ.Frame(gtx.Ops)

		// and this is sent when the application should exit
		case app.DestroyEvent:
			return e.Err
		}

	}	
}

func main() {
	fmt.Println("Start soundboard.")
	//playAudio("Rosies_Soundboard.mp3")

	go func() {
		w := new(app.Window)
		w.Option(app.Title("Egg timer"))
		w.Option(app.Size(unit.Dp(400), unit.Dp(600)))

		if err := draw(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)





	}()
	app.Main()
	
}
