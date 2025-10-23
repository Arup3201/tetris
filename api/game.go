package api

import "github.com/Arup3201/tetris/api/tetrimino"

const (
	BLANK_ID = "blank"
	SHAPE_I  = "I"
)

func createTetrimino(shape string) TetriminoInterface {
	switch shape {
	case SHAPE_I:
		return tetrimino.CreateITetrimino(SHAPE_I)
	}

	return nil
}

type TetriminoInterface interface {
	Spawn([][]string)
}

type Game struct {
	rows, columns int
	playfield     [][]string
	score         int
	spawned       TetriminoInterface
}

func CreateGame() *Game {
	totalRows, totalColumns := 20, 10

	grid := make([][]string, totalRows)
	for r := range totalRows {
		grid[r] = make([]string, totalColumns)
	}

	for r := range totalRows {
		for c := range totalColumns {
			grid[r][c] = BLANK_ID
		}
	}

	return &Game{
		rows:      totalRows,
		columns:   totalColumns,
		playfield: grid,
		score:     0,
	}
}

func (g *Game) SpawnTetrimino(shape string) {
	g.spawned = createTetrimino(shape)
	g.spawned.Spawn(g.playfield)
}
