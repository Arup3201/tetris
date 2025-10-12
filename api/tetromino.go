package api

import "log"

const (
	SHAPE_T = "T"
)

func newTetromino(shape, color string) tetromino {
	switch shape {
	case SHAPE_T:
		return newTTetromino(color)
	}

	log.Fatalf("unsupported tetromino shape %s", shape)
	return nil
}

type tTetromino struct {
	color       string
	coordinates [4][2]int
}

func newTTetromino(color string) *tTetromino {
	return &tTetromino{
		color: color,
		coordinates: [4][2]int{
			{-1, -1},
			{-1, -1},
			{-1, -1},
			{-1, -1},
		},
	}
}

func (t *tTetromino) getPosition() [4][2]int {
	return t.coordinates
}

func (t *tTetromino) drop() {
	t.coordinates[0][0] = 0
	t.coordinates[0][1] = 5
}
