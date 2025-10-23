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
		assert.Equal(t, 20, game.rows)
		for r := range game.rows {
			for c := range game.columns {
				assert.Equal(t, BLANK_ID, game.playfield[r][c])
			}
		}
		// score should be 0
		assert.Equal(t, 0, game.score)
	})
}

func TestTetriminoSpawn(t *testing.T) {
	t.Run("T tetrimino spawn at horizontally at position 10, 11, 12, 13", func(t *testing.T) {
		// prepare
		game := CreateGame()

		// act
		game.SpawnTetrimino(SHAPE_I)

		// assert
		assert.Equal(t, SHAPE_I, game.playfield[1][0])
		assert.Equal(t, SHAPE_I, game.playfield[1][1])
		assert.Equal(t, SHAPE_I, game.playfield[1][2])
		assert.Equal(t, SHAPE_I, game.playfield[1][3])
	})
}
