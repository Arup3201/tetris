package gui

import (
	"image/color"

	"github.com/Arup3201/tetris/api"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type guiPlayground struct {
	gameApi          *api.Game
	playgroundSprite *ebiten.Image
}

func CreatePlaygroundGUI(g *api.Game, rows, columns int) *guiPlayground {
	sprite := ebiten.NewImage(columns*squareWidth, rows*squareHeight)
	for r := range rows {
		if r == 0 {
			continue
		}
		vector.StrokeLine(sprite, 0, float32(r*squareHeight), float32(columns*squareWidth), float32(r*squareHeight), 1, color.RGBA{
			R: 127,
			G: 127,
			B: 127,
			A: 1,
		}, false)
	}
	for c := range columns {
		if c == 0 {
			continue
		}
		vector.StrokeLine(sprite, float32(c*squareWidth), 0, float32(c*squareWidth), float32(rows*squareHeight), 1, color.RGBA{
			R: 127,
			G: 127,
			B: 127,
			A: 1,
		}, false)
	}

	return &guiPlayground{
		gameApi:          g,
		playgroundSprite: sprite,
	}
}

func (p *guiPlayground) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(tileWidth), float64(tileHeight))
	screen.DrawImage(p.playgroundSprite, op)
	op.GeoM.Translate(-float64(tileWidth), -float64(tileHeight))
}
