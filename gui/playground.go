package gui

import (
	"image/color"

	"github.com/Arup3201/tetris/api"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type playgroundGui struct {
	gameApi          *api.Game
	playgroundSprite *ebiten.Image
}

func createPlaygroundGUI(g *api.Game) *playgroundGui {
	columns, rows := api.MATRIX_COLUMNS, api.MATRIX_ROWS
	sprite := ebiten.NewImage(columns*BlockWidth, rows*BlockWidth)
	vector.DrawFilledRect(sprite,
		0, 0, float32(columns*BlockWidth), float32(rows*BlockHeight),
		color.RGBA{212, 212, 212, 255}, true)
	for r := range rows {
		for c := range columns {
			vector.DrawFilledRect(sprite,
				float32(c*BlockWidth), float32(r*BlockHeight), float32(BlockWidth-2), float32(BlockHeight-2),
				color.RGBA{245, 245, 245, 255}, true)
		}
	}

	return &playgroundGui{
		gameApi:          g,
		playgroundSprite: sprite,
	}
}

func (p *playgroundGui) Draw(screen *ebiten.Image) {
	screen.DrawImage(p.playgroundSprite, nil)
}
