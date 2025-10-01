package api

type Playground struct {
	grid [][]*Square
}

func CreatePlayground(r, c int) *Playground {
	grid := make([][]*Square, r)
	for i := range r {
		grid[i] = make([]*Square, c)
	}

	return &Playground{
		grid: grid,
	}
}

func (f *Playground) SetGround(squares []*Square) {
	for _, square := range squares {
		f.grid[square.getY()][square.getX()] = square
	}
}
