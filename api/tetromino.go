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
	coordinates         [4][2]int // [[s0.row, s0.column], [s1.row, s1.column], [s2.row, s2.column], [s3.row, s3.column]]
	playgroundDimension [2]int    // [height, width]
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
	if t.coordinates[0][0] == -1 {
		t.coordinates[0][0] = 0
		t.coordinates[0][1] = t.playgroundDimension[1] / 2
	}
}

func (t *tTetromino) fallByOne() {
	if t.coordinates[0][0] == -1 {
		log.Fatal("Tetromino piece has not yet been dropped inside the playground")
	}

	// tetromino touched the ground
	if t.coordinates[0][0] >= t.playgroundDimension[0]-1 {
		return
	}

	t.coordinates[0][0]++

	if t.coordinates[1][0] == -1 {
		t.coordinates[1][0] = 0
		t.coordinates[1][1] = t.playgroundDimension[1] / 2
	} else {
		t.coordinates[1][0]++
	}

	if t.coordinates[2][0] == -1 {
		t.coordinates[2][0] = 0
		t.coordinates[2][1] = t.coordinates[1][1] + 1
	} else {
		t.coordinates[2][0]++
	}

	if t.coordinates[3][0] == -1 {
		t.coordinates[3][0] = 0
		t.coordinates[3][1] = t.coordinates[1][1] - 1
	} else {
		t.coordinates[3][0]++
	}
}
