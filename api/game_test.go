package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameStart(t *testing.T) {
	t.Run("when game starts the score is 0 and matrix is blank", func(t *testing.T) {
		// prepare

		// act
		game := CreateGame()

		// assert
		// check wall
		assert.Equal(t, 10, game.columns)
		assert.Equal(t, 22, game.rows) // extra 2 rows for spawning tetrimino
		for r := range game.rows {
			for c := range game.columns {
				assert.Equal(t, BLANK_ID, game.playfield[r][c])
			}
		}
		// score should be 0
		assert.Equal(t, 0, game.score)
	})
	t.Run("game playfield coordinates are transformed correctly", func(t *testing.T) {
		// prepare
		game := CreateGame()

		// act
		game.playfield[0][3] = SHAPE_I

		// assert
		assert.Equal(t, SHAPE_I, game.At(22, 4))
	})
}

func TestTetriminoSpawn(t *testing.T) {
	t.Run("T tetrimino spawn at horizontally at 22 row", func(t *testing.T) {
		// prepare
		game := CreateGame()

		// act
		game.SpawnTetrimino(SHAPE_I)

		// assert
		assert.Equal(t, SHAPE_I, game.At(21, 4))
		assert.Equal(t, SHAPE_I, game.At(21, 5))
		assert.Equal(t, SHAPE_I, game.At(21, 6))
		assert.Equal(t, SHAPE_I, game.At(21, 7))
	})
}
