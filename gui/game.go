package gui

import (
	"github.com/Arup3201/tetris/api"
	"github.com/hajimehoshi/ebiten/v2"
)

type game struct {
	gameApi *api.Game
	wallGui *guiWall
}

func CreateGame(r, c int) *game {
	g := api.CreateGame(r, c)
	return &game{
		gameApi: g,
		wallGui: &guiWall{
			gameApi: g,
		},
	}
}

func (g *game) Update() error {
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	g.wallGui.draw(screen)
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func (g *game) Run(width, height int, title string) error {
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle(title)
	return ebiten.RunGame(g)
}
