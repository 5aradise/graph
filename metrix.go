package graph

func Deg(adjM [][]uint8) (degs, halfDegs [][]uint8) {
	degrees := make([][]uint8, 3)
	for i := range degrees {
		degrees[i] = make([]uint8, len(adjM))
	}
	for vert, rels := range adjM {
		for rel, relation := range rels {
			if relation == 1 {
				degrees[0][vert]++
				degrees[0][rel]++
				degrees[1][vert]++
				degrees[2][rel]++
			}
		}
	}
	return degrees[:1], degrees[1:3]
}

func HomoDeg(degs [][]uint8) (uint8, bool) {
	check := degs[0][0]
	for _, vert := range degs[0] {
		if vert != check {
			return 0, false
		}
	}
	return check, true
}

func IsolHang(degs [][]uint8) (isol, hang []int) {
	for vert, deg := range degs[0] {
		if deg == 0 {
			isol = append(isol, vert+1)
		}
		if deg == 1 {
			hang = append(hang, vert+1)
		}
	}
	return
}

func Path(m [][]uint8, pathLen int) [][]int {
	paths := make([][]int, 0)
	pathCountM := adjDeg(m, pathLen)
	usePathCount(pathCountM)
	edges := getEdges(m)
	for start := range m {
		paths = append(paths, createPath(start, edges, pathLen)...)
	}
	return paths
}

func createPath(start int, edges [][2]int, pathLen int) [][]int {
	pathsWithStart := make([][]int, 0)
	if pathLen == 1 {
		for _, edge := range edges {
			if edge[0] == start {
				pathsWithStart = append(pathsWithStart, []int{edge[0] + 1, edge[1] + 1})
			}
		}
		return pathsWithStart
	}
	for edgeCount, edge := range edges {
		if edge[0] == start {
			edgesC := append(make([][2]int, 0), edges...)
			paths := createPath(edge[1], append(edgesC[:edgeCount], edgesC[edgeCount+1:]...), pathLen-1)
			for _, path := range paths {
				pathWithStart := []int{start + 1}
				pathWithStart = append(pathWithStart, path...)
				pathsWithStart = append(pathsWithStart, pathWithStart)
			}
		}
	}
	return pathsWithStart
}

func getEdges(m [][]uint8) [][2]int {
	edges := make([][2]int, 0)
	for vert, rels := range m {
		for vertRel, rel := range rels {
			if rel == 1 {
				edge := [2]int{vert, vertRel}
				edges = append(edges, edge)
			}
		}
	}
	return edges
}

func usePathCount(m [][]uint8) {
	clear(m)
}

func adjDeg(m [][]uint8, deg int) [][]uint8 {
	degM := composeM(m, m)
	for i := 0; i < deg-2; i++ {
		degM = composeM(degM, m)
	}
	return degM
}

func Reachab(m [][]uint8) [][]uint8 {
	vertCount := len(m)
	r := make([][]uint8, vertCount)
	acum := make([][]uint8, vertCount)
	copyM(r, m)
	copyM(acum, m)
	for i := 1; i < vertCount-1; i++ {
		acum = composeM(acum, m)
		addM(r, acum)
	}
	addM(r, identM(vertCount))
	r = toBinM(r)
	return r
}

func StrongCon(m [][]uint8) [][]uint8 {
	r := Reachab(m)
	strongConM := transM(r)
	multM(strongConM, r)
	return strongConM
}

func Components(m [][]uint8) ([][]int, map[int]int) {
	s := StrongCon(m)
	comps := make([][]int, 0, len(s))
	rel := make(map[int]int)
	for v, row := range s {
		if _, include := rel[v+1]; include {
			continue
		}
		con := make([]int, 0)
		for vert, isInCon := range row {
			if isInCon == 1 {
				con = append(con, vert+1)
				rel[vert+1] = len(comps) + 1
			}
		}
		comps = append(comps, con)
	}
	return comps, rel
}

func CondensMatrix(m [][]uint8) [][]uint8 {
	ks, vertRelToComp := Components(m)
	kCount := len(ks)
	condM := make([][]uint8, kCount)
	for k, comp := range ks {
		condM[k] = make([]uint8, kCount)
		for _, vertex := range comp {
			vertexRels := m[vertex-1]
			for rel, isRel := range vertexRels {
				kRel := vertRelToComp[rel+1] - 1
				if isRel == 1 && k != kRel {
					condM[k][kRel] = 1
				}
			}
		}
	}
	return condM
}
