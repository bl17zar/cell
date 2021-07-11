package cell

type State struct {
	Graph           *Graph
	Map             *Map
	LastIterations  int
	undeletedCycles []*Node
}

type Display struct {
	Row  int
	Col  int
	Sign SignType
}

func NewState(size, xMult int, seed func(*Graph, *Map), features []*Display) *State {
	m := NewMap(size, xMult, features)
	g := NewGraph()

	seed(g, m)

	g.FixSize()
	return &State{
		Graph: g,
		Map:   m,
	}
}

func GetNodeDisplay(n *Node) (*Display, error) {
	return &Display{
		Row:  n.Row,
		Col:  n.Col,
		Sign: SignNode,
	}, nil
}

func GetCycleNodeDisplay(n *Node) (*Display, error) {
	return &Display{
		Row:  n.Row,
		Col:  n.Col,
		Sign: SignCycleNode,
	}, nil
}

func GetNodeEmptyDisplay(n *Node) (*Display, error) {
	return &Display{
		Row:  n.Row,
		Col:  n.Col,
		Sign: SignEmpty,
	}, nil
}

func GetEdgeDisplay(n1, n2 *Node) (*Display, error) {
	var display *Display

	if n1.Row == n2.Row {
		var col int
		if (n2.Col - n1.Col) > 0 {
			col = n1.Col + 1
		} else {
			col = n2.Col + 1
		}

		display = &Display{
			Row:  n1.Row,
			Col:  col,
			Sign: SignEdgeHorizontal,
		}
	} else {
		var row int
		if (n2.Row - n1.Row) > 0 {
			row = n1.Row + 1
		} else {
			row = n2.Row + 1
		}

		display = &Display{
			Row:  row,
			Col:  n1.Col,
			Sign: SignEdgeVertical,
		}
	}

	return display, nil
}

func GetCycleEdgeDisplay(n1, n2 *Node) (*Display, error) {
	var display *Display

	if n1.Row == n2.Row {
		var col int
		if (n2.Col - n1.Col) > 0 {
			col = n1.Col + 1
		} else {
			col = n2.Col + 1
		}

		display = &Display{
			Row:  n1.Row,
			Col:  col,
			Sign: SignCycleEdgeHorizontal,
		}
	} else {
		var row int
		if (n2.Row - n1.Row) > 0 {
			row = n1.Row + 1
		} else {
			row = n2.Row + 1
		}

		display = &Display{
			Row:  row,
			Col:  n1.Col,
			Sign: SignCycleEdgeVertical,
		}
	}

	return display, nil
}

func (s *State) Mutate() {
	s.LastIterations = 0
	if s.Graph.Size() == 0 {
		return
	}

	s.Graph.ClearCycles(s.undeletedCycles)
	s.Graph.FixSize()

	visited := make(map[*Node]struct{}, s.Graph.Size())
	nextMap := s.Map.Copy()

	s.clearWithNeighbours(s.undeletedCycles, nextMap)

	nextGraph := s.Graph.Copy()

	for len(visited) < s.Graph.Size() {
		for id := range s.Graph.Nodes {
			n := s.Graph.Nodes[id]

			if _, ok := visited[n]; ok {
				continue
			}

			s.giveBirth(n, nextGraph)

			visited[n] = struct{}{}
			s.LastIterations++
		}
	}

	s.drawChildren(nextGraph, nextMap)

	cycles := nextGraph.FindCycles()
	s.drawCycles(cycles, nextGraph, nextMap)
	s.undeletedCycles = cycles

	s.Graph = nextGraph
	s.Map = nextMap
	s.Graph.FixSize()

	return
}

func (s *State) clearWithNeighbours(nodes []*Node, dest *Map) {
	for _, n := range nodes {
		d, err := GetNodeEmptyDisplay(n)
		if err != nil {
			panic(err)
		}

		dest.Set(d.Row, d.Col, d.Sign)

		ds, err := s.GetNeighboursEmptyDisplays(n)
		if err != nil {
			panic(err)
		}

		for _, d := range ds {
			dest.Set(d.Row, d.Col, d.Sign)
		}
	}

}

