package graph

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

const MaxNeighboursNum = 4

type NodeID string

type Node struct {
	ID         NodeID
	neighbours []*Node
}

func (n *Node) Neighbours() []*Node {
	return n.neighbours
}

func (n *Node) AddNeighbours(newNeighbours ...*Node) error {
	if len(n.neighbours)+len(newNeighbours) > MaxNeighboursNum {
		return errors.New("failed to add neighbours; neighbours size will be greater than max")
	}

	n.neighbours = append(n.neighbours, newNeighbours...)
	return nil
}

func (n *Node) RemoveNeighbourByIdx(idx int) error {
	if idx >= MaxNeighboursNum || idx > len(n.neighbours)-1 {
		return errors.New("failed to remove neighbour by index; index out of range")
	}

	if idx == len(n.neighbours)-1 {
		n.neighbours = n.neighbours[:idx]
		return nil
	}

	n.neighbours = append(n.neighbours[:idx], n.neighbours[idx+1:]...)
	return nil
}

func (n *Node) RemoveNeighbourByID(id NodeID) error {
	for pos, neighbour := range n.neighbours {
		if neighbour.ID == id {
			return n.RemoveNeighbourByIdx(pos)
		}
	}

	return errors.New(fmt.Sprint("neighbour with id: ", id, " not found"))
}

func NewNode(neighbours ...*Node) (*Node, error) {
	n := &Node{
		ID: NodeID(uuid.New().String()),
	}

	if err := n.AddNeighbours(neighbours...); err != nil {
		return nil, err
	}

	return n, nil
}

type Graph struct {
	Nodes  map[NodeID]*Node
	Cycles map[*Node]struct{}
}

func (g *Graph) AddNodes(ns ...*Node) error {
	if g.Nodes == nil {
		g.Nodes = make(map[NodeID]*Node, len(ns))
	}

	for _, n := range ns {
		if _, ok := g.Nodes[n.ID]; ok {
			if err := g.DeleteNodes(ns...); err != nil {
				return err
			}

			return errors.New("id conflict")
		}

		g.Nodes[n.ID] = n

		for _, neighbour := range n.neighbours {
			if _, ok := g.Nodes[neighbour.ID]; !ok {
				return errors.New("wrong neighbour")
			}

			if err := g.Nodes[neighbour.ID].AddNeighbours(n); err != nil {
				return err
			}
		}
	}

	return nil
}

func (g *Graph) DeleteNodes(ns ...*Node) error {
	for _, n := range ns {
		if _, ok := g.Nodes[n.ID]; !ok {
			continue
		}

		for _, neighbour := range n.neighbours {
			selfRefPos := make([]int, 0, 1)
			for i, innerNeighbour := range neighbour.neighbours {
				if innerNeighbour.ID == n.ID {
					selfRefPos = append(selfRefPos, i)
				}
			}

			for _, i := range selfRefPos {
				if err := neighbour.RemoveNeighbourByIdx(i); err != nil {
					return err
				}
			}
		}

		delete(g.Nodes, n.ID)
	}
	return nil
}

func (g *Graph) Copy() *Graph {
	newG := &Graph{
		Nodes: make(map[NodeID]*Node),
	}

	for k := range g.Nodes {
		n := g.Nodes[k]

		newG.Nodes[k] = &Node{ID: n.ID}
	}

	for k := range g.Nodes {
		n := g.Nodes[k]
		newN := newG.Nodes[k]

		for _, neighbour := range n.neighbours {
			newN.neighbours = append(newN.neighbours, newG.Nodes[neighbour.ID])
		}
	}

	return newG
}

func (g *Graph) FindCycles() {
	visited := make(map[NodeID]struct{}, len(g.Nodes))
	cycles := [][]NodeID{}

	for id := range g.Nodes {
		if _, ok := visited[id]; ok {
			continue
		}

		cycles = append(cycles, g.findCycles(id, newNodeSet(), visited)...)
	}

	cycleNodes := map[NodeID]struct{}{}
	g.Cycles = map[*Node]struct{}{}
	for _, cycle := range cycles {
		for _, nodeId := range cycle {
			if _, ok := g.Nodes[nodeId]; ok {
				if _, ok := cycleNodes[nodeId]; !ok {
					cycleNodes[nodeId] = struct{}{}
					g.Cycles[g.Nodes[nodeId]] = struct{}{}
				}
			}
		}
	}
}

func (g *Graph) findCycles(nodeID NodeID, parents nodeSet, visited map[NodeID]struct{}) (cycles [][]NodeID) {
	visited[nodeID] = struct{}{}

	cycles = [][]NodeID{}

	for _, neighbour := range g.Nodes[nodeID].neighbours {
		if parents.GetLast() == neighbour.ID {
			continue
		}

		if parents.HasNotLast(neighbour.ID) {
			cycle := parents.Unwind(neighbour.ID)
			cycle = append(cycle, nodeID)
			cycles = append(cycles, cycle)
			continue
		}

		if _, ok := visited[neighbour.ID]; ok {
			continue
		}

		cycles = append(cycles, g.findCycles(neighbour.ID, parents.Put(nodeID), visited)...)
	}

	return
}
