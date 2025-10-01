package gui

import (
	"github.com/Arup3201/tetris/api"
	"github.com/hajimehoshi/ebiten/v2"
)

var tileSprite = mustLoadImage("assets/tile.png")
var (
	tileWidth  = tileSprite.Bounds().Dx()
	tileHeight = tileSprite.Bounds().Dy()
)

type guiWall struct {
	offsetX, offsetY int
	gameApi          *api.Game
}

func (w *guiWall) Draw(screen *ebiten.Image) {
	newTileX, newTileY := 0, 0
	op := &ebiten.DrawImageOptions{}

	border := w.gameApi.GetWall()
	for _, coord := range border {
		newTileX, newTileY = w.offsetX+coord.X*tileWidth, w.offsetY+coord.Y*tileHeight
		op.GeoM.Translate(float64(newTileX), float64(newTileY))
		screen.DrawImage(tileSprite, op)
		op.GeoM.Translate(-float64(newTileX), -float64(newTileY)) // revert relative position
	}
}
