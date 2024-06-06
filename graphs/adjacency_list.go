package graphs

import (
	"bufio"
	"fmt"
	"main/utils"
	"os"
	"strconv"
	"strings"
)

type AdjacencyList struct {
	Edges map[int][]int
}

func (m *AdjacencyList) BuildFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		edgeParts := strings.Fields(line)

		v1, _ := strconv.Atoi(edgeParts[0])
		v2, _ := strconv.Atoi(edgeParts[1])
		m.AddEdge(v1, v2)
	}

	return nil
}

func NewAdjacencyList() AdjacencyList {
	return AdjacencyList{
		Edges: make(map[int][]int),
	}
}

func (graph *AdjacencyList) AddEdge(v1, v2 int) {
	graph.Edges[v1] = append(graph.Edges[v1], v2)
}

func (graph AdjacencyList) isSafe(v int, path []int, pos int) bool {
	lastVertex := path[pos-1]
	isAdjacent := false
	for _, vertex := range graph.Edges[lastVertex] {
		if vertex == v {
			isAdjacent = true
			break
		}
	}

	if !isAdjacent {
		return false
	}

	for i := 0; i < pos; i++ {
		if path[i] == v {
			return false
		}
	}

	return true
}

func (graph AdjacencyList) HamCycleDFS(path []int, pos int, vertices int) bool {
	if pos == vertices {
		if utils.Contains(graph.Edges[path[pos-1]], path[0]) {
			return true
		} else {
			return false
		}
	}

	for _, v := range graph.Edges[path[pos-1]] {
		if graph.isSafe(v, path, pos) {
			path[pos] = v

			if graph.HamCycleDFS(path, pos+1, vertices) == true {
				return true
			}

			path[pos] = -1
		}
	}

	return false
}

func (graph AdjacencyList) FindHamiltonianCycle() bool {
	vertices := len(graph.Edges)
	path := make([]int, vertices)
	for i := range path {
		path[i] = -1
	}

	path[0] = 0
	if graph.HamCycleDFS(path, 1, vertices) == false {
		fmt.Println("Cycle does not exist")
		return false
	}

	fmt.Println("Hamiltonian cycle:")
	utils.PrintPath(path)

	return true
}

func (g *AdjacencyList) CheckEulerianCycle() bool {
	inDegree := make(map[int]int)
	outDegree := make(map[int]int)

	for v, edges := range g.Edges {
		outDegree[v] = len(edges)
		for _, u := range edges {
			inDegree[u]++
		}
	}

	for v := range g.Edges {
		if inDegree[v] != outDegree[v] {
			return false
		}
	}

	return true
}

func (g *AdjacencyList) FindEulerianCycle() []int {
	if !g.CheckEulerianCycle() {
		return nil
	}

	cycle := []int{}
	currentPath := []int{}

	var startVertex int
	for v := range len(g.Edges) {
		if len(g.Edges[v]) > 0 {
			startVertex = v
			break
		}
	}

	currentPath = append(currentPath, startVertex)
	for len(currentPath) > 0 {
		currentVertex := currentPath[len(currentPath)-1]
		if len(g.Edges[currentVertex]) > 0 {
			nextVertex := g.Edges[currentVertex][0]
			g.Edges[currentVertex] = g.Edges[currentVertex][1:]
			currentPath = append(currentPath, nextVertex)
		} else {
			cycle = append([]int{currentVertex}, cycle...)
			currentPath = currentPath[:len(currentPath)-1]
		}
	}

	fmt.Println("Eulerian cycle:")
	utils.PrintPath(cycle[:len(cycle)-1])
	return cycle
}
