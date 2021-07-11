package main

import (
	"math/rand"
	"time"

	"github.com/bl17zar/cell/cell"
	"github.com/bl17zar/cell/machine"
)

const (
	cellSize = 31
	xMult    = 2
)

func Alive() func(g *cell.Graph, m *cell.Map) {
	return func(g *cell.Graph, m *cell.Map) {
		g.AddNode(8, 16)

		for n := range g.Nodes {
			d, err := cell.GetNodeDisplay(g.Nodes[n])
			if err != nil {
				panic(err)
			}

			m.Set(d.Row, d.Col, d.Sign)
		}

		for e := range g.Edges {
			for _, eN := range g.Edges[e] {
				d, err := cell.GetEdgeDisplay(g.Nodes[e], eN)
				if err != nil {
					panic(err)
				}
				m.Set(d.Row, d.Col, d.Sign)
			}
		}
	}
}

func Dead() func(g *cell.Graph, m *cell.Map) {
	return func(g *cell.Graph, m *cell.Map) {
		g.AddNode(16, 16)

		for n := range g.Nodes {
			d, err := cell.GetNodeDisplay(g.Nodes[n])
			if err != nil {
				panic(err)
			}

			m.Set(d.Row, d.Col, d.Sign)
		}

		for e := range g.Edges {
			for _, eN := range g.Edges[e] {
				d, err := cell.GetEdgeDisplay(g.Nodes[e], eN)
				if err != nil {
					panic(err)
				}
				m.Set(d.Row, d.Col, d.Sign)
			}
		}
	}
}

func Random() func(g *cell.Graph, m *cell.Map) {
	return func(g *cell.Graph, m *cell.Map) {
		rand.Seed(time.Now().UnixNano())
		g.AddNode(rand.Intn(cellSize), rand.Intn(cellSize))

		for n := range g.Nodes {
			d, err := cell.GetNodeDisplay(g.Nodes[n])
			if err != nil {
				panic(err)
			}

			m.Set(d.Row, d.Col, d.Sign)
		}

		for e := range g.Edges {
			for _, eN := range g.Edges[e] {
				d, err := cell.GetEdgeDisplay(g.Nodes[e], eN)
				if err != nil {
					panic(err)
				}
				m.Set(d.Row, d.Col, d.Sign)
			}
		}
	}
}

func main() {
	m := machine.NewMachine(cellSize, xMult, Random())

	m.Run()
}
