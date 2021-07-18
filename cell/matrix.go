package cell

import "fmt"

func GetAdjacencyMatrix(g *Graph, matrixSize int) [][]int {
	matrix := make([][]int, matrixSize)

	for i := 0; i < matrixSize; i++ {
		for j := 0; j < matrixSize; j++ {
			matrix[i][j] = 0
		}
	}

	for _, Edge := range g.Edges {
		Col := Edge['0'].Col
		Row := Edge['0'].Row

		matrix[Col][Row] = 1
		matrix[Row][Col] = 1
	}

	return matrix
}

func TestGetAdjacencyMatrix() int {
	testMatrix := [3][3]int{
		{0, 1, 0},
		{1, 0, 0},
		{0, 0, 0},
	}
	testGraph := NewGraph()

	for i := 0; i < 3; i++ {
		for j := 0; i < 3; i++ {
			testGraph.AddNode(i, j)
		}
	}

	testGraph.AddEdge(testGraph.Nodes["0"], testGraph.Nodes["1"])
	adjacencyMatrix := GetAdjacencyMatrix(testGraph, 3)

	for i := 0; i < 3; i++ {
		for j := 0; i < 3; i++ {
			if testMatrix[i][j] != adjacencyMatrix[i][j] {
				fmt.Print("FAIL\n")
				return -1
			}
		}
	}
	fmt.Print("PASS\n")
	return 0
}
