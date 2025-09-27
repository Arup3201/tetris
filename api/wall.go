package api

type Wall struct {
	width, height int
}

func (w *Wall) GetBorder() []coord {
	border := []coord{}

	for r := range w.height {
		for c := range w.width {
			if r != 0 && c != 0 && r != w.height-1 && c != w.width-1 {
				continue
			}
			border = append(border, coord{
				X: c,
				Y: r,
			})
		}
	}

	return border
}
