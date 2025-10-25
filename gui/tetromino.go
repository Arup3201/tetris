package gui

import (
	"github.com/Arup3201/tetris/api"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	MaxTransitionCount = 40
)

type tetrominoGui struct {
	gameApi         *api.Game
	transitionCount int
}

func createTetrominoGui(g *api.Game) *tetrominoGui {
	return &tetrominoGui{
		gameApi:         g,
		transitionCount: MaxTransitionCount,
	}
}

func (t *tetrominoGui) Update() error {
	if !t.gameApi.HasSpawned() {
		t.gameApi.SpawnTetrimino(api.SHAPE_I)
	} else {
		t.gameApi.DropByOne()
	}
	return nil
}

func (t *tetrominoGui) Draw(screen *ebiten.Image) {

}
