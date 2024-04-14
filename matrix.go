package graph

import "math/rand"

type number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func RandAdjDirMatrix(size int, k float64, r *rand.Rand) [][]uint8 {
	dirM := make([][]uint8, size)
	for i := range dirM {
		dirM[i] = make([]uint8, size)
		for j := range dirM[i] {
			dirM[i][j] = uint8(2 * r.Float64() * k)
		}
	}
	return dirM
}

func DirToUndir(dirM [][]uint8) [][]uint8 {
	size := len(dirM)
	undirM := make([][]uint8, size)
	for i := range undirM {
		undirM[i] = make([]uint8, size)
		for j := range i {
			if dirM[i][j] == 1 || dirM[j][i] == 1 {
				undirM[i][j] = 1
				undirM[j][i] = 1
			}
		}
	}
	return undirM
}

func composeM(m1, m2 [][]uint8) [][]uint8 {
	mc := make([][]uint8, len(m1))
	for row1 := range m1 {
		mc[row1] = make([]uint8, len(m1))
		for col1 := range m1[row1] {
			for col2 := range m2[col1] {
				mc[row1][col2] += m1[row1][col1] * m2[col1][col2]
			}
		}
	}
	return mc
}

func addM[N number](sum, adder [][]N) {
	for row, sumRow := range sum {
		for col := range sumRow {
			sumRow[col] += adder[row][col]
		}
	}
}

func multM[N number](mult, multer [][]N) {
	for row, multRow := range mult {
		for col := range multRow {
			multRow[col] *= multer[row][col]
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

func toBinM[N number](m [][]N) [][]uint8 {
	size := len(m)
	bin := make([][]uint8, size)
	for i := range bin {
		bin[i] = make([]uint8, size)
		for j := range bin[i] {
			if m[i][j] != 0 {
				bin[i][j] = 1
			}
		}
	}
	return bin
}

func identM[N uint8](size int) [][]uint8 {
	ident := make([][]uint8, size)
	for i := range ident {
		ident[i] = make([]uint8, size)
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
