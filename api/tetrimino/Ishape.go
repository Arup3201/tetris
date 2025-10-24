package tetrimino

import "github.com/Arup3201/tetris/api/matrix"

type MatrixSetFunc func(r, c int)

type ITetrimino struct {
	BoundingBox [4][4]uint8
	Name        string
}

func CreateITetrimino(name string) *ITetrimino {
	return &ITetrimino{
		BoundingBox: [4][4]uint8{
			{0, 0, 0, 0},
			{1, 1, 1, 1},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		},
		Name: name,
	}
}

func (t *ITetrimino) Spawn(matrix *matrix.ReverseMatrix) {
	matrix.Set(21, 4, t.Name)
	matrix.Set(21, 5, t.Name)
	matrix.Set(21, 6, t.Name)
	matrix.Set(21, 7, t.Name)
}
