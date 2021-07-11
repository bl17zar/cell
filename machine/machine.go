package machine

import (
	"time"

	"github.com/bl17zar/cell/cell"
	"github.com/bl17zar/cell/drawers"
)

type Machine struct {
	Cells      []*cell.Cell
	generation int
	drawer     drawers.Drawer
}

func NewMachine(cellSize, xMult int, seed func(*cell.Graph, *cell.Map), features []*cell.Display) *Machine {
	return &Machine{
		Cells:  []*cell.Cell{cell.NewCell(cellSize, xMult, seed, features)},
		drawer: &drawers.ConsoleDrawer{},
	}
}

func (m *Machine) Run() {
	m.drawer.Draw(m.Cells[0].State.Map)

	t := time.NewTicker(time.Second)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			for _, c := range m.Cells {
				if m.generation%2 == 0 {
					c.Evolve()
				} else {
					c.ClearCycles()
				}
			}

			m.drawer.Draw(m.Cells[0].State.Map)
			m.generation++
		}
	}
}
