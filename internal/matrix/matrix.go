package matrix

import (
	"github.com/bl17zar/cell/internal/graph"
)

func getNeighbourIndex(node *graph.Node, nodesQueue []*graph.Node) int {
	for index, queueNode := range nodesQueue {
		if queueNode == node {
			return index
		}
	}
	return 0
}

func GetAdjacencyMatrix(g *graph.Graph) [][]int {
	var matrix [][]int

	// create matrix and fill with zeros
	for i := 0; i < len(g.Nodes); i++ {
		matrix = append(matrix, make([]int, len(g.Nodes)))
	}

	nodesQueue := make([]*graph.Node, 0, len(g.Nodes))

	// order node from graph
	for _, node := range g.Nodes {
		nodesQueue = append(nodesQueue, node)
	}

	for position, node := range nodesQueue {
		for _, neigbour := range node.Neighbours() {
			matrix[position][getNeighbourIndex(neigbour, nodesQueue)] = 1
		}
	}

	return matrix
}
