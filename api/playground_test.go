package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGamePlayground(t *testing.T) {
	// prepare
	rows, columns := 20, 10

	// act
	playground := CreatePlayground(rows, columns)

	// assert
	for r := range playground.grid {
		for c := range playground.grid[r] {
			assert.Equal(t, (*Square)(nil), playground.grid[r][c])
		}
	}
}

func TestGameSetPlayground(t *testing.T) {
	// prepare
	rows, columns := 20, 10
	playground := CreatePlayground(rows, columns)

	// act
	playground.SetGround([]*Square{
		createSquare(4, 19, COLOR_ORANGE),
		createSquare(5, 19, COLOR_ORANGE),
		createSquare(6, 19, COLOR_ORANGE),
	})

	// assert
	expected := []*Square{
		{
			position: coord{
				X: 4,
				Y: 19,
			},
			Color: COLOR_ORANGE,
		},
		{
			position: coord{
				X: 5,
				Y: 19,
			},
			Color: COLOR_ORANGE,
		},
		{
			position: coord{
				X: 6,
				Y: 19,
			},
			Color: COLOR_ORANGE,
		},
	}
	assert.Equal(t, expected[0], playground.grid[19][4])
	assert.Equal(t, expected[1], playground.grid[19][5])
	assert.Equal(t, expected[2], playground.grid[19][6])
}
