package graph

import (
	"image/color"
	"slices"

	"fyne.io/fyne/v2"
)

func (g *graph) Kruskal(color color.RGBA) int {
	mstWeight := 0
	path := make([]drawer, 0)
	edges := g.getEdges()
	slices.SortFunc(edges, func(a, b *edge) int {
		return a.weight - b.weight
	})
	groups := make(map[*vertex]int)
	groupsCount := 0
	for _, e := range edges[:len(edges)-1] {
		sGroup, sOk := groups[e.start]
		eGroup, eOk := groups[e.end]
		if !sOk && !eOk {
			groups[e.start] = groupsCount
			groups[e.end] = groupsCount
			groupsCount++
			mstWeight += e.weight
			path = append(path, e, e.start, e.end)
			continue
		}
		if sOk && !eOk {
			groups[e.end] = sGroup
			mstWeight += e.weight
			path = append(path, e, e.end)
			continue
		}
		if !sOk && eOk {
			groups[e.start] = eGroup
			mstWeight += e.weight
			path = append(path, e, e.start)
			continue
		}
		if sGroup != eGroup {
			for vert, group := range groups {
				if group == eGroup {
					groups[vert] = sGroup
				}
			}
			mstWeight += e.weight
			path = append(path, e)
		}
	}

	drawNextEdge := getPathDrawer(g, path, color)
	buttonPos := sumPos(g.pos, fyne.NewPos(-g.vertR*2, g.r+g.vertR))
	setButton(g.c, g.vertR, buttonPos, "Kruskal", drawNextEdge)

	return mstWeight
}
