package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameStart(t *testing.T) {
	// prepare
	// game grid where user can play
	rows, columns := 20, 10

	// act
	game := CreateGame(rows, columns)

	// assert
	// check wall
	assert.Equal(t, 12, game.width)
	assert.Equal(t, 22, game.height)
	for c := range game.width {
		assert.Equal(t, "wall", game.grid[0][c])
		assert.Equal(t, "wall", game.grid[game.height-1][c])
	}
	for r := range game.height {
		assert.Equal(t, "wall", game.grid[r][0])
		assert.Equal(t, "wall", game.grid[r][game.width-1])
	}
	// playground should be empty
	for r := range game.height - 2 {
		for c := range game.width - 2 {
			assert.Equal(t, "blank", game.grid[r+1][c+1])
		}
	}
	// score should be 0
	assert.Equal(t, 0, game.score)
}
