package main

import (
	"time"

	"github.com/bl17zar/cell/internal/cell"
	"github.com/bl17zar/cell/internal/drawer"
	"github.com/bl17zar/cell/internal/machine"
)

const (
	cellSize = 21
	ww       = 2
)

func main() {
	m := machine.Machine{
		Cell:   cell.NewWithDownShiftedSeed(cellSize),
		Drawer: drawer.NewConsoleDrawer(ww),
		Speed:  time.Second,
	}

	m.Run()
}
