package main

import (
	"fmt"
	"log"
	"os"
	// "time"

	"gioui.org/app"
	"gioui.org/io/key"
	// "gioui.org/f32"
	"gioui.org/unit"
	"gioui.org/op"
	// "gioui.org/op/clip"
	// "gioui.org/op/paint"
	"gioui.org/layout"
	// "gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// https://jonegil.github.io/gui-with-gio/egg_timer/

type C = layout.Context
type D = layout.Dimensions

type AudioButton struct {
	Label	string
	Width	int
	Height	int
	AudioFile	string
	widget.Clickable
}

func draw(w *app.Window) error {

	// ops are the operations from the UI
	var ops op.Ops

	// startButton is a clickable widget
	//var startButton widget.Clickable

	// th defines the material design style
	th := material.NewTheme()

	audioButton := AudioButton{
		Label:	"q",
		Width:	50,
		Height:	50,
		AudioFile:	"Rosies_Soundboard.mp3",
	}


	// listen for events in the window.
	for {

		// then detect the type
		switch e := w.Event().(type) {

		// this is sent when the application should re-render.
		case app.FrameEvent:

			gtx := app.NewContext(&ops, e)
			

			layout.Flex{
				Axis: layout.Vertical,
				Spacing: layout.SpaceStart,
			}.Layout(gtx,

				layout.Rigid(
					// 4. The button
					func(gtx C) D {
						margins := layout.Inset{
							Top:	unit.Dp(25),
							Bottom:	unit.Dp(25),
							Right:	unit.Dp(35),
							Left:	unit.Dp(35),
						}
						return margins.Layout(gtx,
							func(gtx C) D {

								if audioButton.Clicked(gtx) {
									fmt.Println("Clicked " + audioButton.Label)
									playAudio(audioButton.AudioFile)
								}

								btn := material.Button(th, &audioButton.Clickable, audioButton.Label)
								return btn.Layout(gtx)
							},
						)

					},
				),
			)
			e.Frame(gtx.Ops)

		case key.Event:
			if e.State == key.Press {
				fmt.Println("Key pressed:", e.Name)
			}
			
		// and this is sent when the application should exit
		case app.DestroyEvent:
			return e.Err
		}

	}	
}

func main() {

	go func() {
		w := new(app.Window)
		w.Option(app.Title("Soundboard"))
		w.Option(app.Size(unit.Dp(400), unit.Dp(400)))

		if err := draw(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)

	}()
	app.Main()
	
}
