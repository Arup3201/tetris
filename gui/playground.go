package gui

import (
	"image"
	"image/color"

	"github.com/Arup3201/tetris/api"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	TetriminoSheet = mustLoadImage("assets/tetrominoes.png")
	BlockWidth     = 37
	BlockHeight    = 37
)

type playgroundGui struct {
	gameApi          *api.Game
	playgroundSprite *ebiten.Image
	blocks           map[string]*ebiten.Image
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

	blocks := map[string]*ebiten.Image{
		api.SHAPE_I: TetriminoSheet.SubImage(image.Rect(0, BlockHeight, BlockWidth, 2*BlockHeight)).(*ebiten.Image),
		api.SHAPE_J: TetriminoSheet.SubImage(image.Rect(5*BlockWidth, BlockHeight, 6*BlockWidth, 2*BlockHeight)).(*ebiten.Image),
		api.SHAPE_L: TetriminoSheet.SubImage(image.Rect(9*BlockWidth, BlockHeight, 10*BlockWidth, 2*BlockHeight)).(*ebiten.Image),
		api.SHAPE_O: TetriminoSheet.SubImage(image.Rect(13*BlockWidth, BlockHeight, 14*BlockWidth, 2*BlockHeight)).(*ebiten.Image),
		api.SHAPE_S: TetriminoSheet.SubImage(image.Rect(16*BlockWidth, BlockHeight, 17*BlockWidth, 2*BlockHeight)).(*ebiten.Image),
		api.SHAPE_Z: TetriminoSheet.SubImage(image.Rect(21*BlockWidth, BlockHeight, 22*BlockWidth, 2*BlockHeight)).(*ebiten.Image),
		api.SHAPE_T: TetriminoSheet.SubImage(image.Rect(24*BlockWidth, BlockHeight, 25*BlockWidth, 2*BlockHeight)).(*ebiten.Image),
	}

	return &playgroundGui{
		gameApi:          g,
		playgroundSprite: sprite,
		blocks:           blocks,
	}
}

func (p *playgroundGui) Draw(screen *ebiten.Image) {
	screen.DrawImage(p.playgroundSprite, nil)
	opt := &ebiten.DrawImageOptions{}

	for r := 1; r <= api.MATRIX_ROWS; r++ {
		for c := 1; c <= api.MATRIX_COLUMNS; c++ {
			if cell := p.gameApi.GetPlayfield(r, c); cell != api.BLANK_ID {
				row, column := api.MATRIX_ROWS-r, c-1
				opt.GeoM.Translate(float64(column*BlockWidth), float64(row*BlockHeight))
				screen.DrawImage(p.blocks[cell], opt)

				opt.GeoM.Translate(-float64(column*BlockWidth), -float64(row*BlockHeight))
			}
		}
	}
}
