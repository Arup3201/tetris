package api

import (
	"github.com/Arup3201/tetris/api/matrix"
	"github.com/Arup3201/tetris/api/tetrimino"
)

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
	Spawn(*matrix.ReverseMatrix)
}

type Game struct {
	rows, columns int
	playfield     *matrix.ReverseMatrix
	score         int
	spawned       TetriminoInterface
}

func CreateGame() *Game {
	totalRows, totalColumns := 22, 10

	return &Game{
		rows:      totalRows,
		columns:   totalColumns,
		playfield: matrix.CreateReverseMatrix(totalRows, totalColumns, BLANK_ID),
		score:     0,
	}
}

func (g *Game) SetPlayfield(row, column int, value string) {
	ok := g.playfield.Set(row, column, value)
	if !ok {
		panic("error setting value to game playground")
	}
}

func (g *Game) GetPlayfield(row, column int) string {
	got, ok := g.playfield.Get(row, column)

	if !ok {
		panic("error getting value to game playground")
	}

	return got
}

func (g *Game) SpawnTetrimino(shape string) {
	g.spawned = createTetrimino(shape)
	g.spawned.Spawn(g.playfield)
}
