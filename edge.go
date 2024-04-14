package graph

import (
	"image/color"
	"math"
)

const (
	Line   uint8 = 1
	Loop   uint8 = 2
	ArcL   uint8 = 3
	ArcR   uint8 = 4
	Double uint8 = 5
)

type edge struct {
	start *vertex
	end   *vertex
	dir   bool
	typ   uint8
}

func newEdge(verts []vertex, start, end *vertex, isTwoCons bool, vertR float32, dir bool) edge {
	if start.num == end.num {
		return edge{start, end, dir, Loop}
	}
	if isVertexBetween(verts, start.pos, end.pos, vertR) {
		if start.num > end.num && !isTwoCons {
			return edge{start, end, dir, ArcR}
		}
		return edge{start, end, dir, ArcL}
	}
	if isTwoCons {
		return edge{start, end, dir, Double}
	}
	return edge{start, end, dir, Line}
}

func (e edge) draw(g *graph, color color.RGBA, width float32) {
	dirV := getVector(e.start.pos, e.end.pos)
	switch e.typ {
	case Line:
		drawLine(g.c, sumPos(e.start.pos, scalarPos(dirV, g.vertR)), sumPos(e.end.pos, scalarPos(dirV, -g.vertR)), color, width)
		if e.dir {
			drawArrow(g.c, g.vertR, e.start.pos, sumPos(e.end.pos, scalarPos(dirV, -g.vertR)), color, width)
		}
	case Loop:
		drawLoop(g.c, g.vertR, e.start.pos, g.isDir, color, width)
	case ArcL:
		drowArc(g.c, g.vertR, sumPos(e.start.pos, scalarPos(rotateVector(dirV, -math.Pi/2), g.vertR)), sumPos(e.end.pos, scalarPos(rotateVector(dirV, -math.Pi/2), g.vertR)), Left, g.isDir, color, width)
	case ArcR:
		drowArc(g.c, g.vertR, sumPos(e.start.pos, scalarPos(rotateVector(dirV, math.Pi/2), g.vertR)), sumPos(e.end.pos, scalarPos(rotateVector(dirV, math.Pi/2), g.vertR)), Right, g.isDir, color, width)
	case Double:
		draw2SidedArrow(g.c, g.vertR, sumPos(e.start.pos, scalarPos(dirV, g.vertR-3)), sumPos(e.end.pos, scalarPos(dirV, -g.vertR+3)), color, width)
	}
}

func setEdges(verts []vertex, vertR float32, m [][]uint8, isDir bool) {
	if isDir {
		for vert := range m {
			for rel := range m[vert] {
				if m[vert][rel] == 1 {
					edge := newEdge(verts, &verts[vert], &verts[rel], m[rel][vert] == 1, vertR, isDir)
					verts[vert].edges = append(verts[vert].edges, &edge)
				}
			}
		}
		return
	}

	for vert := range verts {
		for rel := 0; rel < vert+1; rel++ {
			if m[vert][rel] == 1 {
				edge := newEdge(verts, &verts[vert], &verts[rel], false, vertR, isDir)
				verts[vert].edges = append(verts[vert].edges, &edge)
			}
		}
	}
}
