package matrix

// 2D slice with bottom to top 1-based indexing
type ReverseMatrix struct {
	Matrix        [][]string
	columns, rows int
}

func CreateReverseMatrix(rowsCount, columnsCount int,
	initialValue string) *ReverseMatrix {
	matrix := make([][]string, rowsCount)
	for r := range rowsCount {
		matrix[r] = make([]string, columnsCount)
	}

	for r := range rowsCount {
		for c := range columnsCount {
			matrix[r][c] = initialValue
		}
	}
	return &ReverseMatrix{
		Matrix:  matrix,
		rows:    rowsCount,
		columns: columnsCount,
	}
}

func (m *ReverseMatrix) Set(row, column int, value string) bool {
	if row < 1 || row > m.rows {
		return false
	}
	if column < 1 || column > m.columns {
		return false
	}

	m.Matrix[m.rows-row][column-1] = value
	return true
}

func (m *ReverseMatrix) Get(row, column int) (string, bool) {
	if row < 1 || row > m.rows {
		return "", false
	}
	if column < 1 || column > m.columns {
		return "", false
	}

	return m.Matrix[m.rows-row][column-1], true
}
