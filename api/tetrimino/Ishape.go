package tetrimino

type ITetrimino struct {
	BoundingBox [4][4]uint8
	name        string
}

func CreateITetrimino(name string) *ITetrimino {
	return &ITetrimino{
		BoundingBox: [4][4]uint8{
			{0, 0, 0, 0},
			{1, 1, 1, 1},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		},
		name: name,
	}
}

func (t *ITetrimino) Spawn(matrix [][]string) {
	for r := range 4 {
		for c := range 4 {
			if t.BoundingBox[r][c] == 1 {
				matrix[r][c] = t.name
			}
		}
	}
}
