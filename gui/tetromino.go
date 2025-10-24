package gui

import (
	"image"

	"github.com/Arup3201/tetris/api"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	MaxTransitionCount = 40
)

var (
	TetriminoSheet = mustLoadImage("assets/tetrominoes.png")
	BlockWidth     = 37
	BlockHeight    = 37
)

type tetrominoGui struct {
	gameApi         *api.Game
	transitionCount int
	pieces          map[string]*ebiten.Image
}

func createTetrominoGui(g *api.Game) *tetrominoGui {
	pieces := map[string]*ebiten.Image{
		api.SHAPE_I: TetriminoSheet.SubImage(image.Rect(0, BlockHeight, BlockWidth, 2*BlockHeight)).(*ebiten.Image),
		api.SHAPE_J: TetriminoSheet.SubImage(image.Rect(5*BlockWidth, BlockHeight, 6*BlockWidth, 2*BlockHeight)).(*ebiten.Image),
		api.SHAPE_L: TetriminoSheet.SubImage(image.Rect(9*BlockWidth, BlockHeight, 10*BlockWidth, 2*BlockHeight)).(*ebiten.Image),
		api.SHAPE_O: TetriminoSheet.SubImage(image.Rect(13*BlockWidth, BlockHeight, 14*BlockWidth, 2*BlockHeight)).(*ebiten.Image),
		api.SHAPE_S: TetriminoSheet.SubImage(image.Rect(16*BlockWidth, BlockHeight, 17*BlockWidth, 2*BlockHeight)).(*ebiten.Image),
		api.SHAPE_Z: TetriminoSheet.SubImage(image.Rect(21*BlockWidth, BlockHeight, 22*BlockWidth, 2*BlockHeight)).(*ebiten.Image),
		api.SHAPE_T: TetriminoSheet.SubImage(image.Rect(24*BlockWidth, BlockHeight, 25*BlockWidth, 2*BlockHeight)).(*ebiten.Image),
	}

	return &tetrominoGui{
		gameApi:         g,
		transitionCount: MaxTransitionCount,
		pieces:          pieces,
	}
}

func (t *tetrominoGui) Update() error {
	return nil
}

func (t *tetrominoGui) Draw(screen *ebiten.Image) {
	screen.DrawImage(t.pieces[api.SHAPE_T], nil)
}
