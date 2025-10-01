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
	squares := t.gameApi.GetTetrominoSquares()
	op := &ebiten.DrawImageOptions{}

	for _, sq := range squares {
		if sq.GetX() != -1 && sq.GetY() != -1 {
			squareX, squareY = (sq.GetX()+1)*squareWidth, (sq.GetY()+1)*squareHeight
			op.GeoM.Translate(float64(squareX), float64(squareY))
			op.ColorScale.SetR(float32(sq.Color.R))
			op.ColorScale.SetG(float32(sq.Color.G))
			op.ColorScale.SetB(float32(sq.Color.B))
			op.ColorScale.SetA(float32(sq.Color.A))
			screen.DrawImage(squareSprite, op)
			op.GeoM.Translate(-float64(squareX), -float64(squareY)) // revert relative position
		}
	}
}
