package gui

import (
	"github.com/Arup3201/tetris/api"
	"github.com/hajimehoshi/ebiten/v2"
)

var squareSprite = mustLoadImage("assets/square.png")
var (
	squareWidth  = tileSprite.Bounds().Dx()
	squareHeight = tileSprite.Bounds().Dy()
)

type guiTetromino struct {
	gameApi *api.Game
}

func (t *guiTetromino) Update() error {
	if t.gameApi.TetrominoDidNotEnterField() {
		t.gameApi.InitTetromino()
	} else if t.gameApi.TetrominoHasHit() {
		squares := t.gameApi.GetTetrominoSquares()
		t.gameApi.SetPlayground([]*api.Square{squares[0], squares[1], squares[2], squares[3]})
	} else {
		t.gameApi.TetrminoGoDown()
	}
	return nil
}

func (t *guiTetromino) Draw(screen *ebiten.Image) {
	squareX, squareY := 0, 0
	positions := t.gameApi.GetTetrominoPosition()
	op := &ebiten.DrawImageOptions{}

	for _, pos := range positions {
		if pos.X != -1 && pos.Y != -1 {
			squareX, squareY = pos.X*squareWidth, pos.Y*squareHeight
			op.GeoM.Translate(float64(squareX), float64(squareY))
			screen.DrawImage(squareSprite, op)
			op.GeoM.Translate(-float64(squareX), -float64(squareY)) // revert relative position
		}
	}
}
