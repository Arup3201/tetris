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

func CreateGame(gamePlaygroundGridRows, gamePlaygroundGridColumns int,
	gameWindowWidth, gameWindowHeight int,
	gameWallOffsetX, gameWallOffsetY int) *game {
	g := api.CreateGame(gamePlaygroundGridRows, gamePlaygroundGridColumns)
	return &game{
		width:   gameWindowWidth,
		height:  gameWindowHeight,
		gameApi: g,
		wallGui: &guiWall{
			offsetX: gameWallOffsetX,
			offsetY: gameWallOffsetY,
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
