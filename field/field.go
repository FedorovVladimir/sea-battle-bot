package field

import (
	"fmt"
	"sea-battle-bot/cell"
	"strings"
)

type Field struct {
	size  int
	cells [][]*cell.Cell
}

func NewField(size int) *Field {
	cells := make([][]*cell.Cell, size)
	for i := 0; i < size; i++ {
		cells[i] = make([]*cell.Cell, size)
		for j := 0; j < size; j++ {
			cells[i][j] = cell.NewCell()
		}
	}
	return &Field{size: size, cells: cells}
}

func (f *Field) Shot(row, col int) bool {
	if row >= f.size || col >= f.size || row < 0 || col < 0 {
		return false
	}
	return f.cells[row][col].Shot()
}

func (f *Field) SetShip(row, col int) {
	if row >= f.size || col >= f.size || row < 0 || col < 0 {
		return
	}
	f.cells[row][col].SetShip()
}

func (f *Field) InBattle() bool {
	for _, row := range f.cells {
		for _, c := range row {
			if c.GetState() == cell.Ship {
				return true
			}
		}
	}
	return false
}

func (f *Field) String() string {
	s := strings.Builder{}
	for _, row := range f.cells {
		for _, c := range row {
			s.Write([]byte(fmt.Sprintf("%s ", c)))
		}
		s.Write([]byte("\n"))
	}
	return s.String()
}
