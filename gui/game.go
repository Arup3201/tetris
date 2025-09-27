package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type game struct {
	rows, columns int
	gameWall      *wall
}

func CreateGame(r, c int) *game {
	return &game{
		rows:    r,
		columns: c,
		gameWall: &wall{
			width:  c + 2,
			height: r + 2,
		},
	}
}

func (g *game) Update() error {
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	g.gameWall.draw(screen)
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func (g *game) Run(width, height int, title string) error {
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle(title)
	return ebiten.RunGame(g)
}
