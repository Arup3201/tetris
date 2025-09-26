package tetris

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlankField(t *testing.T) {
	rows, columns := 20, 10

	field := CreateField(rows, columns)

	assert.Equal(t, field[0][0], 0)
	assert.Equal(t, field[rows-1][columns-1], 0)
}

func TestTetrominoEnterFieldAtTopCenter(t *testing.T) {
	rows, columns := 20, 10
	tetromino := CreateTetromino(rows, columns)

	tetromino.EnterField()

	x, y := tetromino.GetPosition()
	assert.Equal(t, x, columns/2)
	assert.Equal(t, y, 0)
}

func TestTetrominoGoDown(t *testing.T) {
	rows, columns := 20, 10
	tetromino := CreateTetromino(rows, columns)
	tetromino.EnterField()

	tetromino.GoDown(10)

	x, y := tetromino.GetPosition()
	assert.Equal(t, x, columns/2)
	assert.Equal(t, y, 10)
}
