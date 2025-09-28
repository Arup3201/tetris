package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGamePlayground(t *testing.T) {
	rows, columns := 20, 10

	playground := CreatePlayground(rows, columns)

	assert.Equal(t, 0, playground.grid[0][0])
	assert.Equal(t, 0, playground.grid[rows-1][columns-1])
}
