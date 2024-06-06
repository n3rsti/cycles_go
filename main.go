package main

import (
	"fmt"
	"main/graphs"
)

func main() {
	adjList := graphs.NewAdjacencyList()
	adjList.BuildFromFile("input")

	adjMatrix := graphs.NewAdjacencyMatrix(0)
	adjMatrix.BuildFromFile("input")

	fmt.Println("Adjacency List:")
	adjList.FindHamiltonianCycle()
	adjList.FindEulerianCycle()

	fmt.Println("\nAdjacency Matrix:")
	adjMatrix.FindHamiltonianCycle()
	adjMatrix.FindEulerianCycle()
}
