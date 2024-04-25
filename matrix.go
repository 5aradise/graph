package graph

import "math/rand"

type number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func RandIntMatrix(size int, k float64, r *rand.Rand) [][]int {
	m := make([][]int, size)
	for i := range m {
		m[i] = make([]int, size)
		for j := range m[i] {
			m[i][j] = int(r.Float64() * k)
		}
	}
	return m
}

func DirToUndir(dirM [][]int) [][]int {
	size := len(dirM)
	undirM := make([][]int, size)
	for i := range undirM {
		undirM[i] = make([]int, size)
		for j := range i + 1 {
			if dirM[i][j] == 1 || dirM[j][i] == 1 {
				undirM[i][j] = 1
				undirM[j][i] = 1
			}
		}
	}
	return undirM
}

func composeM[N number](m1, m2 [][]N) [][]N {
	mc := make([][]N, len(m1))
	for row1 := range m1 {
		mc[row1] = make([]N, len(m1))
		for col1 := range m1[row1] {
			for col2 := range m2[col1] {
				mc[row1][col2] += m1[row1][col1] * m2[col1][col2]
			}
		}
	}
	return mc
}

func AddM[N number](sum, adder [][]N) {
	for row, sumRow := range sum {
		for col := range sumRow {
			sumRow[col] += adder[row][col]
		}
	}
}

func MultM[N number](mult, multer [][]N) {
	for row, multRow := range mult {
		for col := range multRow {
			multRow[col] *= multer[row][col]
		}
	}
}

func ScalarM[N number](mult [][]N, multer N) {
	for _, multRow := range mult {
		for col := range multRow {
			multRow[col] *= multer
		}
	}
}

func transM[T any](m [][]T) [][]T {
	size := len(m)
	t := make([][]T, size)
	for i := range t {
		t[i] = make([]T, size)
		for j := range t[i] {
			t[i][j] = m[j][i]
		}
	}
	return t
}

func ToBinM[N number](m [][]N) [][]int {
	size := len(m)
	bin := make([][]int, size)
	for i := range bin {
		bin[i] = make([]int, size)
		for j := range bin[i] {
			if m[i][j] != 0 {
				bin[i][j] = 1
			}
		}
	}
	return bin
}

func identM(size int) [][]int {
	ident := make([][]int, size)
	for i := range ident {
		ident[i] = make([]int, size)
		ident[i][i] = 1
	}
	return ident
}

func copyM[T any](dst, src [][]T) int {
	copyCount := 0
	minSize := len(dst)
	if minSize > len(src) {
		minSize = len(src)
	}
	for i := 0; i < minSize; i++ {
		dst[i] = make([]T, len(src[i]))
		copyCount += copy(dst[i], src[i])
	}
	return copyCount
}
