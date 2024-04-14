package graph

import (
	"image/color"
	"math"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

const (
	Left  = false
	Right = true
)

func drowArc(c *fyne.Container, size float32, sP, eP fyne.Position, side bool, isDir bool, color color.RGBA, width float32) {
	var k float32 = 1.25
	v := getVector(sP, eP)
	sfi := -math.Pi / float64(3-k/4)
	if side {
		sfi *= -1
	}
	mfi := sfi / 7
	v1Rot := rotateVector(v, sfi)
	v2Rot := rotateVector(v, math.Pi-sfi)
	vRotM := rotateVector(v, mfi)

	s1P := fyne.NewPos(sP.X+v1Rot.X*k*size, sP.Y+v1Rot.Y*k*size)
	e1P := fyne.NewPos(eP.X+v2Rot.X*k*size, eP.Y+v2Rot.Y*k*size)
	mL := getLen(s1P, e1P) / float32(2*math.Cos(mfi))
	mP := fyne.NewPos(s1P.X+vRotM.X*mL, s1P.Y+vRotM.Y*mL)

	drawLine(c, sP, s1P, color, width)
	drawLine(c, s1P, mP, color, width)
	drawLine(c, mP, e1P, color, width)
	drawLine(c, e1P, eP, color, width)
	if isDir {
		drawArrow(c, size, e1P, eP, color, width)
	}
}

func drawLoop(c *fyne.Container, size float32, iP fyne.Position, isDir bool, color color.RGBA, width float32) {
	fi := rand.Float64() * 2 * math.Pi
	sV := rotateVector(fyne.NewPos(1, 0), fi)
	mV := rotateVector(sV, math.Pi/6)
	eV := rotateVector(mV, math.Pi/6)
	sP := sumPos(iP, scalarPos(sV, 2.5*size))
	mP := sumPos(iP, scalarPos(mV, 3*size))
	eP := sumPos(iP, scalarPos(eV, 2.5*size))

	drawLine(c, sumPos(iP, scalarPos(sV, size)), sP, color, width)
	drawLine(c, sP, mP, color, width)
	drawLine(c, mP, eP, color, width)
	drawLine(c, eP, sumPos(iP, scalarPos(eV, size)), color, width)
	if isDir {
		drawArrow(c, size, eP, sumPos(iP, scalarPos(eV, size)), color, width)
	}
}

func drawArrow(c *fyne.Container, size float32, I, O fyne.Position, color color.RGBA, widthL float32) {
	length := size
	width := size / 3
	s := getVector(O, I)
	h := fyne.NewPos(-s.Y, s.X)
	A := fyne.NewPos(O.X+s.X*length, O.Y+s.Y*length)
	R2 := fyne.NewPos(A.X+h.X*width, A.Y+h.Y*width)
	R3 := fyne.NewPos(A.X-h.X*width, A.Y-h.Y*width)

	drawLine(c, O, R2, color, widthL)
	drawLine(c, O, R3, color, widthL)
}

func draw2SidedArrow(c *fyne.Container, size float32, sP, eP fyne.Position, color color.RGBA, width float32) {
	v := getVector(sP, eP)
	h := fyne.NewPos(-v.Y, v.X)
	s1P := sumPos(sP, scalarPos(h, size/5))
	e1P := sumPos(eP, scalarPos(h, size/5))
	drawLine(c, s1P, e1P, color, width)
	drawArrow(c, size, s1P, e1P, color, width)
}

func drawLine(c *fyne.Container, startP, endP fyne.Position, color color.RGBA, width float32) {
	line := canvas.NewLine(color)
	line.StrokeWidth = 1
	line.StrokeWidth = width
	line.Position1 = startP
	line.Position2 = endP

	c.Add(line)
}

func DrawText(c *fyne.Container, size float32, pos fyne.Position, text string, color color.RGBA) {
	title := canvas.NewText(text, color)
	title.TextSize = 1.5 * size
	title.Move(pos)
	c.Add(title)
}
