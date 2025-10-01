package api

type Game struct {
	gridWidth, gridHeight int
	wall                  *Wall
	playground            *Playground
	tetromino             *Tetromino
}

func CreateGame(r, c int) *Game {
	return &Game{
		gridHeight: r,
		gridWidth:  c,
		wall: &Wall{
			height: r + 2,
			width:  c + 2,
		},
		tetromino:  CreateTetromino(r, c),
		playground: CreatePlayground(r, c),
	}
}

func (g *Game) GetWall() []coord {
	return g.wall.GetBorder()
}

func (g *Game) InitTetromino() {
	g.tetromino.EnterField()
}

func (g *Game) TetrominoDidNotEnterField() bool {
	return g.tetromino.DidNotEnterField()
}

func (g *Game) GetTetrominoPosition() []coord {
	positions := g.tetromino.GetPosition()

	// account for the wall
	tetrominoPos := []coord{}
	for _, position := range positions {
		if position.X != -1 || position.Y != -1 {
			tetrominoPos = append(tetrominoPos, coord{
				X: position.X + 1,
				Y: position.Y + 1,
			})
		}
	}

	return tetrominoPos
}

func (g *Game) TetrminoGoDown() {
	g.tetromino.GoDown(g.playground.grid, 1)
}

func (g *Game) TetrominoHasHit() bool {
	return g.tetromino.HasHit(g.playground.grid)
}
