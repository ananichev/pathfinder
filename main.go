package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var p = fmt.Println

type Point struct {
  X, Y, F, G, H int
  Parent *Point
}

type Dimension struct {
  maxX, maxY int
}

var (
  startPoint, destPoint Point
  openList, closedList, obstacles []Point
  dimension Dimension
)

var ortogonalDirections = [][]int{
  {-1, -1},
  {1, -1},
  {-1, 1},
  {1, 1},
}

func main() {
	readFile("./field")
  // search
}

// func search() {
//   for len(openList) > 0 {
//
//   }
// }
//
// func findSiblings(p *Point) {
//     diagonalSiblings(p)
//     ortogonalSiblings(p)
// }

func diagonalSiblings(p *Point) {
  var newX, newY int
  for _, d := range ortogonalDirections {
    newX, newY = p.X + d[0], p.Y + d[1]
    if withinTheBoard(newX, newY) && !(obstacle(newX, newY) || closed(newX, newY) || open(newX, newY)) {
      openList = append(openList, Point{X: newX, Y: newY})
    }
  }
}

// func ortogonalSiblings(p *Point) {
//
// }

func obstacle(x, y int) bool {
  for _, p := range obstacles {
    if p.X == x && p.Y == y {
      return true
    }
  }
  return false
}

func closed(x, y int) bool {
  for _, p := range closedList {
    if p.X == x && p.Y == y {
      return true
    }
  }
  return false
}

func open(x, y int) bool {
  for _, p := range openList {
    if p.X == x && p.Y == y {
      return true
    }
  }
  return false
}

func withinTheBoard(x, y int) bool {
  return x >= 0 && y >= 0 && x <= dimension.maxX && y <= dimension.maxY
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
  dimension = Dimension{x, y}
  diagonalSiblings(&startPoint)
  p(openList)
  // p(closed(1, 1))
  // p(obstacle(1, 1))
}
