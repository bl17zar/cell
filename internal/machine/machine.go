package machine

import (
	"fmt"
	"time"

	"github.com/bl17zar/cell/internal/cell"
	"github.com/bl17zar/cell/internal/drawer"
)

type Machine struct {
	Cell   *cell.Cell
	Drawer drawer.Drawer
	Speed  time.Duration
}

func (m *Machine) Run() {
	t := time.NewTicker(m.Speed)
	defer t.Stop()

	m.Drawer.Draw(m.Cell.State)

	for {
		select {
		case <-t.C:
			tStart := time.Now()
			err := m.Cell.Evolve()
			if err != nil {
				panic(err)
			}
			tElapsed := time.Since(tStart)

			fmt.Println(fmt.Sprint("generation: ", m.Cell.Age, " | elapsed for evolve: ", tElapsed))
			if m.Cell.IsCycled() {
				fmt.Println(fmt.Sprint("cycle time: ", m.Cell.CycledAge(), " | cycle size: ", m.Cell.CycleSize()))
			}

			m.Drawer.Draw(m.Cell.State)
		}
	}
}
