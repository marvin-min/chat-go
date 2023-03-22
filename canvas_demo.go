package main

import (
	"image/color"
	"math"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func setTextContent(c fyne.Canvas, t string) {
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	text := canvas.NewText(t, green)
	text.TextStyle.Bold = true
	c.SetContent(text)
}

func SetContentToCircle(c fyne.Canvas) {
	red := color.NRGBA{R: 0xff, G: 0x33, B: 0x33, A: 0xff}
	circle := canvas.NewCircle(color.White)
	circle.StrokeWidth = 4
	circle.StrokeColor = red
	c.SetContent(circle)
}

func SetContentToRectangle(c fyne.Canvas) {
	blue := color.NRGBA{0, 0, 255, 255}
	rect := canvas.NewRectangle(blue)
	c.SetContent(rect)
	// random integer
	go func() {
		time.Sleep(1 * time.Second)
		rand := uint8(rand.Intn(math.MaxUint8 + 1))
		green := color.NRGBA{0, rand, 0, 255}
		rect.FillColor = green
		rect.Refresh()
	}()
}

func MyCanvas() {
	app := app.New()
	win := app.NewWindow("Canvas Demo")
	myCanvas := win.Canvas()
	// setTextContent(myCanvas, "Hello World")
	// SetContentToCircle(myCanvas)
	SetContentToRectangle(myCanvas)
	win.Resize(fyne.NewSize(200, 200))
	win.ShowAndRun()
}
