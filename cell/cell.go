package cell

type cellState string

const (
	sea         = cellState("SEA")
	ruinsOfSea  = cellState("RUINS_OF_SEA")
	ship        = cellState("SHIP")
	ruinsOfShip = cellState("RUINS_OF_SHIP")
)

type cell struct {
	state cellState
}

func (c *cell) shot() {
	if c.state == sea {
		c.state = ruinsOfSea
	}
	if c.state == ship {
		c.state = ruinsOfShip
	}
}
