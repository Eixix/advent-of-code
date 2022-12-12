package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	id        int
	ways      map[int]int
	isStart   bool
	isEnd     bool
	cost      int
	precessor int
	deleted   bool
}

func main() {
	first()
}

func abs(x rune) int {
	number := int(x)
	if x < 0 {
		return -number
	} else {
		return number
	}
}

func parseText(filename string) ([]rune, int) {
	elevationMap := []rune{}
	file, _ := os.Open(filename)
	rowLength := 0
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rowLength = len(scanner.Text())
		for _, letter := range scanner.Text() {
			elevationMap = append(elevationMap, letter)
		}
	}
	return elevationMap, rowLength
}

func getStartAndEndNode(elevationMap []rune) (int, int) {
	start := 0
	end := 0

	for i, v := range elevationMap {
		if v == 'S' {
			start = i
		} else if v == 'E' {
			end = i
		}
	}

	return start, end
}

func walkAndMap(elevationMap []rune, start int, end int, rowLength int) []Node {
	toReturn := []Node{}

	for i, node := range elevationMap {
		if i == start {
			node = 'a'
		} else if i == end {
			node = 'z'
		}

		ways := make(map[int]int)

		if ((i + 1) % rowLength) != 0 {
			distance := abs(node - elevationMap[i+1])
			ways[i+1] = distance
		}
		if i != 0 && (i%(rowLength+1)) != 0 {
			distance := abs(node - elevationMap[i-1])
			ways[i-1] = distance
		}
		if i+8 <= len(elevationMap)-1 {
			distance := abs(node - elevationMap[i+8])
			ways[i+8] = distance
		}
		if i-8 >= 0 {
			distance := abs(node - elevationMap[i-8])
			ways[i-8] = distance
		}

		newNode := Node{i, ways, i == start, i == end, 1000000, 0, false}
		toReturn = append(toReturn, newNode)
	}

	return toReturn
}

func getMinDistanceNode(Q []Node) Node {
	distance := 1000000
	var toReturn Node

	for _, node := range Q {
		if node.cost < distance && !node.deleted {
			distance = node.cost
			toReturn = node
		}
	}
	return toReturn
}

func dijkstra(graph []Node, start int) []Node {
	Q := []Node{}

	graph[start].cost = 0
	for i := range graph {
		Q = append(Q, graph[i])
	}

	deleted := 0

	for len(Q) != deleted {
		minDistanceNode := getMinDistanceNode(Q)
		Q[minDistanceNode.id].deleted = true
		deleted++

		for neighbour, way := range minDistanceNode.ways {
			if Q[neighbour].id != -1 {
				a := minDistanceNode.cost + way
				if a < graph[neighbour].cost {
					graph[neighbour].cost = a
					graph[neighbour].precessor = neighbour
				}
			}
		}
	}

	return graph
}

func first() {
	elevationMap, rowLength := parseText("small.txt")
	start, end := getStartAndEndNode(elevationMap)
	graph := walkAndMap(elevationMap, start, end, rowLength)
	fmt.Println(dijkstra(graph, start))
	//nodes := make(map[int]Node)

}
