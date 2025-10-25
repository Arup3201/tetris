package api

import (
	"github.com/Arup3201/tetris/api/matrix"
	"github.com/Arup3201/tetris/api/tetrimino"
)

const (
	BLANK_ID       = "blank"
	SHAPE_I        = "I"
	SHAPE_J        = "J"
	SHAPE_L        = "L"
	SHAPE_O        = "O"
	SHAPE_S        = "S"
	SHAPE_Z        = "Z"
	SHAPE_T        = "T"
	MATRIX_ROWS    = 22
	MATRIX_COLUMNS = 10
)

func createTetrimino(shape string) TetriminoInterface {
	switch shape {
	case SHAPE_I:
		return tetrimino.CreateITetrimino(SHAPE_I)
	}

	return nil
}

type TetriminoInterface interface {
	Spawn(*matrix.ReverseMatrix, string) bool
	DropByOne(*matrix.ReverseMatrix, string) bool
}

type Game struct {
	rows, columns int
	playfield     *matrix.ReverseMatrix
	score         int
	spawned       TetriminoInterface
}

func CreateGame() *Game {
	return &Game{
		rows:      MATRIX_ROWS,
		columns:   MATRIX_COLUMNS,
		playfield: matrix.CreateReverseMatrix(MATRIX_ROWS, MATRIX_COLUMNS, BLANK_ID),
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

func (g *Game) SpawnTetrimino(shape string) bool {
	g.spawned = createTetrimino(shape)
	if ok := g.spawned.Spawn(g.playfield, BLANK_ID); !ok {
		g.spawned = nil
		return false
	}

	return true
}

func (g *Game) HasSpawned() bool {
	return g.spawned != nil
}

func (g *Game) DropByOne() bool {
	droppable := g.spawned.DropByOne(g.playfield, BLANK_ID)

	if !droppable {
		g.spawned = nil
	}

	return droppable
}
