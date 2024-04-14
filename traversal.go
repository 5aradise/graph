package graph

import (
	"image/color"

	"fyne.io/fyne/v2"
)

type drawer interface {
	draw(g *graph, color color.RGBA, width float32)
}

func (g *graph) BFS(color color.RGBA) ([][]uint8, []int) {
	path := make([]drawer, 0)
	visited := make(map[int]bool)
	for _, vert := range g.verts {
		if _, ok := visited[vert.num]; !ok && len(vert.edges) != 0 {
			visited[vert.num] = true
			path = append(path, vert)
			vertBFS(vert, &path, visited)
		}
	}

	drawNextEdge := getPathDrawer(g, path, color)
	buttonPos := sumPos(g.pos, fyne.NewPos(-g.vertR*2, g.r+g.vertR))
	setButton(g.c, g.vertR, buttonPos, "BFS", drawNextEdge)

	return getPathMetrics(len(g.verts), path)
}

func vertBFSt(toCheck *[]*vertex, path *[]drawer, visited map[int]bool) {
	currCheck := make([]*vertex, len(*toCheck))
	copy(currCheck, *toCheck)
	*toCheck = nil
	for _, vert := range currCheck {
		for _, edge := range vert.edges {
			if _, ok := visited[edge.end.num]; !ok {
				visited[edge.end.num] = true
				*path = append(*path, *edge, *edge.end)
				*toCheck = append(*toCheck, edge.end)
			}
		}
	}
}

func vertBFS(vert vertex, path *[]drawer, visited map[int]bool) {
	toCheck := vert.edges
	for len(toCheck) != 0 {
		currCheck := make([]*edge, len(toCheck))
		copy(currCheck, toCheck)
		toCheck = nil
		for _, edge := range currCheck {
			if _, ok := visited[edge.end.num]; !ok {
				visited[edge.end.num] = true
				*path = append(*path, *edge, *edge.end)
				toCheck = append(toCheck, edge.end.edges...)
			}
		}
	}
}

func (g *graph) DFS(color color.RGBA) ([][]uint8, []int) {
	path := make([]drawer, 0)
	visited := make(map[int]bool)
	for _, vert := range g.verts {
		if _, ok := visited[vert.num]; !ok && len(vert.edges) != 0 {
			visited[vert.num] = true
			path = append(path, vert)
			vertDFS(vert, &path, visited)
		}
	}

	drawNextEdge := getPathDrawer(g, path, color)
	buttonPos := sumPos(g.pos, fyne.NewPos(-g.vertR*2, g.r+g.vertR))
	setButton(g.c, g.vertR, buttonPos, "DFS", drawNextEdge)

	return getPathMetrics(len(g.verts), path)
}

func vertDFS(vert vertex, path *[]drawer, visited map[int]bool) {
	var step *edge
	toCheck := vert.edges
	for len(toCheck) != 0 {
		toCheck, step = pop(toCheck)
		if _, ok := visited[step.end.num]; !ok {
			visited[step.end.num] = true
			*path = append(*path, *step, *step.end)
			toCheck = append(toCheck, step.end.edges...)
		}
	}
}

func getPathDrawer(g *graph, path []drawer, color color.RGBA) func() {
	currEdge := 0
	return func() {
		if currEdge < len(path) {
			step := path[currEdge]
			step.draw(g, color, 5)
			currEdge++
		}
	}
}

func getPathMetrics(vertsCount int, path []drawer) ([][]uint8, []int) {
	pathM := make([][]uint8, vertsCount)
	for i := range pathM {
		pathM[i] = make([]uint8, vertsCount)
	}

	seq := make([]int, 0, vertsCount)

	for _, step := range path {
		switch step := step.(type) {
		case vertex:
			seq = append(seq, step.num)
		case edge:
			pathM[step.start.num-1][step.end.num-1] = 1
		}
	}

	return pathM, seq
}

func pop[T any](slice []T) ([]T, T) {
	return slice[:len(slice)-1], slice[len(slice)-1]
}
