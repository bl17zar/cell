package cell

import (
	"github.com/bl17zar/cell/internal/graph"
	cellMap "github.com/bl17zar/cell/internal/map"
)

func NewWithCentralSeed(size int) *Cell {
	state := cellMap.New(size)

	g := &graph.Graph{}
	n, _ := graph.NewNode()
	_ = g.AddNodes(n)

	m := map[graph.NodeID]*Coords{
		n.ID: &Coords{
			Row: size / 2,
			Col: size / 2,
		},
	}

	newC, _ := New(state, g, m)

	return newC
}
