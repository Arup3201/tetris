package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	playground := CreatePlayground(rows, columns)
	tetromino := CreateTetromino(rows, columns)
	tetromino.EnterField()

	// act
	tetromino.GoDown(playground.grid, 10)

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
	playground := CreatePlayground(rows, columns)
	tetromino := CreateTetromino(rows, columns)
	tetromino.EnterField()

	// act
	tetromino.GoDown(playground.grid, 20)

	// assert
	expected := true
	got := tetromino.HasHit(playground.grid)
	assert.Equal(t, expected, got)

	expectedPositions := [][]int{{4, rows - 1}, {5, rows - 1}, {5, rows - 2}, {4, rows - 2}}
	positions := tetromino.GetPosition()
	gotPositions := [][]int{{positions[0].X, positions[0].Y}, {positions[1].X, positions[1].Y}, {positions[2].X, positions[2].Y}, {positions[3].X, positions[3].Y}}
	assert.Equal(t, expectedPositions, gotPositions)
}

func TestTetrominoHitAnotherTetromino(t *testing.T) {
	// prepare
	rows, columns := 20, 10
	playground := CreatePlayground(rows, columns)
	playground.SetGround([]*square{
		createSquare(4, 19, "gray"),
		createSquare(5, 19, "gray"),
		createSquare(5, 18, "gray"),
		createSquare(4, 18, "gray"),
	})
	tetromino := CreateTetromino(rows, columns)
	tetromino.EnterField()

	// act
	tetromino.GoDown(playground.grid, 20)

	// assert
	expected := true
	got := tetromino.HasHit(playground.grid)
	assert.Equal(t, expected, got)

	expectedPositions := [][]int{{4, 17}, {5, 17}, {5, 16}, {4, 16}}
	positions := tetromino.GetPosition()
	gotPositions := [][]int{{positions[0].X, positions[0].Y}, {positions[1].X, positions[1].Y}, {positions[2].X, positions[2].Y}, {positions[3].X, positions[3].Y}}
	assert.Equal(t, expectedPositions, gotPositions)
}
