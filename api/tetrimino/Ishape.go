package tetrimino

import "github.com/Arup3201/tetris/api/matrix"

type MatrixSetFunc func(r, c int)

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

func (t *ITetrimino) Spawn(matrix *matrix.ReverseMatrix) {
	matrix.Set(t.top-1, t.left, t.name)
	matrix.Set(t.top-1, t.left+1, t.name)
	matrix.Set(t.top-1, t.left+2, t.name)
	matrix.Set(t.top-1, t.left+3, t.name)
}
