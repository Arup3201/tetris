package api

type coord struct {
	X, Y int
}

type square struct {
	position coord
	color    string
}

func createSquare(x, y int, color string) *square {
	return &square{
		position: coord{
			X: x,
			Y: y,
		},
		color: color,
	}
}

func (sq *square) changePosition(x, y int) {
	sq.position.X = x
	sq.position.Y = y
}

func (sq *square) getX() int {
	return sq.position.X
}

func (sq *square) getY() int {
	return sq.position.Y
}

type Tetromino struct {
	squares               [4]*square
	gridRows, gridColumns int
}

func CreateTetromino(playgroundGridRows, playgroundGridColumns int) *Tetromino {
	squares := [4]*square{
		createSquare(-1, -1, "gray"),
		createSquare(-1, -1, "gray"),
		createSquare(-1, -1, "gray"),
		createSquare(-1, -1, "gray"),
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
	return t.squares[0].getY() == -1 && t.squares[1].getY() == -1
}

func (t *Tetromino) GetPosition() []coord {
	return []coord{t.squares[0].position, t.squares[1].position, t.squares[2].position, t.squares[3].position}
}

func (t *Tetromino) GoDown(fieldGrid [][]int, by int) {
	lastPossibleY := t.gridRows - 1
	for r := range t.gridRows {
		if fieldGrid[r][t.squares[0].getX()] == 1 || fieldGrid[r][t.squares[1].getX()] == 1 {
			lastPossibleY = r - 1
			break
		}
	}

	t.squares[0].changePosition(t.squares[0].getX(), min(t.squares[0].getY()+by, lastPossibleY))
	t.squares[1].changePosition(t.squares[1].getX(), min(t.squares[1].getY()+by, lastPossibleY))
	t.squares[2].changePosition(t.squares[1].getX(), t.squares[1].getY()-1)
	t.squares[3].changePosition(t.squares[0].getX(), t.squares[0].getY()-1)
}

func (t *Tetromino) HasHit(fieldGrid [][]int) bool {
	if t.squares[0].getY() == t.gridRows-1 {
		return true
	}
	if t.squares[1].getY() == t.gridRows-1 {
		return true
	}

	if fieldGrid[t.squares[0].getY()+1][t.squares[0].getX()] == 1 || fieldGrid[t.squares[1].getY()+1][t.squares[0].getX()] == 1 {
		return true
	}

	return false
}
