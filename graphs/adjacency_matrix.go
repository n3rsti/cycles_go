package graphs

import (
	"bufio"
	"fmt"
	"main/utils"
	"os"
	"strconv"
	"strings"
)

type AdjacencyMatrix struct {
	Matrix [][]int
}

func (m *AdjacencyMatrix) BuildFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	firstLine := scanner.Text()
	parts := strings.Fields(firstLine)

	numVertices, _ := strconv.Atoi(parts[0])
	graph := NewAdjacencyMatrix(numVertices)

	for scanner.Scan() {
		line := scanner.Text()
		edgeParts := strings.Fields(line)

		v1, _ := strconv.Atoi(edgeParts[0])
		v2, _ := strconv.Atoi(edgeParts[1])
		graph.AddEdge(v1, v2)
	}

	m.Matrix = graph.Matrix

	return nil
}

func NewAdjacencyMatrix(vertices int) *AdjacencyMatrix {
	matrix := make([][]int, vertices)
	for i := range matrix {
		matrix[i] = make([]int, vertices)
	}
	return &AdjacencyMatrix{matrix}
}

func (graph AdjacencyMatrix) isSafe(v int, path []int, pos int) bool {
	if graph.Matrix[path[pos-1]][v] == 0 {
		return false
	}

	for i := 0; i < pos; i++ {
		if path[i] == v {
			return false
		}
	}

	return true
}

func (graph AdjacencyMatrix) HamCycleUtil(path []int, pos int, vertices int) bool {
	if pos == vertices {
		if graph.Matrix[path[pos-1]][path[0]] == 1 {
			return true
		} else {
			return false
		}
	}

	for v := 1; v < vertices; v++ {
		if graph.isSafe(v, path, pos) {
			path[pos] = v

			if graph.HamCycleUtil(path, pos+1, vertices) == true {
				return true
			}

			path[pos] = -1
		}
	}

	return false
}

func (graph AdjacencyMatrix) FindHamiltonianCycle() bool {
	vertices := len(graph.Matrix)
	path := make([]int, vertices)
	for i := range path {
		path[i] = -1
	}

	path[0] = 0
	if graph.HamCycleUtil(path, 1, vertices) == false {
		fmt.Println("Cycle does not exist")
		return false
	}

	fmt.Println("Hamiltonian cycle:")
	utils.PrintPath(path)

	return true
}

func (graph *AdjacencyMatrix) AddEdge(v1, v2 int) {
	graph.Matrix[v1][v2] = 1
	graph.Matrix[v2][v1] = 1
}

func (graph *AdjacencyMatrix) CheckEulerianCycle() bool {
	size := len(graph.Matrix)
	inDegree := make([]int, size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			inDegree[j] += graph.Matrix[i][j]
		}
	}

	for i := 0; i < size; i++ {
		if inDegree[i]%2 != 0 {
			return false
		}
	}
	return true
}

func (graph *AdjacencyMatrix) FindEulerianCycle() []int {
	if !graph.CheckEulerianCycle() {
		return nil
	}

	size := len(graph.Matrix)
	visited := make([]bool, size)

	cycle := []int{}
	stack := []int{0}

	for len(stack) > 0 {
		u := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for v := 0; v < size; v++ {
			if graph.Matrix[u][v] != 0 && !visited[u] {
				stack = append(stack, v)
				cycle = append(cycle, u)
				visited[u] = true
				u = v
			}
		}
	}

	fmt.Println("Eulerian cycle:")
	utils.PrintPath(cycle)

	return cycle
}
