package gui

import (
	"image/color"

	"github.com/Arup3201/tetris/api"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	squareSprite = mustLoadImage("assets/square.png")
	squareHeight = squareSprite.Bounds().Dy()
	squareWidth  = squareSprite.Bounds().Dx()
)

type playgroundGui struct {
	gameApi          *api.Game
	playgroundSprite *ebiten.Image
}

func createPlaygroundGUI(g *api.Game, rows, columns int) *playgroundGui {
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

	return &playgroundGui{
		gameApi:          g,
		playgroundSprite: sprite,
	}
}

func (p *playgroundGui) Draw(screen *ebiten.Image) {
	playground := p.gameApi.GetPlayground()
	opt := &ebiten.DrawImageOptions{}

	// move after the wall
	opt.GeoM.Translate(float64(wallBlockWidth), float64(wallBlockHeight))
	screen.DrawImage(p.playgroundSprite, opt)
	opt.GeoM.Translate(-float64(wallBlockWidth), -float64(wallBlockHeight))

	for _, id := range playground {
		if id != api.BLANK_ID {
			// draw square
		}
	}
}
