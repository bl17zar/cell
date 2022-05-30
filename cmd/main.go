package main

import (
	"time"

	"github.com/bl17zar/cell/internal/cell"
	"github.com/bl17zar/cell/internal/drawer"
	"github.com/bl17zar/cell/internal/machine"
)

const (
	cellSize     = 33
	boundarySize = 7
)

func main() {
	m := machine.Machine{
		Cell: cell.NewWithCentralBoundedSeed(cellSize, boundarySize),
		Drawer: drawer.NewConsoleDrawer(2, &drawer.Opts{
			UseClear:  true,
			ClearSize: cellSize + 1, // + 1 to also clear stats
		}),
		Speed: time.Second,
	}

	m.Run()
}
