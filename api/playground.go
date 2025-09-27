package api

type Playground struct {
	grid [][]int
}

func CreatePlayground(r, c int) *Playground {
	grid := make([][]int, r)
	for i := range r {
		grid[i] = make([]int, c)
	}

	return &Playground{
		grid: grid,
	}
}

func (f *Playground) SetGround(tetrominoCoordinates []tetrominoCoord) {
	for _, tCoord := range tetrominoCoordinates {
		f.grid[tCoord.s1.Y][tCoord.s1.X] = 1
		f.grid[tCoord.s2.Y][tCoord.s2.X] = 1
		f.grid[tCoord.s3.Y][tCoord.s3.X] = 1
		f.grid[tCoord.s4.Y][tCoord.s4.X] = 1
	}
}
