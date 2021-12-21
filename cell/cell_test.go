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
			fields: fields{state: Sea},
			want:   RuinsOfSea,
		},
		{
			name:   "shot to ship",
			fields: fields{state: Ship},
			want:   RuinsOfShip,
		},
		{
			name:   "shot to ruins on sea",
			fields: fields{state: RuinsOfSea},
			want:   RuinsOfSea,
		},
		{
			name:   "shot to ruins on ship",
			fields: fields{state: RuinsOfShip},
			want:   RuinsOfShip,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cell{state: tt.fields.state}
			c.shot()
			if c.state != tt.want {
				t.Errorf("%s got cell state: %s, but want: %s\n", tt.name, c.state, tt.want)
			}
		})
	}
}
