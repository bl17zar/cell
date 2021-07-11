package cell

import (
	"fmt"
)

type Node struct {
	Row int
	Col int
}

func (n *Node) id() string {
	return fmt.Sprintf("%v,%v", n.Row, n.Col)
}

type Graph struct {
	size  int
	Nodes map[string]*Node
	Edges map[string][]*Node
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]*Node),
		Edges: make(map[string][]*Node),
	}
}

func (g *Graph) Copy() *Graph {
	res := NewGraph()

	for k, v := range g.Nodes {
		res.Nodes[k] = v
	}

	for k, v := range g.Edges {
		res.Edges[k] = v
	}

	return res
}

func (g *Graph) AddNode(row, col int) *Node {
	n := &Node{row, col}
	g.Nodes[n.id()] = n

	return n
}

func (g *Graph) AddEdge(n1, n2 *Node) {
	g.Edges[n1.id()] = append(g.Edges[n1.id()], n2)
	g.Edges[n2.id()] = append(g.Edges[n2.id()], n1)
}

func (g *Graph) DeleteNode(n *Node) {
	delete(g.Nodes, n.id())

	if g.Edges == nil {
		return
	}

	for _, linkedNode := range g.Edges[n.id()] {
		if len(g.Edges[linkedNode.id()]) == 0 {
			continue
		}

		newLinkedNodeNodes := make([]*Node, 0, len(g.Edges[linkedNode.id()])-1)
		for _, linkedNodeNode := range g.Edges[linkedNode.id()] {
			if linkedNodeNode == n {
				continue
			}

			newLinkedNodeNodes = append(newLinkedNodeNodes, linkedNodeNode)
		}

		g.Edges[linkedNode.id()] = newLinkedNodeNodes
	}

	delete(g.Edges, n.id())
}

func (g *Graph) ClearCycles(nodes []*Node) {
	for _, n := range nodes {
		g.DeleteNode(n)
	}
}

func (g *Graph) FindCycles() []*Node {
	visited := make(map[string]struct{}, len(g.Nodes))
	cycles := [][]string{}

	for id := range g.Nodes {
		if _, ok := visited[id]; ok {
			continue
		}

		cycles = append(cycles, g.findCycles(id, newOrderedSet(), visited)...)
	}

	cycleNodes := map[string]struct{}{}
	res := []*Node{}
	for _, cycle := range cycles {
		for _, nodeId := range cycle {
			if _, ok := g.Nodes[nodeId]; ok {
				if _, ok := cycleNodes[nodeId]; !ok {
					cycleNodes[nodeId] = struct{}{}
					res = append(res, g.Nodes[nodeId])
				}
			}
		}
	}

	return res
}

func (g *Graph) findCycles(nodeID string, parents orderedSet, visited map[string]struct{}) (cycles [][]string) {
	visited[nodeID] = struct{}{}

	cycles = [][]string{}

	for _, linkedNode := range g.Edges[nodeID] {
		if parents.GetLast() == linkedNode.id() {
			continue
		}

		if parents.HasNotLast(linkedNode.id()) {
			cycle := parents.Unwind(linkedNode.id())
			cycle = append(cycle, nodeID)

			cycles = append(cycles, cycle)

			continue
		}

		if _, ok := visited[linkedNode.id()]; ok {
			continue
		}

		cycles = append(cycles, g.findCycles(linkedNode.id(), parents.Put(nodeID), visited)...)
	}

	return
}
