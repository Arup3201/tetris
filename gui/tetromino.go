package gui

import (
	"bytes"
	"image/color"
	"log"

	"github.com/Arup3201/tetris/api"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/text/language"
)

var (
	mplusFaceSource *text.GoTextFaceSource
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	mplusFaceSource = s
}

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
		op := &text.DrawOptions{}
		op.ColorScale.ScaleWithColor(color.RGBA{0, 0, 0, 255})
		f := &text.GoTextFace{
			Source:    mplusFaceSource,
			Direction: text.DirectionLeftToRight,
			Size:      24,
			Language:  language.English,
		}
		text.Draw(screen, gameOverText, f, op)
	}
}
