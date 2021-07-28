package main

import (
	"time"

	"github.com/bl17zar/cell/internal/cell"
	"github.com/bl17zar/cell/internal/drawer"
	"github.com/bl17zar/cell/internal/machine"
)

const (
	cellSize = 41
	ww       = 2
)

func main() {
	m := machine.Machine{
		Cell:   cell.NewWithCentralSeed(cellSize),
		Drawer: drawer.NewConsoleDrawer(ww),
		Speed:  time.Second,
	}

	m.Run()
}
