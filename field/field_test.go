package field

import (
	"sea-battle-bot/cell"
	"testing"
)

func Test_field_shot(t *testing.T) {
	type fields struct {
		cells [][]*cell.Cell
	}
	type args struct {
		row int
		col int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "shot to sea (0;0)",
			fields: fields{cells: [][]*cell.Cell{{cell.NewCell()}}},
			args:   args{row: 0, col: 0},
			want:   false,
		},
		{
			name:   "shot to ship (0;0)",
			fields: fields{cells: [][]*cell.Cell{{cell.NewCell(cell.WithState(cell.Ship))}}},
			args:   args{row: 0, col: 0},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &field{cells: tt.fields.cells}
			got := f.Shot(tt.args.row, tt.args.col)
			if got != tt.want {
				t.Errorf("%s got shot result: %v, but want: %v\n", tt.name, got, tt.want)
			}
		})
	}
}

func Test_field_setShip(t *testing.T) {
	type fields struct {
		cells [][]*cell.Cell
	}
	type args struct {
		row int
		col int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   cell.State
	}{
		{
			name:   "set ship to sea (0;0)",
			fields: fields{cells: [][]*cell.Cell{{cell.NewCell()}}},
			args:   args{row: 0, col: 0},
			want:   cell.Ship,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &field{cells: tt.fields.cells}
			f.SetShip(tt.args.row, tt.args.col)
			got := f.cells[tt.args.row][tt.args.col].GetState()
			if got != tt.want {
				t.Errorf("%s got cell state: %s, but want: %s\n", tt.name, got, tt.want)
			}
		})
	}
}

func Test_field_inBattle(t *testing.T) {
	type fields struct {
		cells [][]*cell.Cell
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "start game",
			fields: fields{cells: [][]*cell.Cell{
				{cell.NewCell(cell.WithState(cell.Ship))},
				{cell.NewCell(cell.WithState(cell.Sea))},
			}},
			want: true,
		},
		{
			name: "middle game",
			fields: fields{cells: [][]*cell.Cell{
				{cell.NewCell(cell.WithState(cell.RuinsOfShip))},
				{cell.NewCell(cell.WithState(cell.Ship))},
			}},
			want: true,
		},
		{
			name: "end game",
			fields: fields{cells: [][]*cell.Cell{
				{cell.NewCell(cell.WithState(cell.RuinsOfShip))},
				{cell.NewCell(cell.WithState(cell.Sea))},
			}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &field{cells: tt.fields.cells}
			if got := f.inBattle(); got != tt.want {
				t.Errorf("inBattle() = %v, want %v", got, tt.want)
			}
		})
	}
}
