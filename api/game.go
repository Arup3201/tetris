package api

const (
	WALL_ID      = "wall"
	BLANK_ID     = "blank"
	SQUARE_ID    = "square"
	COLOR_YELLOW = "yellow"
)

type tetromino interface {
	getPosition() [4][2]int
	drop()
	fallByOne([][]string)
	moveRightByOne([][]string)
	hasHit([][]string) bool
}

type Game struct {
	rows, columns     int
	grid              [][]string
	score             int
	droppingTetromino tetromino

	HasTetrominoDropped bool
}

func CreateGame(gamePlaygroundRows, gamePlaygroundColumns int) *Game {
	totalRows := gamePlaygroundRows + 2
	totalColumns := gamePlaygroundColumns + 2

	grid := make([][]string, totalRows)
	for r := range totalRows {
		grid[r] = make([]string, totalColumns)
	}

	for c := range totalColumns {
		grid[0][c] = WALL_ID
		grid[totalRows-1][c] = WALL_ID
	}
	for r := range totalRows {
		grid[r][0] = WALL_ID
		grid[r][totalColumns-1] = WALL_ID
	}

	for r := range totalRows - 2 {
		for c := range totalColumns - 2 {
			grid[r+1][c+1] = BLANK_ID
		}
	}

	return &Game{
		rows:    totalRows,
		columns: totalColumns,
		grid:    grid,
		score:   0,
	}
}

func (g *Game) GetWallCoordinates() [][2]int {
	coordinates := [][2]int{}
	for r := range g.rows {
		for c := range g.columns {
			if g.grid[r][c] == WALL_ID {
				coordinates = append(coordinates, [2]int{r, c})
			}
		}
	}

	return coordinates
}

func (g *Game) GetPlayground() map[[2]int]string {
	playground := make(map[[2]int]string, 0)
	for r := range g.rows {
		for c := range g.columns {
			if g.grid[r][c] != WALL_ID {
				playground[[2]int{r, c}] = g.grid[r][c]
			}
		}
	}

	return playground
}

func (g *Game) GetTetrominoPosition() [4][2]int {
	coordinates := g.droppingTetromino.getPosition()
	// consider the walls
	for i := range coordinates {
		if coordinates[i][0] != -1 { // coordinates[i][1] != -1
			coordinates[i][0] += 1
			coordinates[i][1] += 1
		}
	}
	return coordinates
}

func (g *Game) DropTetromino(shape, color string) {
	// tetromino does not consider walls internally
	g.droppingTetromino = newTetromino(shape, color, [2]int{g.rows - 2, g.columns - 2})
	g.droppingTetromino.drop()
	g.HasTetrominoDropped = true
}

func (g *Game) TetrominoFallsByOne() {
	if !g.HasTetrominoDropped {
		return
	}

	playground := make([][]string, g.rows-2)
	for r := range playground {
		playground[r] = make([]string, g.columns-2)
	}
	for r := 1; r <= g.rows-2; r++ {
		for c := 1; c <= g.columns-2; c++ {
			playground[r-1][c-1] = g.grid[r][c]
		}
	}

	g.droppingTetromino.fallByOne(playground)

	if g.droppingTetromino.hasHit(playground) {
		coordinates := g.GetTetrominoPosition()
		g.grid[coordinates[0][0]][coordinates[0][1]] = SQUARE_ID
		g.grid[coordinates[1][0]][coordinates[1][1]] = SQUARE_ID
		g.grid[coordinates[2][0]][coordinates[2][1]] = SQUARE_ID
		g.grid[coordinates[3][0]][coordinates[3][1]] = SQUARE_ID

		g.HasTetrominoDropped = false
	}
}

func (g *Game) MoveTetrominoRight() {
	playground := make([][]string, g.rows-2)
	for r := range playground {
		playground[r] = make([]string, g.columns-2)
	}
	for r := 1; r <= g.rows-2; r++ {
		for c := 1; c <= g.columns-2; c++ {
			playground[r-1][c-1] = g.grid[r][c]
		}
	}

	g.droppingTetromino.moveRightByOne(playground)
}
