package cell

type Cell struct {
	State *State
}

func NewCell(size, xMult int, seed func(*Graph, *Map), features []*Display) *Cell {
	return &Cell{
		State: NewState(size, xMult, seed, features),
	}
}

func (c *Cell) Evolve() {
	c.State.Mutate()
}

func (c *Cell) ClearCycles() {
	c.State.ClearCycles()
}
