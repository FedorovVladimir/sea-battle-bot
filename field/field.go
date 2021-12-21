package field

import "sea-battle-bot/cell"

const n = 10

type field struct {
	cells [][]*cell.Cell
}

func newField() *field {
	cells := make([][]*cell.Cell, n)
	for i := 0; i < n; i++ {
		cells = append(cells, make([]*cell.Cell, n))
		for j := 0; j < n; j++ {
			cells[i] = append(cells[i], cell.NewCell())
		}
	}
	return &field{cells: cells}
}

func (f *field) shot(row, col int) {
	f.cells[row][col].Shot()
}
