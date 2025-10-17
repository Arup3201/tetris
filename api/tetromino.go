package api

import "log"

const (
	SHAPE_T = "T"
)

func newTetromino(shape, color string, playgroundDimension [2]int) tetromino {
	switch shape {
	case SHAPE_T:
		return newTTetromino(color, playgroundDimension)
	}

	log.Fatalf("unsupported tetromino shape %s", shape)
	return nil
}

type tTetromino struct {
	color               string
	coordinates         [4][2]int
	playgroundDimension [2]int // [width, height]
}

func newTTetromino(color string, playgroundDimension [2]int) *tTetromino {
	return &tTetromino{
		color: color,
		coordinates: [4][2]int{
			{-1, -1},
			{-1, -1},
			{-1, -1},
			{-1, -1},
		},
		playgroundDimension: playgroundDimension,
	}
}

func (t *tTetromino) getPosition() [4][2]int {
	return t.coordinates
}

func (t *tTetromino) drop() {
	// consider the walls
	t.coordinates[0][0] = 0
	t.coordinates[0][1] = t.playgroundDimension[0] / 2
}
