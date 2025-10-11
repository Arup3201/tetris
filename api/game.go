package api

const (
	WALL_ID  = "wall"
	BLANK_ID = "blank"
)

type Game struct {
	width, height int
	grid          [][]string
	score         int
}

func CreateGame(gamePlaygroundRows, gamePlaygroundColumns int) *Game {
	gameWidth, gameHeight := gamePlaygroundColumns+2, gamePlaygroundRows+2

	grid := make([][]string, gameHeight)
	for r := range gameHeight {
		grid[r] = make([]string, gameWidth)
	}

	for c := range gameWidth {
		grid[0][c] = WALL_ID
		grid[gameHeight-1][c] = WALL_ID
	}
	for r := range gameHeight {
		grid[r][0] = WALL_ID
		grid[r][gameWidth-1] = WALL_ID
	}

	for r := range gameHeight - 2 {
		for c := range gameWidth - 2 {
			grid[r+1][c+1] = BLANK_ID
		}
	}

	return &Game{
		width:  gameWidth,
		height: gameHeight,
		grid:   grid,
		score:  0,
	}
}

func (g *Game) GetWallCoordinates() [][2]int {
	coordinates := [][2]int{}
	for r := range g.height {
		for c := range g.width {
			if g.grid[r][c] == WALL_ID {
				coordinates = append(coordinates, [2]int{r, c})
			}
		}
	}

	return coordinates
}

func (g *Game) GetPlayground() map[[2]int]string {
	playground := make(map[[2]int]string, 0)
	for r := range g.height {
		for c := range g.width {
			if g.grid[r][c] != WALL_ID {
				playground[[2]int{r, c}] = g.grid[r][c]
			}
		}
	}

	return playground
}
