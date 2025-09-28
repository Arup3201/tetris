package api

type coord struct {
	X, Y int
}

type tetrominoCoord struct {
	s1, s2, s3, s4 coord
}

type tetromino struct {
	coord                 tetrominoCoord
	gridRows, gridColumns int
}

func CreateTetromino(playgroundGridRows, playgroundGridColumns int) *tetromino {
	return &tetromino{
		coord: tetrominoCoord{
			s1: coord{
				X: -1,
				Y: -1,
			},
			s2: coord{
				X: -1,
				Y: -1,
			},
			s3: coord{
				X: -1,
				Y: -1,
			},
			s4: coord{
				X: -1,
				Y: -1,
			},
		},
		gridRows:    playgroundGridRows,
		gridColumns: playgroundGridColumns,
	}
}

func (t *tetromino) EnterField() {
	t.coord.s1.X = t.gridColumns/2 - 1
	t.coord.s1.Y = 0
	t.coord.s2.X = t.gridColumns / 2
	t.coord.s2.Y = 0
}

func (t *tetromino) GetPosition() []coord {
	return []coord{t.coord.s1, t.coord.s2, t.coord.s3, t.coord.s4}
}

func (t *tetromino) GoDown(fieldGrid [][]int, by int) {
	lastPossibleY := t.gridRows - 1
	for r := range t.gridRows {
		if fieldGrid[r][t.coord.s1.X] == 1 || fieldGrid[r][t.coord.s2.X] == 1 {
			lastPossibleY = r - 1
			break
		}
	}

	t.coord.s1.Y = min(t.coord.s1.Y+by, lastPossibleY)
	t.coord.s2.Y = min(t.coord.s2.Y+by, lastPossibleY)

	if t.coord.s3.X == -1 && t.coord.s3.Y == -1 {
		t.coord.s3.X = t.coord.s2.X
		t.coord.s3.Y = t.coord.s2.Y - 1
	}
	if t.coord.s4.X == -1 && t.coord.s4.Y == -1 {
		t.coord.s4.X = t.coord.s1.X
		t.coord.s4.Y = t.coord.s1.Y - 1
	}
}

func (t *tetromino) HasHit(fieldGrid [][]int) bool {
	if t.coord.s1.Y == t.gridRows-1 {
		return true
	}
	if t.coord.s2.Y == t.gridRows-1 {
		return true
	}

	if fieldGrid[t.coord.s1.Y+1][t.coord.s1.X] == 1 || fieldGrid[t.coord.s2.Y+1][t.coord.s2.X] == 1 {
		return true
	}

	return false
}
