package main

import (
	"math/rand"
	"time"

	"github.com/bl17zar/cell/cell"
	"github.com/bl17zar/cell/machine"
)

const (
	cellSize = 41
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

func Random(count int) func(g *cell.Graph, m *cell.Map) {
	return func(g *cell.Graph, m *cell.Map) {
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < count; i++ {
			g.AddNode(rand.Intn(cellSize-1), rand.Intn(cellSize-1))
		}

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

func drawDown(sign cell.SignType, startRow, col, distance int) []*cell.Display {
	res := make([]*cell.Display, 0, distance)
	for i := 1; i <= distance; i++ {
		res = append(res, &cell.Display{
			Row:  startRow + i,
			Col:  col,
			Sign: sign,
		})
	}

	return res
}

func drawRight(sign cell.SignType, row, startCol, distance int) []*cell.Display {
	res := make([]*cell.Display, 0, distance)
	for i := 1; i <= distance; i++ {
		res = append(res, &cell.Display{
			Row:  row,
			Col:  startCol + i,
			Sign: sign,
		})
	}

	return res
}

func main() {
	features := []*cell.Display{}
	features = append(features, drawDown(cell.SignObstacle, 5, 13, 21)...)
	features = append(features, drawDown(cell.SignObstacle, 5, 29, 21)...)
	features = append(features, drawRight(cell.SignObstacle, 5, 12, 17)...)
	features = append(features, drawRight(cell.SignObstacle, 27, 12, 6)...)
	features = append(features, drawRight(cell.SignObstacle, 27, 23, 6)...)

	m := machine.NewMachine(cellSize, xMult, Alive(), features)

	m.Run()
}
