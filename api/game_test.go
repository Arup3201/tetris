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
		assert.Equal(t, WALL_ID, game.grid[0][c])
		assert.Equal(t, WALL_ID, game.grid[game.height-1][c])
	}
	for r := range game.height {
		assert.Equal(t, WALL_ID, game.grid[r][0])
		assert.Equal(t, WALL_ID, game.grid[r][game.width-1])
	}
	// playground should be empty
	for r := range game.height - 2 {
		for c := range game.width - 2 {
			assert.Equal(t, BLANK_ID, game.grid[r+1][c+1])
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
	expectedWallCoordinates := [][2]int{}
	for c := range game.width {
		expectedWallCoordinates = append(expectedWallCoordinates, [2]int{0, c})
		expectedWallCoordinates = append(expectedWallCoordinates, [2]int{game.height - 1, c})
	}
	for r := range game.height {
		expectedWallCoordinates = append(expectedWallCoordinates, [2]int{r, 0})
		expectedWallCoordinates = append(expectedWallCoordinates, [2]int{r, game.width - 1})
	}
	for _, want := range expectedWallCoordinates {
		assert.Equal(t, true, slices.Contains(wall, want))
	}
}

func TestGetPlayground(t *testing.T) {
	// prepare
	rows, columns := 20, 10
	game := CreateGame(rows, columns)

	// act
	playground := game.GetPlayground()

	// assert
	expectedPlayground := map[[2]int]string{}
	for r := range game.height - 2 {
		for c := range game.width - 2 {
			expectedPlayground[[2]int{r + 1, c + 1}] = BLANK_ID
		}
	}
	assert.Equal(t, expectedPlayground, playground)
}

func TestTetrisDropInsidePlayground(t *testing.T) {
	// prepare
	rows, columns := 20, 10
	game := CreateGame(rows, columns)
	tetrominoColor, tetrominoShape := COLOR_YELLOW, SHAPE_T

	// act
	game.DropTetromino(tetrominoShape, tetrominoColor)

	// assert
	expectedTetromino := map[string]any{
		"color": COLOR_YELLOW,
		"shape": SHAPE_T,
		"position": [4][2]int{
			{1, 6},
			{-1, -1},
			{-1, -1},
			{-1, -1},
		},
	}
	got := game.droppingTetromino.getPosition()
	assert.Equal(t, expectedTetromino["shape"], SHAPE_T)
	assert.Equal(t, expectedTetromino["color"], COLOR_YELLOW)
	assert.Equal(t, expectedTetromino["position"], got)
	assert.Equal(t, true, game.HasTetrominoDropped)
}
