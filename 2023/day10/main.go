package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type QueueNode struct {
	Value Node
	Next  *QueueNode
}

type Queue struct {
	front, rear *QueueNode
	length      int
}

type Coord struct {
	X int
	Y int
}

func (q *Queue) Enque(item Node) {
	newNode := &QueueNode{Next: nil, Value: item}
	if q.rear == nil {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.Next = newNode
		q.rear = newNode
	}
	q.length++
}

func (q *Queue) Dequeue() Node {
	if q.front == nil {
		panic("Queue is empty")
	}
	item := q.front.Value
	q.front = q.front.Next
	if q.front == nil {
		q.rear = nil
	}
	q.length--
	return item
}

type Node struct {
	rowIndex          int
	columnIndex       int
	North             bool
	South             bool
	East              bool
	West              bool
	DistanceFromStart int
}

func (n *Node) goNorth(node *Node) bool {
	return n.North && node.South
}

func (n *Node) goSouth(node *Node) bool {
	return n.South && node.North
}

func (n *Node) goEast(node *Node) bool {
	return n.East && node.West
}

func (n *Node) goWest(node *Node) bool {
	return n.West && node.East
}

func createPipesRow(data []rune, rowCount int) []Node {
	var res []Node
	for i, pipe := range data {
		north, south, east, west := getPosibleDirections(pipe)
		res = append(res, Node{North: north, South: south, East: east, West: west, rowIndex: rowCount, columnIndex: i})
	}
	return res
}

func findStartIndex() (row, column int) {
	var rowIndex, columnIndex int
	file, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(file), "\n")

	for i, line := range lines {
		column := strings.IndexRune(line, 'S')

		if column != -1 {
			rowIndex = i
			columnIndex = column

			break
		}

	}
	return rowIndex, columnIndex
}

func getPosibleDirections(pipe rune) (north, south, east, west bool) {
	switch pipe {
	case '|':
		north = true
		south = true
	case '-':
		west = true
		east = true
	case 'L':
		north = true
		east = true
	case 'J':
		north = true
		west = true
	case '7':
		south = true
		west = true
	case 'F':
		south = true
		east = true
	case 'S':
		north = true
		south = true
		east = true
		west = true

	}
	return north, south, east, west
}

func shouldVisitNode(nodes [][]Node, visited [][]bool, currentNode Node, direction rune) bool {
	switch direction {
	case 'N':
		return currentNode.rowIndex > 0 && !visited[currentNode.rowIndex-1][currentNode.columnIndex]
	case 'E':
		return currentNode.columnIndex < len(nodes[currentNode.rowIndex])-1 && !visited[currentNode.rowIndex][currentNode.columnIndex+1]
	case 'S':
		return currentNode.rowIndex < len(nodes)-1 && !visited[currentNode.rowIndex+1][currentNode.columnIndex]
	case 'W':
		return currentNode.columnIndex > 0 && !visited[currentNode.rowIndex][currentNode.columnIndex-1]
	default:
		return false
	}
}

func walk(nodes [][]Node, startRow, startColumn int) (int, [][]bool, []Coord) {
	queue := Queue{}

	var longestPath int

	var coords []Coord

	visited := make([][]bool, len(nodes))
	for i := range visited {
		visited[i] = make([]bool, len(nodes[i]))
	}

	queue.Enque(nodes[startRow][startColumn])

	for {
		if queue.length == 0 {
			break
		}

		currentNode := queue.Dequeue()

		if currentNode.DistanceFromStart > longestPath {
			longestPath = currentNode.DistanceFromStart
		}

		visited[currentNode.rowIndex][currentNode.columnIndex] = true

		coords = append(coords, Coord{X: currentNode.columnIndex, Y: currentNode.rowIndex})

		if shouldVisitNode(nodes, visited, currentNode, 'E') {
			eastNode := nodes[currentNode.rowIndex][currentNode.columnIndex+1]

			if currentNode.goEast(&eastNode) {
				eastNode.DistanceFromStart = currentNode.DistanceFromStart + 1
				queue.Enque(eastNode)
			}

		}

		if shouldVisitNode(nodes, visited, currentNode, 'W') {
			westNode := nodes[currentNode.rowIndex][currentNode.columnIndex-1]

			if currentNode.goWest(&westNode) {
				westNode.DistanceFromStart = currentNode.DistanceFromStart + 1
				queue.Enque(westNode)
			}

		}

		if shouldVisitNode(nodes, visited, currentNode, 'N') {
			northIndex := nodes[currentNode.rowIndex-1][currentNode.columnIndex]

			if currentNode.goNorth(&northIndex) {
				northIndex.DistanceFromStart = currentNode.DistanceFromStart + 1
				queue.Enque(northIndex)
			}

		}

		if shouldVisitNode(nodes, visited, currentNode, 'S') {
			southIndex := nodes[currentNode.rowIndex+1][currentNode.columnIndex]

			if currentNode.goSouth(&southIndex) {
				southIndex.DistanceFromStart = currentNode.DistanceFromStart + 1
				queue.Enque(southIndex)
			}

		}

	}

	return longestPath, visited, coords

}

func shoelaceFormula(vertices []Coord) float64 {
	n := len(vertices)

	area := 0.0
	for i := 0; i < n; i++ {
		nextIndex := (i + 1) % len(vertices)
		area += float64((vertices[i].X * vertices[nextIndex].Y) - (vertices[i].Y * vertices[nextIndex].X))
	}

	return 0.5 * math.Abs(area)
}

func main() {
	var path [][]Node

	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var rowCount int

	startRow, startColumn := findStartIndex()

	for scanner.Scan() {
		pipesInRow := []rune(scanner.Text())

		path = append(path, createPipesRow(pipesInRow, rowCount))
		rowCount++
	}

	res, _, coords := walk(path, startRow, startColumn)

	fmt.Printf("Longest path is: %d\n", res)

	fmt.Println(len(coords) / 2)

	//coords = []Coord{{X: 2, Y: 1}, {X: 5, Y: 0}, {X: 6, Y: 4}, {X: 4, Y: 2}, {X: 1, Y: 3}}

	fmt.Println(coords)

	fmt.Printf("Area is :%f", shoelaceFormula(coords))

	//fmt.Printf("Area is :%f", shoelaceFormula([]Coord{{X: 2, Y: 1}, {X: 5, Y: 0}, {X: 6, Y: 4}, {X: 4, Y: 2}, {X: 1, Y: 3}}))

	var countVisited int

	var tilesInside int

	fmt.Printf("Visited: %d\n", countVisited)

	fmt.Printf("Tiles inside: %d\n", tilesInside)

}
