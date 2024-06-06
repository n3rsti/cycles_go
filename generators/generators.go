package generators

import (
	"fmt"
	"math/rand"
)

func hasEdge(edges [][]int, v1, v2 int) bool {
	for _, edge := range edges {
		if (edge[0] == v1 && edge[1] == v2) || (edge[0] == v2 && edge[1] == v1) {
			return true
		}
	}
	return false
}

func edgeExists(edges [][]int, u, v int) bool {
	for _, edge := range edges {
		if (edge[0] == u && edge[1] == v) || (edge[0] == v && edge[1] == u) {
			return true
		}
	}
	return false
}

func UndirectedEulerian(numVertices, density int) [][]int {
	if numVertices < 3 || density < 0 || density > 100 {
		return nil
	}

	edges := [][]int{}
	degrees := make([]int, numVertices)
	totalEdges := (numVertices * (numVertices - 1)) / 2
	targetEdges := totalEdges * density / 100

	for i := 0; i < numVertices; i++ {
		edges = append(edges, []int{i, (i + 1) % numVertices})
		degrees[i]++
		degrees[(i+1)%numVertices]++
	}

	for len(edges) < targetEdges {
		i := rand.Intn(numVertices)
		j := rand.Intn(numVertices)
		if i != j && !edgeExists(edges, i, j) {
			edges = append(edges, []int{i, j})
			degrees[i]++
			degrees[j]++
		}
	}

	for i := 0; i < numVertices; i++ {
		if degrees[i]%2 != 0 {
			for {
				j := rand.Intn(numVertices)
				if j != i && degrees[j]%2 != 0 && !edgeExists(edges, i, j) {
					edges = append(edges, []int{i, j})
					degrees[i]++
					degrees[j]++
					break
				}
				fmt.Println(1)
			}
		}
	}

	return edges
}

func UndirectedHamiltonian(numVertices, density int) [][]int {
	edges := make([][]int, 0)

	for i := 0; i < numVertices-1; i++ {
		edges = append(edges, []int{i, i + 1})
	}
	edges = append(edges, []int{numVertices - 1, 0})

	totalPossibleEdges := numVertices * (numVertices - 1) / 2

	numEdgesToAdd := (totalPossibleEdges * density) / 100

	for i := 0; i < numEdgesToAdd; i++ {
		src := rand.Intn(numVertices)
		dst := rand.Intn(numVertices)
		edges = append(edges, []int{src, dst})
	}

	return edges
}

func DirectedEulerian(numVertices, density int) [][]int {
	numEdges := numVertices * (numVertices - 1) * density / 100
	edges := make([][]int, 0, numEdges)

	for i := 0; i < numVertices-1; i++ {
		edges = append(edges, []int{i, i + 1})
	}

	for len(edges) < numEdges {
		v1 := rand.Intn(numVertices)
		v2 := rand.Intn(numVertices)

		if v1 == v2 {
			continue
		}

		exists := false
		for _, edge := range edges {
			if edge[0] == v1 && edge[1] == v2 {
				exists = true
				break
			}
		}

		if !exists {
			edges = append(edges, []int{v1, v2})
		}
	}

	return edges
}

func DirectedHamiltonian(numVertices, density int) [][2]int {
	edges := make([][2]int, 0)

	for i := 0; i < numVertices-1; i++ {
		edges = append(edges, [2]int{i, i + 1})
	}
	edges = append(edges, [2]int{numVertices - 1, 0})

	totalPossibleEdges := numVertices * (numVertices - 1)

	numEdgesToAdd := (totalPossibleEdges * density) / 100

	for i := 0; i < numEdgesToAdd; i++ {
		src := rand.Intn(numVertices)
		dst := rand.Intn(numVertices)
		edges = append(edges, [2]int{src, dst})
	}

	return edges
}
