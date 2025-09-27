package gui

import (
	"image"
	_ "image/png"

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

type wall struct {
	offsetX, offsetY int
	width, height    int
}

func (w *wall) draw(screen *ebiten.Image) {
	newTileX, newTileY := 0, 0
	op := &ebiten.DrawImageOptions{}

	// top
	for range w.width {
		op.GeoM.Translate(float64(newTileX), float64(newTileY))
		screen.DrawImage(tileSprite, op)
		op.GeoM.Translate(-float64(newTileX), -float64(newTileY)) // revert relative position
		newTileX += w.offsetX + tileWidth
	}

	// left - avoid rewrite top-left tile
	newTileX = 0
	newTileY += tileHeight
	for range w.height - 1 {
		op.GeoM.Translate(float64(newTileX), float64(newTileY))
		screen.DrawImage(tileSprite, op)
		op.GeoM.Translate(-float64(newTileX), -float64(newTileY)) // revert relative position
		newTileY += w.offsetY + tileHeight
	}

	// right - avoid rewrite top-right tile
	newTileX = (w.width - 1) * tileWidth
	newTileY = tileHeight
	for range w.height - 1 {
		op.GeoM.Translate(float64(newTileX), float64(newTileY))
		screen.DrawImage(tileSprite, op)
		op.GeoM.Translate(-float64(newTileX), -float64(newTileY)) // revert relative position
		newTileY += w.offsetY + tileHeight
	}

	// bottom - avoid bottom-left and bottom-right tile
	newTileX = tileWidth
	newTileY = (w.height - 1) * tileHeight
	for range w.width - 2 {
		op.GeoM.Translate(float64(newTileX), float64(newTileY))
		screen.DrawImage(tileSprite, op)
		op.GeoM.Translate(-float64(newTileX), -float64(newTileY)) // revert relative position
		newTileX += w.offsetX + tileWidth
	}
}
