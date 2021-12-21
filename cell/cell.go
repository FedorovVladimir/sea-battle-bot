package cell

type State string

const (
	Sea         = State("SEA")
	ruinsOfSea  = State("RUINS_OF_SEA")
	Ship        = State("SHIP")
	RuinsOfShip = State("RUINS_OF_SHIP")
)

type Cell struct {
	state State
}

type Option func(*Cell)

func WithState(state State) Option {
	return func(cell *Cell) {
		cell.state = state
	}
}

func NewCell(setters ...Option) *Cell {
	cell := &Cell{state: Sea}
	for _, setter := range setters {
		setter(cell)
	}
	return cell
}

func (c *Cell) Shot() bool {
	if c.state == Ship {
		c.state = RuinsOfShip
		return true
	}
	if c.state == Sea {
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

func (c *Cell) String() string {
	return string(c.state)
}
