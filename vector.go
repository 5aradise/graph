package graph

import (
	"math"

	"fyne.io/fyne/v2"
)

func isPositionInArea(pos fyne.Position, poligons ...fyne.Position) bool {
	v1 := getVector(pos, poligons[len(poligons)-1])
	v2 := getVector(pos, poligons[0])
	cross := crossProd(v1, v2)
	for _, poligon := range poligons[1:] {
		v1 = v2
		v2 = getVector(pos, poligon)
		if crossProd(v1, v2)*cross <= 0.00001 {
			return false
		}
	}
	return true
}

func rotateVector(v fyne.Position, fi float64) fyne.Position {
	sin := float32(math.Sin(fi))
	cos := float32(math.Cos(fi))
	return fyne.NewPos(v.X*cos-v.Y*sin, v.X*sin+v.Y*cos)
}

func getVector(sP, eP fyne.Position) fyne.Position {
	v := fyne.NewPos(eP.X-sP.X, eP.Y-sP.Y)
	vLen := getLen(v)
	v = fyne.NewPos(v.X/vLen, v.Y/vLen)
	return v
}

func getLen(points ...fyne.Position) float32 {
	sP := points[0]
	if len(points) == 1 {
		return float32(math.Sqrt(float64(sP.X*sP.X + sP.Y*sP.Y)))
	}
	eP := points[1]
	return float32(math.Sqrt(float64((sP.X-eP.X)*(sP.X-eP.X) + (sP.Y-eP.Y)*(sP.Y-eP.Y))))
}

func crossProd(v1, v2 fyne.Position) float32 {
	return v1.X*v2.Y - v2.X*v1.Y
}

func sumPos(points ...fyne.Position) fyne.Position {
	sumPos := fyne.NewPos(0, 0)
	for _, point := range points {
		sumPos.X += point.X
		sumPos.Y += point.Y
	}
	return sumPos
}

func scalarPos(point fyne.Position, dims ...float32) fyne.Position {
	point.X *= dims[0]
	if len(dims) == 1 {
		point.Y *= dims[0]
	} else {
		point.Y *= dims[1]
	}
	return point
}
