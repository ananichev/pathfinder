package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

type Node struct {
	X, Y, F, G, H int
	Parent        *Node
}

type Nodes []Node

type Dimension struct {
	maxX, maxY int
}

var (
	startNode, destNode             Node
	openList, closedList, obstacles Nodes
	dimension                       Dimension
)

var diagonalDirections = [][]int{{-1, -1}, {1, -1}, {-1, 1}, {1, 1}}
var ortogonalDirections = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func main() {
	readFile("./field")
	checkRequires()

	if node, ok := search(); ok {
		printPath(node)
	} else {
		fmt.Printf("Can not find path to Node{X: %d, Y: %d}", destNode.X, destNode.Y)
	}
}

func checkRequires() {
	if (startNode == Node{} || destNode == Node{}) {
		fmt.Println("Missing start/dest node!")
		os.Exit(1)
	}
}

func search() (current *Node, ok bool) {
	for len(openList) > 0 {
		current, openList = &openList[0], openList[1:]
		closedList = append(closedList, *current)
		if current.equal(destNode) {
			ok = true
			return
		}
		findSiblings(current)
	}
	return
}

func findSiblings(n *Node) {
	siblings(n, diagonalDirections, 14)
	siblings(n, ortogonalDirections, 10)
	sort.Sort(openList)
}

func siblings(n *Node, directions [][]int, cost int) {
	for _, d := range directions {
		if unreachable(n, d) {
			continue
		}

		cur, ok := openList.find(n.X+d[0], n.Y+d[1])
		if !ok {
			cur = Node{X: n.X + d[0], Y: n.Y + d[1], Parent: n, G: cost}
		}

		manhattanDistance(&cur)
		if cur.in(openList) && (pathCost(&cur)+cost) <= pathCost(n) {
			n.Parent = &cur
		}
		openList = append(openList, cur)
	}
}

func unreachable(n *Node, d []int) bool {
	newX, newY := n.X+d[0], n.Y+d[1]
	_, xObstacle := obstacles.find(newX, n.Y)
	_, yObstacle := obstacles.find(n.X, newY)
	_, obstacle := obstacles.find(newX, newY)
	_, closed := closedList.find(newX, newY)
	return !withinTheBoard(newX, newY) || obstacle || closed || (xObstacle && yObstacle)
}

func pathCost(n *Node) (cost int) {
	for n != nil {
		cost = cost + n.G
		n = n.Parent
	}
	return
}

func manhattanDistance(current *Node) {
	current.H = int(math.Abs(float64(destNode.X-current.X))+math.Abs(float64(destNode.Y-current.Y))) * 10
	current.F = current.G + current.H
}

func (n *Node) in(list Nodes) bool {
	for _, el := range list {
		if n.equal(el) {
			return true
		}
	}
	return false
}

func (l Nodes) find(x, y int) (Node, bool) {
	for _, n := range l {
		if n.X == x && n.Y == y {
			return n, true
		}
	}
	return Node{}, false
}

func (n Node) equal(o Node) bool {
	return n.X == o.X && n.Y == o.Y
}

func (l Nodes) Less(i, j int) bool {
	return l[i].F < l[j].F
}

func (l Nodes) Len() int {
	return len(l)
}

func (l Nodes) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func withinTheBoard(x, y int) bool {
	return x >= 0 && y >= 0 && x < dimension.maxX && y < dimension.maxY
}

func readFile(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var x, y int
	var str string
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		for x, str = range line {
			switch str {
			case "X":
				obstacles = append(obstacles, Node{X: x, Y: y})
			case "1":
				startNode = Node{X: x, Y: y}
				openList = append(openList, startNode)
			case "2":
				destNode = Node{X: x, Y: y}
			}
		}
		y++
	}
	dimension = Dimension{maxX: x + 1, maxY: y}
}

func printPath(n *Node) {
	var rightWay Nodes
	for n != nil {
		rightWay = append(rightWay, *n)
		n = n.Parent
	}

	var c Node
	for y := 0; y < dimension.maxY; y++ {
		for x := 0; x < dimension.maxX; x++ {
			c = Node{X: x, Y: y}
			switch {
			case c.in(obstacles):
				fmt.Print("X", " ")
			case c.equal(startNode):
				fmt.Print("1", " ")
			case c.equal(destNode):
				fmt.Print("2", " ")
			case c.in(rightWay):
				fmt.Print("+", " ")
			default:
				fmt.Print(".", " ")
			}
		}
		fmt.Print("\n")
	}
}
