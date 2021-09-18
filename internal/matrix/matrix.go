package matrix

import "strings"

type row []rune

type matrix struct {
	rows          []row
	size          int
	currentRow    int
	currentColumn int
	columns       int
}

func NewMatrix(n int, key int) *matrix {
	rows := make([]row, getRowsQuantity(n, key))
	return &matrix{
		rows:    rows,
		columns: key,
	}
}

func getRowsQuantity(n, key int) int {
	if n%key == 0 {
		return n / key
	} else {
		return n / (key - 1)
	}
}

func (m *matrix) AddToMatrix(symbol rune) {
	rows := m.rows
	if m.currentColumn == m.columns {
		m.currentColumn = 0
		m.currentRow++
	}
	rows[m.currentRow] = append(rows[m.currentRow], symbol)
	m.currentColumn++
	m.size++
}

func (m matrix) TransformMatrix() *matrix {
	line := make([]rune, m.size)
	newRow := make([]rune, m.columns)
	newRows := make([]row, m.currentRow+1)
	n := 0
	for i := 0; i < m.columns; i++ {
		for j := 0; j <= m.currentRow; j++ {
			line[n] = m.rows[j][i]
			n++
		}
	}
	n = 0
	for i := 0; i <= m.currentRow; i++ {
		for j := 0; j < m.columns; j++ {
			newRow[j] = line[n]
			n++
		}
		newRows[i] = make([]rune, m.columns)
		copy(newRows[i], newRow)
	}
	return &matrix{
		rows: newRows,
		size: m.size,
	}
}

func (m matrix) String() string {
	sb := strings.Builder{}
	for _, row := range m.rows {
		for _, symbol := range row {
			sb.WriteRune(symbol)
		}
	}
	return sb.String()
}
