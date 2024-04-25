package graph

import (
	"image/color"
	"math"
)

type shape uint8

const (
	Line   shape = 1
	Loop   shape = 2
	ArcL   shape = 3
	ArcR   shape = 4
	Double shape = 5
)

type edge struct {
	start  *vertex
	end    *vertex
	weight int
	dir    bool
	shape  shape
}

func newEdge(verts []vertex, start, end *vertex, weight int, isTwoCons bool, vertR float32, dir bool) edge {
	edge := edge{start, end, weight, dir, Line}
	if start.num == end.num {
		edge.shape = Loop
		return edge
	}
	if isVertexBetween(verts, start.pos, end.pos, vertR) {
		if start.num > end.num && !isTwoCons {
			edge.shape = ArcR
			return edge
		}
		edge.shape = ArcL
		return edge
	}
	if isTwoCons {
		edge.shape = Double
		return edge
	}
	return edge
}

func (e edge) draw(g *graph, color color.RGBA, width float32) {
	dirV := getVector(e.start.pos, e.end.pos)
	switch e.shape {
	case Line:
		drawLine(g.c, sumPos(e.start.pos, scalarPos(dirV, g.vertR)), sumPos(e.end.pos, scalarPos(dirV, -g.vertR)), color, width, e.weight)
		if e.dir {
			drawArrow(g.c, g.vertR, e.start.pos, sumPos(e.end.pos, scalarPos(dirV, -g.vertR)), color, width)
		}
	case Loop:
		drawLoop(g.c, g.vertR, e.start.pos, g.isDir, color, width, e.weight)
	case ArcL:
		drowArc(g.c, g.vertR, sumPos(e.start.pos, scalarPos(rotateVector(dirV, -math.Pi/2), g.vertR)), sumPos(e.end.pos, scalarPos(rotateVector(dirV, -math.Pi/2), g.vertR)), Left, g.isDir, color, width, e.weight)
	case ArcR:
		drowArc(g.c, g.vertR, sumPos(e.start.pos, scalarPos(rotateVector(dirV, math.Pi/2), g.vertR)), sumPos(e.end.pos, scalarPos(rotateVector(dirV, math.Pi/2), g.vertR)), Right, g.isDir, color, width, e.weight)
	case Double:
		draw2SidedArrow(g.c, g.vertR, sumPos(e.start.pos, scalarPos(dirV, g.vertR-3)), sumPos(e.end.pos, scalarPos(dirV, -g.vertR+3)), color, width, e.weight)
	}
}

func setEdges(verts []vertex, vertR float32, m [][]int, isDir bool) {
	if isDir {
		for vert := range m {
			for rel := range m[vert] {
				if m[vert][rel] != 0 {
					edge := newEdge(verts, &verts[vert], &verts[rel], m[vert][rel], m[rel][vert] != 0, vertR, true)
					verts[vert].edges = append(verts[vert].edges, &edge)
				}
			}
		}
		return
	}

	for vert := range verts {
		for rel := range vert + 1 {
			if m[vert][rel] != 0 {
				edge := newEdge(verts, &verts[vert], &verts[rel], m[vert][rel], false, vertR, false)
				verts[vert].edges = append(verts[vert].edges, &edge)
			}
		}
	}
}
