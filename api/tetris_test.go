package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGamePlayground(t *testing.T) {
	rows, columns := 20, 10

	game := CreateGame(rows, columns)

	assert.Equal(t, 0, game.playground.grid[0][0])
	assert.Equal(t, 0, game.playground.grid[rows-1][columns-1])
}

func TestWallBorder(t *testing.T) {
	rows, colums := 4, 4

	game := CreateGame(rows, colums)

	expected := []coord{
		{
			X: 0,
			Y: 0,
		},
		{
			X: 1,
			Y: 0,
		},
		{
			X: 2,
			Y: 0,
		},
		{
			X: 3,
			Y: 0,
		},
		{
			X: 4,
			Y: 0,
		},
		{
			X: 5,
			Y: 0,
		},
		{
			X: 0,
			Y: 1,
		},
		{
			X: 5,
			Y: 1,
		},
		{
			X: 0,
			Y: 2,
		},
		{
			X: 5,
			Y: 2,
		},
		{
			X: 0,
			Y: 3,
		},
		{
			X: 5,
			Y: 3,
		},
		{
			X: 0,
			Y: 4,
		},
		{
			X: 5,
			Y: 4,
		},
		{
			X: 0,
			Y: 5,
		},
		{
			X: 1,
			Y: 5,
		},
		{
			X: 2,
			Y: 5,
		},
		{
			X: 3,
			Y: 5,
		},
		{
			X: 4,
			Y: 5,
		},
		{
			X: 5,
			Y: 5,
		},
	}
	border := game.wall.GetBorder()
	assert.Equal(t, expected, border)
}

func TestTetrominoEnterFieldAtTopCenter(t *testing.T) {
	// prepare
	rows, columns := 20, 10
	tetromino := CreateTetromino(rows, columns)

	// act
	tetromino.EnterField()

	// assert
	positions := tetromino.GetPosition()
	assert.Equal(t, 4, positions[0].X)
	assert.Equal(t, 0, positions[0].Y)
	assert.Equal(t, 5, positions[1].X)
	assert.Equal(t, 0, positions[1].Y)
	assert.Equal(t, -1, positions[2].X)
	assert.Equal(t, -1, positions[2].Y)
	assert.Equal(t, -1, positions[3].X)
	assert.Equal(t, -1, positions[3].Y)
}

func TestTetrominoGoDown(t *testing.T) {
	// prepare
	rows, columns := 20, 10
	field := CreatePlayground(rows, columns)
	tetromino := CreateTetromino(rows, columns)
	tetromino.EnterField()

	// act
	tetromino.GoDown(field.grid, 10)

	// assert
	positions := tetromino.GetPosition()
	assert.Equal(t, 4, positions[0].X)
	assert.Equal(t, 10, positions[0].Y)
	assert.Equal(t, 4, positions[3].X)
	assert.Equal(t, 9, positions[3].Y)
	assert.Equal(t, 5, positions[1].X)
	assert.Equal(t, 10, positions[1].Y)
	assert.Equal(t, 5, positions[2].X)
	assert.Equal(t, 9, positions[2].Y)
}

func TestTetrominoHitGround(t *testing.T) {
	// prepare
	rows, columns := 20, 10
	field := CreatePlayground(rows, columns)
	tetromino := CreateTetromino(rows, columns)
	tetromino.EnterField()

	// act
	tetromino.GoDown(field.grid, 20)

	// assert
	positions := tetromino.GetPosition()
	assert.Equal(t, 4, positions[0].X)
	assert.Equal(t, rows-1, positions[0].Y)
	assert.Equal(t, 4, positions[3].X)
	assert.Equal(t, rows-2, positions[3].Y)
	assert.Equal(t, 5, positions[1].X)
	assert.Equal(t, rows-1, positions[1].Y)
	assert.Equal(t, 5, positions[2].X)
	assert.Equal(t, rows-2, positions[2].Y)
}

func TestTetrominoHitAnotherTetromino(t *testing.T) {
	// prepare
	rows, columns := 20, 10
	field := CreatePlayground(rows, columns)
	field.SetGround([]tetrominoCoord{
		{
			s1: coord{
				X: 4,
				Y: 19,
			},
			s2: coord{
				X: 5,
				Y: 19,
			},
			s3: coord{
				X: 5,
				Y: 18,
			},
			s4: coord{
				X: 4,
				Y: 18,
			},
		},
	})
	tetromino := CreateTetromino(rows, columns)
	tetromino.EnterField()

	// act
	tetromino.GoDown(field.grid, 20)

	// assert
	positions := tetromino.GetPosition()
	assert.Equal(t, 4, positions[0].X)
	assert.Equal(t, 17, positions[0].Y)
	assert.Equal(t, 4, positions[3].X)
	assert.Equal(t, 16, positions[3].Y)
	assert.Equal(t, 5, positions[1].X)
	assert.Equal(t, 17, positions[1].Y)
	assert.Equal(t, 5, positions[2].X)
	assert.Equal(t, 16, positions[2].Y)
}
