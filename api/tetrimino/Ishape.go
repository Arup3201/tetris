package tetrimino

type MatrixSetFunc func(r, c int)

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

func (t *ITetrimino) Spawn(fn MatrixSetFunc) {
	fn(21, 4)
	fn(21, 5)
	fn(21, 6)
	fn(21, 7)
}
