package field

import (
	"fmt"
	"sea-battle-bot/cell"
	"strings"
)

const n = 10

type field struct {
	cells [][]*cell.Cell
}

func NewField() *field {
	cells := make([][]*cell.Cell, n)
	for i := 0; i < n; i++ {
		cells[i] = make([]*cell.Cell, n)
		for j := 0; j < n; j++ {
			cells[i][j] = cell.NewCell()
		}
	}
	return &field{cells: cells}
}

func (f *field) shot(row, col int) bool {
	return f.cells[row][col].Shot()
}

func (f *field) setShip(row, col int) {
	f.cells[row][col].SetShip()
}

func (f *field) inBattle() bool {
	for _, row := range f.cells {
		for _, c := range row {
			if c.GetState() == cell.Ship {
				return true
			}
		}
	}
	return false
}

func (f *field) String() string {
	s := strings.Builder{}
	for _, row := range f.cells {
		for _, c := range row {
			s.Write([]byte(fmt.Sprintf("%s ", c)))
		}
		s.Write([]byte("\n"))
	}
	return s.String()
}
