package tetris

func CreateField(r, c int) [][]int {
	field2D := make([][]int, r)
	for i := range r {
		field2D[i] = make([]int, c)
	}

	return field2D
}

type tetromino struct {
	x, y, bottom, right int
}

func CreateTetromino(bottom, right int) *tetromino {
	return &tetromino{
		x:      -1,
		y:      -1,
		bottom: bottom,
		right:  right,
	}
}

func (t *tetromino) EnterField() {
	t.x = t.right / 2
	t.y = 0
}

func (t *tetromino) GetPosition() (int, int) {
	return t.x, t.y
}

func (t *tetromino) GoDown(by int) {
	t.y += by
}
