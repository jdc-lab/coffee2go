package fyne

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
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
	m.window.SetContent(fyne.NewContainerWithLayout(layout.NewGridLayout(1)))
	m.window.ShowAndRun()
}
