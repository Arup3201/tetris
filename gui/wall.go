package gui

import (
	"github.com/Arup3201/tetris/api"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	wallBlockSprite = mustLoadImage("assets/tile.png")
	wallBlockWidth  = wallBlockSprite.Bounds().Dx()
	wallBlockHeight = wallBlockSprite.Bounds().Dy()
)

type wallGui struct {
	gameApi *api.Game
}

func (w *wallGui) Draw(screen *ebiten.Image) {
	coordinates := w.gameApi.GetWallCoordinates()
	opt := &ebiten.DrawImageOptions{}

	for _, rc := range coordinates {
		opt.GeoM.Translate(float64(wallBlockWidth*rc[1]),
			float64(wallBlockHeight*rc[0]))
		screen.DrawImage(wallBlockSprite, opt)

		opt.GeoM.Translate(-float64(wallBlockWidth*rc[1]),
			-float64(wallBlockHeight*rc[0]))
	}
}
