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
	assert.Equal(t, 12, game.columns)
	assert.Equal(t, 22, game.rows)
	for c := range game.columns {
		assert.Equal(t, WALL_ID, game.grid[0][c])
		assert.Equal(t, WALL_ID, game.grid[game.rows-1][c])
	}
	for r := range game.rows {
		assert.Equal(t, WALL_ID, game.grid[r][0])
		assert.Equal(t, WALL_ID, game.grid[r][game.columns-1])
	}
	// playground should be empty
	for r := range game.rows - 2 {
		for c := range game.columns - 2 {
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
	for c := range game.columns {
		expectedWallCoordinates = append(expectedWallCoordinates, [2]int{0, c})
		expectedWallCoordinates = append(expectedWallCoordinates, [2]int{game.rows - 1, c})
	}
	for r := range game.rows {
		expectedWallCoordinates = append(expectedWallCoordinates, [2]int{r, 0})
		expectedWallCoordinates = append(expectedWallCoordinates, [2]int{r, game.columns - 1})
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
	for r := range game.rows - 2 {
		for c := range game.columns - 2 {
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
	got := game.GetTetrominoPosition()
	assert.Equal(t, expectedTetromino["shape"], SHAPE_T)
	assert.Equal(t, expectedTetromino["color"], COLOR_YELLOW)
	assert.Equal(t, expectedTetromino["position"], got)
	assert.Equal(t, true, game.HasTetrominoDropped)
}

func TestTetrisGoDown(t *testing.T) {
	// prepare
	rows, columns := 20, 10
	game := CreateGame(rows, columns)
	tetrominoColor, tetrominoShape := COLOR_YELLOW, SHAPE_T
	game.DropTetromino(tetrominoShape, tetrominoColor)

	// act
	game.TetrominoFallsByOne()

	// assert
	expectedTetromino := map[string]any{
		"position": [4][2]int{
			{2, 6},
			{1, 6},
			{1, 7},
			{1, 5},
		},
	}
	got := game.GetTetrominoPosition()
	assert.Equal(t, expectedTetromino["position"], got)
}

func TestTetrisGoDownBy10(t *testing.T) {
	// prepare
	rows, columns := 20, 10
	game := CreateGame(rows, columns)
	tetrominoColor, tetrominoShape := COLOR_YELLOW, SHAPE_T
	game.DropTetromino(tetrominoShape, tetrominoColor)

	// act
	for range 10 {
		game.TetrominoFallsByOne()
	}

	// assert
	expectedTetromino := map[string]any{
		"position": [4][2]int{
			{11, 6},
			{10, 6},
			{10, 7},
			{10, 5},
		},
	}
	got := game.GetTetrominoPosition()
	assert.Equal(t, expectedTetromino["position"], got)
}

func TestTetrisHitGround(t *testing.T) {
	// prepare
	rows, columns := 20, 10
	game := CreateGame(rows, columns)
	tetrominoColor, tetrominoShape := COLOR_YELLOW, SHAPE_T
	game.DropTetromino(tetrominoShape, tetrominoColor)

	// act
	for range 20 {
		game.TetrominoFallsByOne()
	}

	// assert
	assert.Equal(t, false, game.HasTetrominoDropped)
	assert.Equal(t, SQUARE_ID, game.grid[20][6])
	assert.Equal(t, SQUARE_ID, game.grid[19][6])
	assert.Equal(t, SQUARE_ID, game.grid[19][7])
	assert.Equal(t, SQUARE_ID, game.grid[19][5])
}

func TestTetrisHitAnotherTetris(t *testing.T) {
	// prepare
	rows, columns := 20, 10
	game := CreateGame(rows, columns)
	tetrominoColor, tetrominoShape := COLOR_YELLOW, SHAPE_T
	game.DropTetromino(tetrominoShape, tetrominoColor)
	for range 20 {
		game.TetrominoFallsByOne()
	}
	game.DropTetromino(tetrominoShape, tetrominoColor)

	// act
	for range 20 {
		game.TetrominoFallsByOne()
	}

	// assert
	assert.Equal(t, false, game.HasTetrominoDropped)
	assert.Equal(t, SQUARE_ID, game.grid[18][6])
	assert.Equal(t, SQUARE_ID, game.grid[17][6])
	assert.Equal(t, SQUARE_ID, game.grid[17][7])
	assert.Equal(t, SQUARE_ID, game.grid[17][5])
}
