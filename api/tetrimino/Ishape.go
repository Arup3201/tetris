package tetrimino

import "github.com/Arup3201/tetris/api/matrix"

const (
	DEGREE_0   = "0deg"
	DEGREE_90  = "90deg"
	DEGREE_180 = "180deg"
	DEGREE_270 = "270deg"
)

type ITetrimino struct {
	boundingBox [4][4]uint8
	name        string
	top, left   int
	rotation    string
}

func CreateITetrimino(name string) *ITetrimino {
	return &ITetrimino{
		boundingBox: [4][4]uint8{
			{0, 0, 0, 0},
			{1, 1, 1, 1},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		},
		name:     name,
		top:      22,
		left:     4,
		rotation: DEGREE_0,
	}
}

func (t *ITetrimino) removeFromPlayfield(matrix *matrix.ReverseMatrix, emptyCellValue string) {
	for r := range len(t.boundingBox) {
		for c := range len(t.boundingBox[0]) {
			if t.boundingBox[r][c] == 1 {
				matrix.Set(t.top-r, t.left+c, emptyCellValue)
			}
		}
	}
}

func (t *ITetrimino) addToPlayfield(matrix *matrix.ReverseMatrix) {
	for r := range len(t.boundingBox) {
		for c := range len(t.boundingBox[0]) {
			if t.boundingBox[r][c] == 1 {
				matrix.Set(t.top-r, t.left+c, t.name)
			}
		}
	}
}

func (t *ITetrimino) Spawn(matrix *matrix.ReverseMatrix, emptyCellValue string) bool {
	for r := range len(t.boundingBox) {
		for c := range len(t.boundingBox[0]) {
			if t.boundingBox[r][c] == 1 {
				cell, _ := matrix.Get(t.top-r, t.left+c)
				if cell != emptyCellValue {
					return false
				}
			}
		}
	}

	t.addToPlayfield(matrix)
	return true
}

func (t *ITetrimino) DropByOne(matrix *matrix.ReverseMatrix, emptyCellValue string) bool {
	// hit the ground
	switch {
	case t.rotation == DEGREE_0 && t.top == 2:
		return false
	case t.rotation == DEGREE_90 && t.top == 4:
		return false
	case t.rotation == DEGREE_180 && t.top == 3:
		return false
	case t.rotation == DEGREE_270 && t.top == 4:
		return false
	}

	// hit another tetrimino piece
	switch t.rotation {
	case DEGREE_0, DEGREE_180:
		for r := range len(t.boundingBox) {
			for c := range len(t.boundingBox[0]) {
				if t.boundingBox[r][c] == 1 {
					cell, _ := matrix.Get(t.top-r-1, t.left+c)
					if cell != emptyCellValue {
						return false
					}
				}
			}
		}
	case DEGREE_90:
		cell, _ := matrix.Get(t.top-4, t.left+2)
		if cell != emptyCellValue {
			return false
		}
	case DEGREE_270:
		cell, _ := matrix.Get(t.top-4, t.left+1)
		if cell != emptyCellValue {
			return false
		}
	}

	t.removeFromPlayfield(matrix, emptyCellValue)
	t.top--
	t.addToPlayfield(matrix)

	return true
}

func (t *ITetrimino) Rotate(matrix *matrix.ReverseMatrix, emptyCellValue string) bool {
	t.removeFromPlayfield(matrix, emptyCellValue)

	switch t.rotation {
	case DEGREE_0:
		t.boundingBox = [4][4]uint8{
			{0, 0, 1, 0},
			{0, 0, 1, 0},
			{0, 0, 1, 0},
			{0, 0, 1, 0},
		}
		t.rotation = DEGREE_90
	case DEGREE_90:
		t.boundingBox = [4][4]uint8{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{1, 1, 1, 1},
			{0, 0, 0, 0},
		}
		t.rotation = DEGREE_180
	case DEGREE_180:
		t.boundingBox = [4][4]uint8{
			{0, 1, 0, 0},
			{0, 1, 0, 0},
			{0, 1, 0, 0},
			{0, 1, 0, 0},
		}
		t.rotation = DEGREE_270
	case DEGREE_270:
		t.boundingBox = [4][4]uint8{
			{0, 0, 0, 0},
			{1, 1, 1, 1},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		}
		t.rotation = DEGREE_0
	}

	t.addToPlayfield(matrix)

	return true
}
