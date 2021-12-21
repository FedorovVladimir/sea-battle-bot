package cell

import (
	"testing"
)

func Test_cell_shot(t *testing.T) {
	type fields struct {
		state cellState
	}
	tests := []struct {
		name   string
		fields fields
		want   cellState
	}{
		{
			name:   "shot to sea",
			fields: fields{state: sea},
			want:   ruinsOfSea,
		},
		{
			name:   "shot to ship",
			fields: fields{state: ship},
			want:   ruinsOfShip,
		},
		{
			name:   "shot to ruins on sea",
			fields: fields{state: ruinsOfSea},
			want:   ruinsOfSea,
		},
		{
			name:   "shot to ruins on ship",
			fields: fields{state: ruinsOfShip},
			want:   ruinsOfShip,
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
