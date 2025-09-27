package api

type coord struct {
	X, Y int
}

type tetrominoCoord struct {
	S1, S2, S3, S4 coord
}

type tetromino struct {
	coord                 tetrominoCoord
	gridRows, gridColumns int
}

func CreateTetromino(bottom, right int) *tetromino {
	return &tetromino{
		coord: tetrominoCoord{
			S1: coord{
				X: -1,
				Y: -1,
			},
			S2: coord{
				X: -1,
				Y: -1,
			},
			S3: coord{
				X: -1,
				Y: -1,
			},
			S4: coord{
				X: -1,
				Y: -1,
			},
		},
		gridRows:    bottom,
		gridColumns: right,
	}
}

func (t *tetromino) EnterField() {
	t.coord.S1.X = t.gridColumns/2 - 1
	t.coord.S1.Y = 0
	t.coord.S2.X = t.gridColumns / 2
	t.coord.S2.Y = 0
}

func (t *tetromino) GetPosition() []coord {
	return []coord{t.coord.S1, t.coord.S2, t.coord.S3, t.coord.S4}
}

func (t *tetromino) GoDown(fieldGrid [][]int, by int) {
	lastPossibleY := t.gridRows - 1
	for r := range t.gridRows {
		if fieldGrid[r][t.coord.S1.X] == 1 || fieldGrid[r][t.coord.S2.X] == 1 {
			lastPossibleY = r - 1
			break
		}
	}

	t.coord.S1.Y = min(t.coord.S1.Y+by, lastPossibleY)
	t.coord.S2.Y = min(t.coord.S2.Y+by, lastPossibleY)

	if t.coord.S3.X == -1 && t.coord.S3.Y == -1 {
		t.coord.S3.X = t.coord.S2.X
		t.coord.S3.Y = t.coord.S2.Y - 1
	}
	if t.coord.S4.X == -1 && t.coord.S4.Y == -1 {
		t.coord.S4.X = t.coord.S1.X
		t.coord.S4.Y = t.coord.S1.Y - 1
	}
}
