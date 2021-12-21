package cell

type State string

const (
	sea         = State("SEA")
	ruinsOfSea  = State("RUINS_OF_SEA")
	Ship        = State("SHIP")
	ruinsOfShip = State("RUINS_OF_SHIP")
)

type Cell struct {
	state State
}

func NewCell() *Cell {
	return &Cell{state: sea}
}

func (c *Cell) Shot() bool {
	if c.state == Ship {
		c.state = ruinsOfShip
		return true
	}
	if c.state == sea {
		c.state = ruinsOfSea
	}
	return false
}

func (c *Cell) SetShip() {
	c.state = Ship
}

func (c *Cell) GetState() State {
	return c.state
}
