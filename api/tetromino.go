package api

type Tetromino struct {
	squares               [4]*Square
	gridRows, gridColumns int
}

func CreateTetromino(playgroundGridRows, playgroundGridColumns int) *Tetromino {
	squares := [4]*Square{
		createSquare(-1, -1, COLOR_ORANGE),
		createSquare(-1, -1, COLOR_ORANGE),
		createSquare(-1, -1, COLOR_ORANGE),
		createSquare(-1, -1, COLOR_ORANGE),
	}

	return &Tetromino{
		squares:     squares,
		gridRows:    playgroundGridRows,
		gridColumns: playgroundGridColumns,
	}
}

func (t *Tetromino) EnterField() {
	t.squares[0].changePosition(t.gridColumns/2-1, 0)
	t.squares[1].changePosition(t.gridColumns/2, 0)
}

func (t *Tetromino) DidNotEnterField() bool {
	return t.squares[0].GetY() == -1 && t.squares[1].GetY() == -1
}

func (t *Tetromino) GetPosition() []coord {
	return []coord{t.squares[0].position, t.squares[1].position, t.squares[2].position, t.squares[3].position}
}

func (t *Tetromino) GoDown(grid [][]*Square, by int) {
	lastPossibleY := t.gridRows - 1
	for r := range t.gridRows {
		if grid[r][t.squares[0].GetX()] != nil || grid[r][t.squares[1].GetX()] != nil {
			lastPossibleY = r - 1
			break
		}
	}

	t.squares[0].changePosition(t.squares[0].GetX(), min(t.squares[0].GetY()+by, lastPossibleY))
	t.squares[1].changePosition(t.squares[1].GetX(), min(t.squares[1].GetY()+by, lastPossibleY))
	t.squares[2].changePosition(t.squares[1].GetX(), t.squares[1].GetY()-1)
	t.squares[3].changePosition(t.squares[0].GetX(), t.squares[0].GetY()-1)
}

func (t *Tetromino) HasHit(fieldGrid [][]*Square) bool {
	if t.squares[0].GetY() == t.gridRows-1 {
		return true
	}
	if t.squares[1].GetY() == t.gridRows-1 {
		return true
	}

	if fieldGrid[t.squares[0].GetY()+1][t.squares[0].GetX()] != nil || fieldGrid[t.squares[1].GetY()+1][t.squares[0].GetX()] != nil {
		return true
	}

	return false
}
