package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWallBorder(t *testing.T) {
	width, height := 6, 6
	wall := &Wall{
		width:  width,
		height: height,
	}

	border := wall.GetBorder()

	expected := []coord{
		{
			X: 0,
			Y: 0,
		},
		{
			X: 1,
			Y: 0,
		},
		{
			X: 2,
			Y: 0,
		},
		{
			X: 3,
			Y: 0,
		},
		{
			X: 4,
			Y: 0,
		},
		{
			X: 5,
			Y: 0,
		},
		{
			X: 0,
			Y: 1,
		},
		{
			X: 5,
			Y: 1,
		},
		{
			X: 0,
			Y: 2,
		},
		{
			X: 5,
			Y: 2,
		},
		{
			X: 0,
			Y: 3,
		},
		{
			X: 5,
			Y: 3,
		},
		{
			X: 0,
			Y: 4,
		},
		{
			X: 5,
			Y: 4,
		},
		{
			X: 0,
			Y: 5,
		},
		{
			X: 1,
			Y: 5,
		},
		{
			X: 2,
			Y: 5,
		},
		{
			X: 3,
			Y: 5,
		},
		{
			X: 4,
			Y: 5,
		},
		{
			X: 5,
			Y: 5,
		},
	}
	assert.Equal(t, expected, border)
}
