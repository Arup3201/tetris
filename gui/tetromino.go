package gui

import (
	"github.com/Arup3201/tetris/api"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/text/language"
)

type tetrominoGui struct {
	gameApi    *api.Game
	isGameOver bool
}

func createTetrominoGui(g *api.Game) *tetrominoGui {
	return &tetrominoGui{
		gameApi: g,
	}
}

func (t *tetrominoGui) Update() error {
	if !t.isGameOver {
		if !t.gameApi.HasSpawned() {
			success := t.gameApi.SpawnTetrimino(api.SHAPE_I)
			if !success {
				t.isGameOver = true
			}
		} else {
			t.gameApi.DropByOne()
		}
	}

	return nil
}

func (t *tetrominoGui) Draw(screen *ebiten.Image) {
	if t.isGameOver {
		const gameOverText = "Game Over!"
		f := &text.GoTextFace{
			Direction: text.DirectionLeftToRight,
			Size:      24,
			Language:  language.English,
		}
		text.Draw(screen, gameOverText, f, nil)
	}
}
