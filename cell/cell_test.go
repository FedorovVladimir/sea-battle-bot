package cell

import (
	"testing"
)

func Test_cell_shot(t *testing.T) {
	type fields struct {
		state State
	}
	tests := []struct {
		name   string
		fields fields
		want   State
	}{
		{
			name:   "shot to sea",
			fields: fields{state: Sea},
			want:   ruinsOfSea,
		},
		{
			name:   "shot to ship",
			fields: fields{state: Ship},
			want:   RuinsOfShip,
		},
		{
			name:   "shot to ruins on sea",
			fields: fields{state: ruinsOfSea},
			want:   ruinsOfSea,
		},
		{
			name:   "shot to ruins on ship",
			fields: fields{state: RuinsOfShip},
			want:   RuinsOfShip,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cell{state: tt.fields.state}
			c.Shot()
			if c.state != tt.want {
				t.Errorf("%s got cell state: %s, but want: %s\n", tt.name, c.state, tt.want)
			}
		})
	}
}

func TestCell_setShip(t *testing.T) {
	type fields struct {
		state State
	}
	tests := []struct {
		name   string
		fields fields
		want   State
	}{
		{
			name:   "set ship to see",
			fields: fields{state: Sea},
			want:   Ship,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cell{state: tt.fields.state}
			c.SetShip()
			if c.state != tt.want {
				t.Errorf("%s got cell state: %s, but want: %s\n", tt.name, c.state, tt.want)
			}
		})
	}
}
