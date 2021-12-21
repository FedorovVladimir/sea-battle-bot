package cell

type cellState string

const (
	Sea         = cellState("SEA")
	RuinsOfSea  = cellState("RUINS_OF_SEA")
	Ship        = cellState("SHIP")
	RuinsOfShip = cellState("RUINS_OF_SHIP")
)

type cell struct {
	state cellState
}

func (c *cell) shot() {
	if c.state == Sea {
		c.state = RuinsOfSea
	}
	if c.state == Ship {
		c.state = RuinsOfShip
	}
}
