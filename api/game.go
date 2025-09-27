package api

type Game struct {
	wall       *Wall
	playground *Playground
}

func CreateGame(r, c int) *Game {
	return &Game{
		wall: &Wall{
			height: r + 2,
			width:  c + 2,
		},
		playground: CreatePlayground(r, c),
	}
}

func (g *Game) GetWall() []coord {
	return g.wall.GetBorder()
}
