package graph

import (
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type vertex struct {
	num   int
	pos   fyne.Position
	edges []*edge
}

func (v vertex) draw(g *graph, colorV color.RGBA, width float32) {
	body := canvas.NewCircle(color.Black)
	body.StrokeColor = colorV
	body.StrokeWidth = width
	body.Position1 = fyne.NewPos(v.pos.X-g.vertR, v.pos.Y-g.vertR)
	body.Position2 = fyne.NewPos(v.pos.X+g.vertR, v.pos.Y+g.vertR)

	g.c.Add(body)
	DrawFormattedNum(g.c, g.vertR, v.pos, v.num, colorV)
}

func createVertices(g *graph, m [][]int) []vertex {
	verts := make([]vertex, len(m))

	for i := range verts {
		verts[i].num = i + 1
	}

	vertsPos := getVerticesPos(len(verts), g.sideC, g.isMidVer, g.pos, g.r)
	for i := range verts {
		verts[i].pos = vertsPos[i]
	}

	setEdges(verts, g.vertR, m, g.isDir)

	return verts
}

func getVerticesPos(vertCount, sideCount int, isVertexInMid bool, midPos fyne.Position, graphR float32) []fyne.Position {
	if sideCount == 0 || sideCount > vertCount {
		sideCount = vertCount
		if isVertexInMid {
			sideCount--
		}
	}

	vertsPos := make([]fyne.Position, vertCount)
	curVert := 0
	if isVertexInMid {
		vertsPos[curVert] = midPos
		curVert++
	}

	if sideCount == 0 {
		return vertsPos
	}

	extrSideSize := (vertCount - curVert) % sideCount
	mainSideSize := (vertCount - curVert - extrSideSize) / sideCount

	sidesSize := make([]int, sideCount)
	for i := range sidesSize {
		sidesSize[i] = mainSideSize
	}
	for i := 0; i < 2*extrSideSize; i += 2 {
		sidesSize[i%sideCount]++
	}

	sideFi := 0.0
	dfi := 2 * math.Pi / float64(sideCount)
	sideLen := 2 * graphR * float32(math.Sin(dfi/2))
	dy, dx := math.Sincos(sideFi)

	curPos := fyne.NewPos(midPos.X-sideLen/2, midPos.Y-float32(math.Sqrt(float64(graphR*graphR-sideLen*sideLen/4))))
	vertsPos[curVert] = curPos
	curVert++

	for _, sideSize := range sidesSize {
		sidePos := getSidePos(sideSize, sideLen, fyne.NewPos(float32(dx), float32(dy)), curPos)
		for j := range sidePos {
			vertsPos[curVert] = sidePos[j]
			curVert++
			if curVert == vertCount {
				return vertsPos
			}
		}
		curPos = sidePos[sideSize-1]
		sideFi += dfi
		dy, dx = math.Sincos(sideFi)
	}

	return vertsPos
}

func getSidePos(size int, length float32, dv, curPos fyne.Position) []fyne.Position {
	sidePos := make([]fyne.Position, size)
	delay := length / float32(size)
	for i := range sidePos {
		curPos = sumPos(curPos, scalarPos(dv, delay))
		sidePos[i] = curPos
	}
	return sidePos
}

func isVertexBetween(verts []vertex, sP, eP fyne.Position, vertR float32) bool {
	v := getVector(sP, eP)
	h := fyne.NewPos(-v.Y, v.X)
	s1P := sumPos(sP, scalarPos(h, vertR, vertR))
	s2P := sumPos(sP, scalarPos(h, -vertR, -vertR))
	e1P := sumPos(eP, scalarPos(h, vertR, vertR))
	e2P := sumPos(eP, scalarPos(h, -vertR, -vertR))
	for _, vert := range verts {
		if isPositionInArea(vert.pos, s1P, s2P, e2P, e1P) {
			return true
		}
	}
	return false
}
