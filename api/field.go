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
		f.grid[tCoord.S1.Y][tCoord.S1.X] = 1
		f.grid[tCoord.S2.Y][tCoord.S2.X] = 1
		f.grid[tCoord.S3.Y][tCoord.S3.X] = 1
		f.grid[tCoord.S4.Y][tCoord.S4.X] = 1
	}
}
