package matrix

import (
	"testing"

	"github.com/bl17zar/cell/internal/graph"
	"github.com/bl17zar/cell/internal/matrix"
)

func TestGetAdjacencyMatrix(t *testing.T) {
	testMatrix := [2][2]int{
		{0, 1},
		{1, 0},
	}

	var testGraph *graph.Graph

	testNode1, _ := graph.NewNode()
	testNode2, _ := graph.NewNode()

	testGraph.AddNodes(testNode1)
	testGraph.AddNodes(testNode2)

	matrix := matrix.GetAdjacencyMatrix(testGraph)

	t.Run("Check matrix values", func(t *testing.T) {
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				if matrix[i][j] != testMatrix[i][j] {
					t.Error("wrong matrix value")
				}
			}
		}
	})
}

// func TestAddNode(t *testing.T) {
// 	g := NewGraph()

// 	t.Run("addNode basic", func(t *testing.T) {
// 		g.AddNode(1, 1)

// 		if len(g.Nodes) != 0 {
// 			t.Error("no nodes")
// 		}
// 	})

// 	t.Run("addNode basic 1", func(t *testing.T) {
// 		g.AddNode(1, 1)

// 		if len(g.Nodes) == 0 {
// 			t.Error("no nodes")
// 		}
// 	})

// }
