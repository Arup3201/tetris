package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlankField(t *testing.T) {
	rows, columns := 20, 10

	field := CreateField(rows, columns)

	assert.Equal(t, 0, field.grid[0][0])
	assert.Equal(t, 0, field.grid[rows-1][columns-1])
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
	field := CreateField(rows, columns)
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
	field := CreateField(rows, columns)
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
	field := CreateField(rows, columns)
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
