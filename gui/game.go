package gui

import (
	"github.com/Arup3201/tetris/api"
	"github.com/hajimehoshi/ebiten/v2"
)

type game struct {
	width, height int
	gameApi       *api.Game
	wallGui       *guiWall
}

func CreateGame(r, c, winWidth, winHeight int) *game {
	g := api.CreateGame(r, c)
	return &game{
		width:   winWidth,
		height:  winHeight,
		gameApi: g,
		wallGui: &guiWall{
			offsetX: 40,
			offsetY: 40,
			gameApi: g,
		},
	}
}

func (g *game) Update() error {
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	g.wallGui.Draw(screen)
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.width, g.height
}

func (g *game) Run(title string) error {
	ebiten.SetWindowSize(g.width, g.height)
	ebiten.SetWindowTitle(title)
	return ebiten.RunGame(g)
}
