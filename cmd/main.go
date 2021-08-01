package main

import (
	"time"

	"github.com/bl17zar/cell/internal/cell"
	"github.com/bl17zar/cell/internal/drawer"
	"github.com/bl17zar/cell/internal/machine"
	"github.com/faiface/pixel/pixelgl"
)

const (
	cellSize     = 32
	boundarySize = 5
)

func main() {
	pixelgl.Run(func() {
		m := machine.Machine{
			Cell:   cell.NewWithCentralBoundedSeed(cellSize, boundarySize),
			Drawer: drawer.NewPixelDrawer(),
			Speed:  time.Second,
		}

		m.Run()
	})
}
