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
		n.ID: {
			Row: size / 2,
			Col: size / 2,
		},
	}

	newC, _ := New(state, g, m)

	return newC
}

func NewWithDownShiftedSeed(size int) *Cell {
	state := cellMap.New(size)

	g := &graph.Graph{}
	n, _ := graph.NewNode()
	_ = g.AddNodes(n)

	m := map[graph.NodeID]*Coords{
		n.ID: {
			Row: size - 1,
			Col: size / 2,
		},
	}

	newC, _ := New(state, g, m)

	return newC
}

func NewWithCentralBoundedSeed(size, obstacleSize int) *Cell {
	state := cellMap.New(size)

	g := &graph.Graph{}
	n, _ := graph.NewNode()
	_ = g.AddNodes(n)

	pos := size / 2
	m := map[graph.NodeID]*Coords{
		n.ID: {
			Row: pos,
			Col: pos,
		},
	}

	for i := 0; i < obstacleSize; i++ {
		state.AddObstacle(pos+obstacleSize, pos+i)
		state.AddObstacle(pos+obstacleSize, pos-i)
		state.AddObstacle(pos-obstacleSize, pos+i)
		state.AddObstacle(pos-obstacleSize, pos-i)
	}

	for i := 0; i < obstacleSize; i++ {
		state.AddObstacle(pos+i, pos+obstacleSize)
		state.AddObstacle(pos+i, pos-obstacleSize)
		state.AddObstacle(pos-i, pos+obstacleSize)
		state.AddObstacle(pos-i, pos-obstacleSize)
	}

	state.SetEmpty(pos, pos+obstacleSize)

	newC, _ := New(state, g, m)

	return newC
}
