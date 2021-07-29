package cell

import (
	"errors"

	"github.com/bl17zar/cell/internal/graph"
	cellMap "github.com/bl17zar/cell/internal/map"
)

type Coords struct {
	Row int
	Col int
}

type Cell struct {
	isCycled  bool
	cycleFrom int
	cycleSize int
	cycledAge int
	Age       int
	graph     *graph.Graph
	mapping   map[graph.NodeID]*Coords
	State     *cellMap.Map
	history   *history
}

func New(state *cellMap.Map, g *graph.Graph, m map[graph.NodeID]*Coords) (*Cell, error) {
	if state == nil || g == nil || m == nil {
		return nil, errors.New("failed to init cell")
	}

	return &Cell{
		graph:   g,
		mapping: m,
		State:   state,
		history: newHistory(),
	}, nil
}

func (c *Cell) startNewAge() error {
	if !c.isCycled {
		if !c.checkIsCycled(c.State) {
			c.history.Put(c.State.Copy())
		}

		if err := c.clearCycles(); err != nil {
			return err
		}
	}

	c.Age++

	return nil
}

func (c *Cell) clearCycles() error {
	for n := range c.graph.Cycles {
		coords := c.mapping[n.ID]
		for _, neighbour := range n.Neighbours() {
			neighbourCoords := c.mapping[neighbour.ID]
			c.State.SetEmptyEdge(coords.Row, coords.Col, neighbourCoords.Row, neighbourCoords.Col)
		}
		c.State.SetEmpty(coords.Col, coords.Row)
	}

	var forDelete []*graph.Node
	for n := range c.graph.Cycles {
		forDelete = append(forDelete, n)
		delete(c.mapping, n.ID)
	}

	if err := c.graph.DeleteNodes(forDelete...); err != nil {
		return err
	}

	return nil
}

func (c *Cell) Evolve() error {
	if err := c.startNewAge(); err != nil {
		return err
	}

	if !c.isCycled {
		g, m, err := c.produceNextGeneration()
		if err != nil {
			return err
		}

		newState, err := c.newState(g, m)
		if err != nil {
			return err
		}

		c.State = newState
		c.graph = g
		c.mapping = m

		return nil
	}

	cycleStateIdx := c.cycleFrom + (c.Age-c.cycleFrom)%c.cycleSize
	state, err := c.history.GetByIdx(cycleStateIdx)
	if err != nil {
		return nil
	}

	c.State = state.(*cellMap.Map)
	c.cycledAge = cycleStateIdx - c.cycleFrom
	c.graph = nil
	c.mapping = nil

	return nil
}

func (c *Cell) newState(g *graph.Graph, mapping map[graph.NodeID]*Coords) (*cellMap.Map, error) {
	newState := c.State.Copy()

	for _, n := range g.Nodes {
		coords := mapping[n.ID]
		newState.AddNode(coords.Row, coords.Col)
		for _, neighbour := range n.Neighbours() {
			neighbourCoords := mapping[neighbour.ID]
			newState.AddEdge(coords.Row, coords.Col, neighbourCoords.Row, neighbourCoords.Col)
		}
	}

	for n := range g.Cycles {
		coords := mapping[n.ID]
		newState.AddCycleNode(coords.Row, coords.Col)
		for _, neighbour := range n.Neighbours() {
			neighbourCoords := mapping[neighbour.ID]
			newState.AddCycleEdge(coords.Row, coords.Col, neighbourCoords.Row, neighbourCoords.Col)
		}
	}

	return newState, nil
}

func (c *Cell) produceNextGeneration() (*graph.Graph, map[graph.NodeID]*Coords, error) {
	g := c.graph.Copy()

	m := make(map[graph.NodeID]*Coords, len(c.mapping))
	for k := range c.mapping {
		m[k] = &Coords{
			Row: c.mapping[k].Row,
			Col: c.mapping[k].Col,
		}
	}

	for nodeID := range c.graph.Nodes {
		n := c.graph.Nodes[nodeID]

		if err := c.giveBirth(n, g, m); err != nil {
			return nil, nil, err
		}
	}

	if err := c.mergeNodeDuplicates(g, &m); err != nil {
		return nil, nil, err
	}

	g.FindCycles()

	return g, m, nil
}

