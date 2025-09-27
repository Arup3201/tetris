package api

type field struct {
	grid [][]int
}

func CreateField(r, c int) *field {
	field2D := make([][]int, r)
	for i := range r {
		field2D[i] = make([]int, c)
	}

	return &field{
		grid: field2D,
	}
}

func (f *field) SetGround(tetrominoCoordinates []tetrominoCoord) {
	for _, tCoord := range tetrominoCoordinates {
		f.grid[tCoord.s1.Y][tCoord.s1.X] = 1
		f.grid[tCoord.s2.Y][tCoord.s2.X] = 1
		f.grid[tCoord.s3.Y][tCoord.s3.X] = 1
		f.grid[tCoord.s4.Y][tCoord.s4.X] = 1
	}
}
