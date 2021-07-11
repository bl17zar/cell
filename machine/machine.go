package machine

import (
	"fmt"
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

	t := time.NewTicker(time.Millisecond * 100)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			m.Cell.Evolve()

			m.drawer.Draw(m.Cell.State.Map)

			fmt.Println("generation:", m.generation)

			m.generation++
		}
	}
}
