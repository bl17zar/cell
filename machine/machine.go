package machine

import (
	"time"

	"github.com/bl17zar/cell/cell"
	"github.com/bl17zar/cell/drawers"
)

type Machine struct {
	Cell       *cell.Cell
	generation int
	drawer     drawers.Drawer
}

func NewMachine(cellSize, xMult int, seed func(*cell.Graph, *cell.Map), features []*cell.Display) *Machine {
	return &Machine{
		Cell:   cell.NewCell(cellSize, xMult, seed, features),
		drawer: &drawers.ConsoleDrawer{},
	}
}

func (m *Machine) Run() {
	m.drawer.Draw(m.Cell.State.Map)

	t := time.NewTicker(time.Second)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			if m.generation%2 == 0 {
				m.Cell.Evolve()
			} else {
				m.Cell.ClearCycles()
			}

			m.drawer.Draw(m.Cell.State.Map)
			m.generation++
		}
	}
}
