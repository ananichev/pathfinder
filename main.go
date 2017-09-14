package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

type Point struct {
	X, Y, F, G, H int
	Parent        *Point
}

type Points []Point

type Dimension struct {
	maxX, maxY int
}

var (
	startPoint, destPoint           Point
	openList, closedList, obstacles Points
	dimension                       Dimension
)

var diagonalDirections = [][]int{{-1, -1}, {1, -1}, {-1, 1}, {1, 1}}
var ortogonalDirections = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func main() {
	readFile("./field")
	if point, ok := search(); ok {
		printPath(point)
	}
}

func search() (current *Point, ok bool) {
	for len(openList) > 0 {
		current, openList = &openList[0], openList[1:]
		closedList = append(closedList, *current)
		if current.X == destPoint.X && current.Y == destPoint.Y {
			ok = true
			return
		}
		findSiblings(current)
	}
	fmt.Printf("Can not found path to Point{X: %d, Y: %d}", destPoint.X, destPoint.Y)
	return
}

func findSiblings(p *Point) {
	siblings(p, diagonalDirections, 14)
	siblings(p, ortogonalDirections, 10)
	sort.Sort(openList)
}

func siblings(p *Point, directions [][]int, cost int) {
	for _, d := range directions {
		cur, ok := openList.find(p.X+d[0], p.Y+d[1])
		if !ok {
			cur = Point{X: p.X + d[0], Y: p.Y + d[1], Parent: p, G: cost}
		}
		if withinTheBoard(cur.X, cur.Y) && !(cur.in(obstacles) || cur.in(closedList)) {
			manhattenDistance(&cur)

			if cur.in(openList) && (pathCost(&cur)+cost) < pathCost(p) {
				p.Parent = &cur
			}
			openList = append(openList, cur)
		}
	}
}

func pathCost(p *Point) (cost int) {
	for p != nil {
		cost = cost + p.G
		p = p.Parent
	}
	return
}

func manhattanDistance(current *Point) {
	current.H = int(math.Abs(float64(destPoint.X-current.X))+math.Abs(float64(destPoint.Y-current.Y))) * 10
	current.F = current.G + current.H
}

func (p *Point) in(list Points) bool {
	for _, el := range list {
		if p.X == el.X && p.Y == el.Y {
			return true
		}
	}
	return false
}

func (l *Points) find(x, y int) (Point, bool) {
	for _, p := range *l {
		if p.X == x && p.Y == y {
			return p, true
		}
	}
	return Point{}, false
}

func (l Points) Less(i, j int) bool {
	return l[i].F < l[j].F
}

func (l Points) Len() int {
	return len(l)
}

func (l Points) Swap(i, j int) {
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
				obstacles = append(obstacles, Point{X: x, Y: y})
			case "1":
				startPoint = Point{X: x, Y: y}
				openList = append(openList, startPoint)
			case "2":
				destPoint = Point{X: x, Y: y}
			}
		}
		y++
	}
	dimension = Dimension{maxX: x + 1, maxY: y}
}

func printPath(p *Point) {
	var rightWay Points
	for p != nil {
		rightWay = append(rightWay, *p)
		p = p.Parent
	}

	var c Point
	for y := 0; y < dimension.maxY; y++ {
		for x := 0; x < dimension.maxX; x++ {
			c = Point{X: x, Y: y}
			switch {
			case c.in(obstacles):
				fmt.Print(" X ")
			case c.X == startPoint.X && c.Y == startPoint.Y:
				fmt.Print(" 1 ")
			case c.X == destPoint.X && c.Y == destPoint.Y:
				fmt.Print(" 2 ")
			case c.in(rightWay):
				fmt.Print(" + ")
			default:
				fmt.Print(" 0 ")
			}
		}
		fmt.Print("\n")
	}
}
