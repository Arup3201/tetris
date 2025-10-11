package api

import (
	"slices"
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

func TestGetWall(t *testing.T) {
	// prepare
	rows, columns := 20, 10
	game := CreateGame(rows, columns)

	// act
	wall := game.GetWallCoordinates()

	// assert
	expectedWallCoordinates := [][]int{}
	for c := range game.width {
		expectedWallCoordinates = append(expectedWallCoordinates, []int{0, c})
		expectedWallCoordinates = append(expectedWallCoordinates, []int{game.height - 1, c})
	}
	for r := range game.height {
		expectedWallCoordinates = append(expectedWallCoordinates, []int{r, 0})
		expectedWallCoordinates = append(expectedWallCoordinates, []int{r, game.width - 1})
	}
	included := false
	for _, want := range expectedWallCoordinates {
		included = false
		for _, got := range wall {
			if slices.Equal(want, got) {
				included = true
			}
		}
		if !included {
			t.Errorf("expected %v to be included in the wall coordinates %v", want, wall)
		}
	}
}
