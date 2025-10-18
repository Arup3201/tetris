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
	if !t.gameApi.HasTetrominoDropped {
		t.gameApi.DropTetromino(api.SHAPE_T, api.COLOR_YELLOW)
	} else if t.transitionCount != 0 {
		t.transitionCount--
	} else {
		t.gameApi.TetrominoFallsByOne()
		t.transitionCount = MaxTransitionCount
	}
	return nil
}

func (t *tetrominoGui) Draw(screen *ebiten.Image) {
	coordinates := t.gameApi.GetTetrominoPosition()
	opt := &ebiten.DrawImageOptions{}

	for _, rc := range coordinates {
		if rc[0] != -1 { // rc[1] != -1
			opt.GeoM.Translate(float64(rc[1]*squareWidth),
				float64(rc[0]*squareHeight))
			screen.DrawImage(squareSprite, opt)

			opt.GeoM.Translate(-float64(rc[1]*squareWidth),
				-float64(rc[0]*squareHeight))
		}
	}
}
