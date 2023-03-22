package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 15:04:05")
	clock.SetText(formatted)
}

func MyClock() {
	a := app.New()
	w := a.NewWindow("Clock")
	w.Resize(fyne.NewSize(200, 200))
	clock := widget.NewLabel("")
	updateTime(clock)
	w.SetContent(clock)
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()
	w.Show()
	w2 := a.NewWindow("Larger")
	w2.SetContent(widget.NewLabel("More content"))
	w2.Resize(fyne.NewSize(100, 100))
	w2.Show()
	a.Run()
}
