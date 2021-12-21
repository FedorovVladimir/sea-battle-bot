package cell

type cellState string

const (
	sea         = cellState("SEA")
	ruinsOfSea  = cellState("RUINS_OF_SEA")
	ship        = cellState("SHIP")
	ruinsOfShip = cellState("RUINS_OF_SHIP")
)

type Cell struct {
	state cellState
}

func NewCell() *Cell {
	return &Cell{state: sea}
}

func (c *Cell) Shot() bool {
	if c.state == ship {
		c.state = ruinsOfShip
		return true
	}
	if c.state == sea {
		c.state = ruinsOfSea
	}
	return false
}

func (c *Cell) setShip() {
	c.state = ship
}
