package tetrimino

import "github.com/Arup3201/tetris/api/matrix"

type ITetrimino struct {
	boundingBox [4][4]uint8
	name        string
	top, left   int
}

func CreateITetrimino(name string) *ITetrimino {
	return &ITetrimino{
		boundingBox: [4][4]uint8{
			{0, 0, 0, 0},
			{1, 1, 1, 1},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		},
		name: name,
		top:  22,
		left: 4,
	}
}

func (t *ITetrimino) Spawn(matrix *matrix.ReverseMatrix, emptyCellValue string) bool {
	x1, _ := matrix.Get(t.top-1, t.left)
	x2, _ := matrix.Get(t.top-1, t.left+1)
	x3, _ := matrix.Get(t.top-1, t.left+2)
	x4, _ := matrix.Get(t.top-1, t.left+3)
	if x1 != emptyCellValue || x2 != emptyCellValue || x3 != emptyCellValue || x4 != emptyCellValue {
		return false
	}

	matrix.Set(t.top-1, t.left, t.name)
	matrix.Set(t.top-1, t.left+1, t.name)
	matrix.Set(t.top-1, t.left+2, t.name)
	matrix.Set(t.top-1, t.left+3, t.name)
	return true
}
