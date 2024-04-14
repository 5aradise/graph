package graph

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func setButton(c *fyne.Container, size float32, pos fyne.Position, label string, callback func()) {
	content := widget.NewButton(label, callback)
	content.Move(pos)
	content.Resize(fyne.NewSize(size*4, size*2))
	c.Add(content)
}
