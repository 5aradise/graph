package graph

func Deg(adjM [][]int, isDir bool) (degs, halfDegs [][]int) {
	if isDir {
		degrees := make([][]int, 3)
		for i := range degrees {
			degrees[i] = make([]int, len(adjM))
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
	degrees := make([][]int, 1)
	degrees[0] = make([]int, len(adjM))
	for vert := range adjM {
		for rel := range vert + 1 {
			if adjM[vert][rel] == 1 {
				degrees[0][vert]++
				degrees[0][rel]++
			}
		}
	}
	return degrees, nil
}

func HomoDeg(degs [][]int) (int, bool) {
	check := degs[0][0]
	for _, vert := range degs[0] {
		if vert != check {
			return 0, false
		}
	}
	return check, true
}

func IsolHang(degs [][]int) (isol, hang []int) {
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

func Path(m [][]int, pathLen int) [][]int {
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

func Path2(m [][]int) [][]int {
	paths := make([][]int, 0)
	m2 := adjDeg(m, 2)
	for start, consLen2 := range m2 {
		for end, conLen2 := range consLen2 {
			if conLen2 != 0 {
				var pathCount int = 0
			endLoop:
				for mid, stmi := range m[start] {
					if stmi != 0 {
						for _, mien := range m[mid] {
							if mien != 0 {
								paths = append(paths, []int{start + 1, mid + 1, end + 1})
								pathCount++
								if pathCount == conLen2 {
									break endLoop
								}
							}
						}
					}
				}
			}
		}
	}
	return paths
}

func getEdges(m [][]int) [][2]int {
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

func usePathCount(m [][]int) {
	clear(m)
}

func adjDeg(m [][]int, deg int) [][]int {
	degM := composeM(m, m)
	for i := 0; i < deg-2; i++ {
		degM = composeM(degM, m)
	}
	return degM
}

func Reachab(m [][]int) [][]int {
	vertCount := len(m)
	r := make([][]int, vertCount)
	acum := make([][]int, vertCount)
	copyM(r, m)
	copyM(acum, m)
	for i := 1; i < vertCount-1; i++ {
		acum = composeM(acum, m)
		AddM(r, acum)
	}
	AddM(r, identM(vertCount))
	r = ToBinM(r)
	return r
}

func StrongCon(m [][]int) [][]int {
	r := Reachab(m)
	strongConM := transM(r)
	MultM(strongConM, r)
	return strongConM
}

func Components(m [][]int) ([][]int, map[int]int) {
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

func CondensMatrix(m [][]int) [][]int {
	ks, vertRelToComp := Components(m)
	kCount := len(ks)
	condM := make([][]int, kCount)
	for k, comp := range ks {
		condM[k] = make([]int, kCount)
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
