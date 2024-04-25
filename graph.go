package graph

import (
	"image/color"

	"fyne.io/fyne/v2"
)

type graph struct {
	c        *fyne.Container
	verts    []vertex
	pos      fyne.Position
	r        float32
	vertR    float32
	sideC    int
	isMidVer bool
	isDir    bool
}

func NewGraph(c *fyne.Container, topLeftCorner fyne.Position, graphRadius, vertexRadius float32, sideCount int, isVertexInMid bool, adjMatrix [][]int, isDir bool) graph {
	g := graph{
		c:        c,
		r:        graphRadius,
		vertR:    vertexRadius,
		sideC:    sideCount,
		isMidVer: isVertexInMid,
		isDir:    isDir,
	}
	g.pos = sumPos(topLeftCorner, fyne.NewPos(g.r, g.r))
	g.verts = createVertices(&g, adjMatrix)
	return g
}

func (g *graph) Draw() {
	for _, vert := range g.verts {
		for _, edge := range vert.edges {
			edge.draw(g, color.RGBA{255, 255, 255, 255}, g.vertR/20)
		}
	}
	for _, vert := range g.verts {
		vert.draw(g, color.RGBA{255, 255, 255, 255}, 1)
	}
}

func (g *graph) getEdges() []*edge {
	edges := make([]*edge, 0)
	for _, vert := range g.verts {
		edges = append(edges, vert.edges...)
	}
	return edges
}
