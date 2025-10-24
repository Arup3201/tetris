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
				assert.Equal(t, BLANK_ID, game.playfield.Matrix[r][c])
			}
		}
		// score should be 0
		assert.Equal(t, 0, game.score)
	})
	t.Run("game playfield set success", func(t *testing.T) {
		// prepare
		game := CreateGame()

		// act
		game.SetPlayfield(22, 4, SHAPE_I)

		// assert
		assert.Equal(t, SHAPE_I, game.playfield.Matrix[0][3])
	})
	t.Run("game playfield get success", func(t *testing.T) {
		// prepare
		game := CreateGame()
		game.SetPlayfield(22, 4, SHAPE_I)

		// act
		got := game.GetPlayfield(22, 4)

		// assert
		assert.Equal(t, SHAPE_I, got)
	})
}

func TestTetriminoSpawn(t *testing.T) {
	t.Run("T tetrimino spawn at horizontally at 21 row", func(t *testing.T) {
		// prepare
		game := CreateGame()

		// act
		game.SpawnTetrimino(SHAPE_I)

		// assert
		assert.Equal(t, BLANK_ID, game.GetPlayfield(22, 4))
		assert.Equal(t, BLANK_ID, game.GetPlayfield(22, 5))
		assert.Equal(t, BLANK_ID, game.GetPlayfield(22, 6))
		assert.Equal(t, BLANK_ID, game.GetPlayfield(22, 7))
		assert.Equal(t, SHAPE_I, game.GetPlayfield(21, 4))
		assert.Equal(t, SHAPE_I, game.GetPlayfield(21, 5))
		assert.Equal(t, SHAPE_I, game.GetPlayfield(21, 6))
		assert.Equal(t, SHAPE_I, game.GetPlayfield(21, 7))
		assert.Equal(t, BLANK_ID, game.GetPlayfield(20, 4))
		assert.Equal(t, BLANK_ID, game.GetPlayfield(20, 5))
		assert.Equal(t, BLANK_ID, game.GetPlayfield(20, 6))
		assert.Equal(t, BLANK_ID, game.GetPlayfield(20, 7))
		assert.Equal(t, BLANK_ID, game.GetPlayfield(19, 4))
		assert.Equal(t, BLANK_ID, game.GetPlayfield(19, 5))
		assert.Equal(t, BLANK_ID, game.GetPlayfield(19, 6))
		assert.Equal(t, BLANK_ID, game.GetPlayfield(19, 7))
	})
	t.Run("T tetrimino spawns but previous tetrimino at 20 and 19 row", func(t *testing.T) {
		// prepare
		game := CreateGame()
		game.SetPlayfield(20, 4, SHAPE_I)
		game.SetPlayfield(19, 4, SHAPE_I)

		// act
		game.SpawnTetrimino(SHAPE_I)

		// assert
		assert.Equal(t, BLANK_ID, game.GetPlayfield(22, 4))
		assert.Equal(t, BLANK_ID, game.GetPlayfield(22, 5))
		assert.Equal(t, BLANK_ID, game.GetPlayfield(22, 6))
		assert.Equal(t, BLANK_ID, game.GetPlayfield(22, 7))
		assert.Equal(t, SHAPE_I, game.GetPlayfield(21, 4))
		assert.Equal(t, SHAPE_I, game.GetPlayfield(21, 5))
		assert.Equal(t, SHAPE_I, game.GetPlayfield(21, 6))
		assert.Equal(t, SHAPE_I, game.GetPlayfield(21, 7))
		assert.Equal(t, SHAPE_I, game.GetPlayfield(20, 4))
		assert.Equal(t, BLANK_ID, game.GetPlayfield(20, 5))
		assert.Equal(t, BLANK_ID, game.GetPlayfield(20, 6))
		assert.Equal(t, BLANK_ID, game.GetPlayfield(20, 7))
		assert.Equal(t, SHAPE_I, game.GetPlayfield(19, 4))
		assert.Equal(t, BLANK_ID, game.GetPlayfield(19, 5))
		assert.Equal(t, BLANK_ID, game.GetPlayfield(19, 6))
		assert.Equal(t, BLANK_ID, game.GetPlayfield(19, 7))
	})
	t.Run("T tetrimino spawn but already occupied 21 row", func(t *testing.T) {
		// prepare
		game := CreateGame()
		game.SetPlayfield(21, 4, SHAPE_I)

		// act
		ok := game.SpawnTetrimino(SHAPE_I)

		// assert
		assert.Equal(t, false, ok)
		assert.Equal(t, nil, game.spawned)
	})
}
