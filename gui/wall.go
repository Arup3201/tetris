package gui

import (
	"image"
	_ "image/png"

	"github.com/Arup3201/tetris/api"
	"github.com/hajimehoshi/ebiten/v2"
)

var tileSprite = mustLoadImage("assets/tile.png")
var (
	tileWidth  = tileSprite.Bounds().Dx()
	tileHeight = tileSprite.Bounds().Dy()
)

func mustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

type guiWall struct {
	gameApi *api.Game
}

func (w *guiWall) draw(screen *ebiten.Image) {
	newTileX, newTileY := 0, 0
	op := &ebiten.DrawImageOptions{}

	border := w.gameApi.GetWall()
	for _, coord := range border {
		newTileX, newTileY = coord.X*tileWidth, coord.Y*tileHeight
		op.GeoM.Translate(float64(newTileX), float64(newTileY))
		screen.DrawImage(tileSprite, op)
		op.GeoM.Translate(-float64(newTileX), -float64(newTileY)) // revert relative position
	}
}
