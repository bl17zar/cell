package matrix

import (
	"fmt"

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

	nodesQueue := make([]*graph.Node, len(g.Nodes))

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

/* ========== TEST ========== */

func TestGetAdjacencyMatrix() int {
	testMatrix := [2][2]int{
		{0, 1},
		{1, 0},
	}

	var testGraph *graph.Graph

	testNode1, _ := graph.NewNode()
	testNode2, _ := graph.NewNode()

	testGraph.AddNodes(testNode1)
	testGraph.AddNodes(testNode2)

	matrix := GetAdjacencyMatrix(testGraph)

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if matrix[i][j] != testMatrix[i][j] {
				fmt.Print("FAIL\n")
			}
		}
	}
	fmt.Print("PASS\n")
	return 0
}