func (s *State) GetNeighboursEmptyDisplays(n *Node) ([]*Display, error) {
	res := make([]*Display, 0, 4)

	if s.Map.IsInsideBorders(n.Row, n.Col+1) && !s.Map.IsObstacle(n.Row, n.Col+1) {
		res = append(res, &Display{
			Row:  n.Row,
			Col:  n.Col + 1,
			Sign: SignEmpty,
		})
	}

	if s.Map.IsInsideBorders(n.Row+1, n.Col) && !s.Map.IsObstacle(n.Row+1, n.Col) {
		res = append(res, &Display{
			Row:  n.Row + 1,
			Col:  n.Col,
			Sign: SignEmpty,
		})
	}

	if s.Map.IsInsideBorders(n.Row, n.Col-1) && !s.Map.IsObstacle(n.Row, n.Col-1) {
		res = append(res, &Display{
			Row:  n.Row,
			Col:  n.Col - 1,
			Sign: SignEmpty,
		})
	}

	if s.Map.IsInsideBorders(n.Row-1, n.Col) && !s.Map.IsObstacle(n.Row-1, n.Col) {
		res = append(res, &Display{
			Row:  n.Row - 1,
			Col:  n.Col,
			Sign: SignEmpty,
		})
	}

	return res, nil
}

func (s *State) drawChildren(nextGen *Graph, dest *Map) {
	for n := range nextGen.Nodes {
		d, err := GetNodeDisplay(nextGen.Nodes[n])
		if err != nil {
			panic(err)
		}

		dest.Set(d.Row, d.Col, d.Sign)
	}

	for e := range nextGen.Edges {
		for _, eN := range nextGen.Edges[e] {
			d, err := GetEdgeDisplay(nextGen.Nodes[e], eN)
			if err != nil {
				panic(err)
			}
			dest.Set(d.Row, d.Col, d.Sign)
		}
	}
}

func (s *State) giveBirth(n *Node, dest *Graph) {
	if len(s.Graph.Edges[n.id()]) == 4 {
		return
	}

	children := make([]*Node, 0, len(s.Graph.Edges[n.id()])/2)

	if !s.Map.IsObstacle(n.Row, n.Col+2) && s.Map.IsInsideBorders(n.Row, n.Col+2) {
		children = append(children, dest.AddNode(n.Row, n.Col+2))
	}

	if !s.Map.IsObstacle(n.Row+2, n.Col) && s.Map.IsInsideBorders(n.Row+2, n.Col) {
		children = append(children, dest.AddNode(n.Row+2, n.Col))
	}

	if !s.Map.IsObstacle(n.Row, n.Col-2) && s.Map.IsInsideBorders(n.Row, n.Col-2) {
		children = append(children, dest.AddNode(n.Row, n.Col-2))
	}

	if !s.Map.IsObstacle(n.Row-2, n.Col) && s.Map.IsInsideBorders(n.Row-2, n.Col) {
		children = append(children, dest.AddNode(n.Row-2, n.Col))
	}

	for _, child := range children {
		dest.AddEdge(n, child)
	}
}

func (s *State) drawCycles(cycles []*Node, graph *Graph, dest *Map) {
	for _, n := range cycles {
		d, err := GetCycleNodeDisplay(n)
		if err != nil {
			panic(err)
		}

		dest.Set(d.Row, d.Col, d.Sign)

		for _, eN := range graph.Edges[n.id()] {
			d, err := GetCycleEdgeDisplay(n, eN)
			if err != nil {
				panic(err)
			}
			dest.Set(d.Row, d.Col, d.Sign)
		}

	}
}

func (g *Graph) FixSize() {
	g.size = len(g.Nodes)
}

func (g *Graph) Size() int {
	return g.size
}
