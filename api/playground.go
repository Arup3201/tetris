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

func (f *Playground) SetGround(squares []*square) {
	for _, square := range squares {
		f.grid[square.getY()][square.getX()] = 1
	}
}