func (c *Cell) giveBirth(n *graph.Node, destG *graph.Graph, destM map[graph.NodeID]*Coords) error {
	if len(c.graph.Nodes[n.ID].Neighbours()) == graph.MaxNeighboursNum {
		return nil
	}

	parentCoords := c.mapping[n.ID]
	children := make([]*graph.Node, 0, graph.MaxNeighboursNum-len(c.graph.Nodes[n.ID].Neighbours()))

	if c.State.CheckPathForNewNode(parentCoords.Row, parentCoords.Col, cellMap.LEFT) {
		newN, err := graph.NewNode(destG.Nodes[n.ID])
		if err != nil {
			return err
		}

		children = append(children, newN)
		destM[newN.ID] = &Coords{
			Row: parentCoords.Row,
			Col: parentCoords.Col - 2,
		}
	}

	if c.State.CheckPathForNewNode(parentCoords.Row, parentCoords.Col, cellMap.UP) {
		newN, err := graph.NewNode(destG.Nodes[n.ID])
		if err != nil {
			return err
		}

		children = append(children, newN)
		destM[newN.ID] = &Coords{
			Row: parentCoords.Row - 2,
			Col: parentCoords.Col,
		}
	}

	if c.State.CheckPathForNewNode(parentCoords.Row, parentCoords.Col, cellMap.DOWN) {
		newN, err := graph.NewNode(destG.Nodes[n.ID])
		if err != nil {
			return err
		}

		children = append(children, newN)
		destM[newN.ID] = &Coords{
			Row: parentCoords.Row + 2,
			Col: parentCoords.Col,
		}
	}

	if c.State.CheckPathForNewNode(parentCoords.Row, parentCoords.Col, cellMap.RIGHT) {
		newN, err := graph.NewNode(destG.Nodes[n.ID])
		if err != nil {
			return err
		}

		children = append(children, newN)
		destM[newN.ID] = &Coords{
			Row: parentCoords.Row,
			Col: parentCoords.Col + 2,
		}
	}

	return destG.AddNodes(children...)
}

func (c *Cell) mergeNodeDuplicates(g *graph.Graph, m *map[graph.NodeID]*Coords) error {
	duplicates := map[Coords][]graph.NodeID{}
	for nID, coords := range *m {
		coordsVal := Coords{Row: coords.Row, Col: coords.Col}
		duplicates[coordsVal] = append(duplicates[coordsVal], nID)
	}

	for _, v := range duplicates {
		if len(v) > 1 {
			sNID := v[0]
			for _, dID := range v {
				if dID != sNID {
					sN := g.Nodes[sNID]
					dN := g.Nodes[dID]

					for _, dNeighbour := range dN.Neighbours() {
						if err := dNeighbour.RemoveNeighbourByID(dN.ID); err != nil {
							return err
						}

						if err := dNeighbour.AddNeighbours(sN); err != nil {
							return err
						}

						if err := sN.AddNeighbours(dNeighbour); err != nil {
							return err
						}
					}

					delete(*m, dID)
					if err := g.DeleteNodes(dN); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func (c *Cell) checkIsCycled(state *cellMap.Map) bool {
	if c.history.Exists(state) {
		c.isCycled = true

		d := c.history.Distance(state)
		if d == -1 {
			panic("wrong history cycle detection")
		}

		c.cycleSize = c.Age - d
		c.cycleFrom = d

		return true
	}

	return false
}

func (c *Cell) IsCycled() bool {
	return c.isCycled
}

func (c *Cell) CycledAge() int {
	return c.cycledAge
}

func (c *Cell) CycleFrom() int {
	return c.cycleFrom
}

func (c *Cell) CycleSize() int {
	return c.cycleSize
}
