package cell

type State struct {
	Graph          *Graph
	Map            *Map
	LastIterations int
}

type Display struct {
	Row  int
	Col  int
	Sign SignType
}

func NewState(size, xMult int, seed func(*Graph, *Map)) *State {
	m := NewMap(size, xMult)
	g := NewGraph()

	seed(g, m)

	g.FixSize()
	return &State{
		Graph: g,
		Map:   m,
	}
}

func GetNodeDisplay(n *Node) (*Display, error) {
	// if n.Row%2 != 0 || n.Col%2 != 0 {
	// 	return nil, errors.New("node has wrong coords")
	// }

	return &Display{
		Row:  n.Row,
		Col:  n.Col,
		Sign: signNode,
	}, nil
}

func GetNodeEmptyDisplay(n *Node) (*Display, error) {
	// if n.Row%2 != 0 || n.Col%2 != 0 {
	// 	return nil, errors.New("node has wrong coords")
	// }

	return &Display{
		Row:  n.Row,
		Col:  n.Col,
		Sign: signEmpty,
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

		// if col%2 == 0 {
		// 	return nil, errors.New("edge has wrong coords")
		// }

		display = &Display{
			Row:  n1.Row,
			Col:  col,
			Sign: signEdgeHorizontal,
		}
	} else {
		var row int
		if (n2.Row - n1.Row) > 0 {
			row = n1.Row + 1
		} else {
			row = n2.Row + 1
		}

		// if row%2 == 0 {
		// 	return nil, errors.New("edge has wrong coords")
		// }

		display = &Display{
			Row:  row,
			Col:  n1.Col,
			Sign: signEdgeVertical,
		}
	}

	return display, nil
}

func (s *State) Mutate() {
	s.LastIterations = 0
	if s.Graph.Size() == 0 {
		return
	}

	visited := make(map[*Node]struct{}, s.Graph.Size())
	nextGraph := s.Graph.Copy()
	nextMap := s.Map.Copy()

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

	s.clearWithNeighbours(nextGraph.ClearCycles(), nextMap)
	s.drawChildren(nextGraph, nextMap)

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

		ds, err := GetNeighboursEmptyDisplays(n)
		if err != nil {
			panic(err)
		}

		for _, d := range ds {
			dest.Set(d.Row, d.Col, d.Sign)
		}
	}

}

func GetNeighboursEmptyDisplays(n *Node) ([]*Display, error) {
	// if n.Row%2 != 0 || n.Col%2 != 0 {
	// 	return nil, errors.New("node has wrong coords")
	// }

	return []*Display{
		{
			Row: n.Row,
			Col: n.Col + 1,
			Sign: signEmpty,
		},
		{
			Row: n.Row + 1,
			Col: n.Col,
			Sign: signEmpty,
		},
		{
			Row: n.Row,
			Col: n.Col - 1,
			Sign: signEmpty,
		},
		{
			Row: n.Row - 1,
			Col: n.Col,
			Sign: signEmpty,
		},
	}, nil
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

	if s.Map.IsEmpty(n.Row, n.Col+2) {
		children = append(children, dest.AddNode(n.Row, n.Col+2))
	}

	if s.Map.IsEmpty(n.Row+2, n.Col) {
		children = append(children, dest.AddNode(n.Row+2, n.Col))
	}

	if s.Map.IsEmpty(n.Row, n.Col-2) {
		children = append(children, dest.AddNode(n.Row, n.Col-2))
	}

	if s.Map.IsEmpty(n.Row-2, n.Col) {
		children = append(children, dest.AddNode(n.Row-2, n.Col))
	}

	for _, child := range children {
		dest.AddEdge(n, child)
	}
}

func (g *Graph) FixSize() {
	g.size = len(g.Nodes)
}

func (g *Graph) Size() int {
	return g.size
}
