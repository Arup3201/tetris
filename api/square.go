package api

type coord struct {
	X, Y int
}

type color struct {
	R, G, B, A uint8
}

var (
	COLOR_ORANGE = color{
		R: 255,
		G: 165,
		B: 0,
		A: 255,
	}
)

type Square struct {
	position coord
	Color    color
}

func createSquare(x, y int, color color) *Square {
	return &Square{
		position: coord{
			X: x,
			Y: y,
		},
		Color: color,
	}
}

func (sq *Square) changePosition(x, y int) {
	sq.position.X = x
	sq.position.Y = y
}

func (sq *Square) GetX() int {
	return sq.position.X
}

func (sq *Square) GetY() int {
	return sq.position.Y
}
