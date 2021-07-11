package machine

import (
	"fmt"
	"time"

	"github.com/bl17zar/cell/cell"
)

type Machine struct {
	Cells      []*cell.Cell
	generation int
	frame      int
}

func NewMachine(cellSize, xMult int, seed func(*cell.Graph, *cell.Map)) *Machine {
	return &Machine{
		Cells: []*cell.Cell{cell.NewCell(cellSize, xMult, seed)},
	}
}

func (m *Machine) Run() {
	m.Draw()

	t := time.NewTicker(time.Second)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			for _, c := range m.Cells {
				if m.frame%2 == 0 {
					c.Evolve()
				} else {
					c.ClearCycles()
				}
			}

			m.Draw()
			m.frame++
		}
	}
}

func (m *Machine) Draw() {
	fmt.Println(fmt.Sprintf("generation: %d/%d, iterations: %d", m.generation, m.frame%2, m.Cells[0].State.LastIterations))

	for _, c := range m.Cells {
		fmt.Print(c.State.Map)
	}

	m.generation++
}
