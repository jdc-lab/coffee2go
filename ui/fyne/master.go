package fyne

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
)

type Master struct {
	window fyne.Window
	app    fyne.App
}

func NewFyne(width int, height int) (*Master, error) {
	m := &Master{
		app: app.New(),
	}

	m.window = m.app.NewWindow("Coffee2Go")
	m.window.Resize(fyne.NewSize(width, height))
	return m, nil
}

func (m *Master) Run(ready func()) {
	ready()
	m.window.ShowAndRun()
}
