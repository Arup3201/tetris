package api

type Game struct {
	width, height int
	grid          [][]string
	score         int
}

func CreateGame(r, c int) *Game {
	gameWidth, gameHeight := c+2, r+2

	grid := make([][]string, gameHeight)
	for r := range gameHeight {
		grid[r] = make([]string, gameWidth)
	}

	for c := range gameWidth {
		grid[0][c] = "wall"
		grid[gameHeight-1][c] = "wall"
	}
	for r := range gameHeight {
		grid[r][0] = "wall"
		grid[r][gameWidth-1] = "wall"
	}

	for r := range gameHeight - 2 {
		for c := range gameWidth - 2 {
			grid[r+1][c+1] = "blank"
		}
	}

	return &Game{
		width:  gameWidth,
		height: gameHeight,
		grid:   grid,
		score:  0,
	}
}
