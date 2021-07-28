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

// func (c *Cell) WithCentralSeed() *Cell {
// 	n, _ := graph.NewNode()
// 	_ = c.applySeed(&seed{
// 		graph: &graph.Graph{
// 			Nodes: map[graph.NodeID]*graph.Node{
// 				n.ID: n,
// 			},
// 		},
// 		mapping: map[graph.NodeID]Coords{
// 			n.ID: {
// 				row: c.size / 2,
// 				col: c.size / 2,
// 			},
// 		},
// 	})
//
// 	_ = c.getState()
//
// 	return c
// }
//
// func (c *Cell) WithCentralBoundedSeed(size int) *Cell {
// 	pos := c.size / 2
//
// 	m := cellMap.New(c.size)
// 	for i := 0; i < size; i++ {
// 		m.Set(pos+size, pos+i, cellMap.SymbolObstacle)
// 		m.Set(pos+size, pos-i, cellMap.SymbolObstacle)
// 		m.Set(pos-size, pos+i, cellMap.SymbolObstacle)
// 		m.Set(pos-size, pos-i, cellMap.SymbolObstacle)
// 	}
//
// 	for i := 0; i < size; i++ {
// 		m.Set(pos+i, pos+size, cellMap.SymbolObstacle)
// 		m.Set(pos+i, pos-size, cellMap.SymbolObstacle)
// 		m.Set(pos-i, pos+size, cellMap.SymbolObstacle)
// 		m.Set(pos-i, pos-size, cellMap.SymbolObstacle)
// 	}
//
// 	m.Set(pos+size, pos, cellMap.SymbolEmpty)
//
// 	n, _ := graph.NewNode()
//
// 	_ = c.applySeed(&seed{
// 		graph: &graph.Graph{
// 			Nodes: map[graph.NodeID]*graph.Node{
// 				n.ID: n,
// 			},
// 		},
// 		mapping: map[graph.NodeID]Coords{
// 			n.ID: {
// 				row: pos,
// 				col: pos,
// 			},
// 		},
// 		features: m,
// 	})
//
// 	_ = c.getState()
//
// 	return c
// }
//
// func (c *Cell) WithCustomSeed(g *graph.Graph, repr map[graph.NodeID]Coords, m *cellMap.Map) (*Cell, error) {
// 	if g == nil {
// 		return nil, errors.New("graph can't be nil")
// 	}
//
// 	if repr == nil || !c.validateRepr(g, repr) {
// 		return nil, errors.New("invalid repr")
// 	}
//
// 	s := &seed{
// 		graph:    g,
// 		mapping:  repr,
// 		features: m,
// 	}
//
// 	err := c.applySeed(s)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return c, c.getState()
// }
