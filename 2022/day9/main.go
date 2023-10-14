package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type Knot struct {
	x int
	y int
}

type Move struct {
	direction string
	times     int
}

type Rope struct {
	length int
	knots  []Knot
}

func NewRope(l int) Rope {
	return Rope{l, make([]Knot, l)}
}

//go:embed example.txt
var example string

//go:embed exampleXL.txt
var exampleXL string

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
	example = strings.TrimRight(example, "\n")
	if len(example) == 0 {
		panic("empty example.txt file")
	}
	exampleXL = strings.TrimRight(exampleXL, "\n")
	if len(example) == 0 {
		panic("empty exampleXL.txt file")
	}
}

func main() {
	var part string
	flag.StringVar(&part, "part", "", "part 1 or 2")
	flag.Parse()

	if part == "1" || part == "" {
		fmt.Println("--- Running part 1 ---")
		ansExample := foo(example, 2)
		fmt.Println("Output Example:", ansExample)
		ansInput := foo(input, 2)
		fmt.Println("Output Input:", ansInput)
	}
	if part == "2" || part == "" {
		fmt.Println("--- Running part 2 ---")
		ansExample := foo(example, 10)
		fmt.Println("Output Example:", ansExample)
		ansExampleXL := foo(exampleXL, 10)
		fmt.Println("Output ExampleXL:", ansExampleXL)
		ansInput := foo(input, 10)
		fmt.Println("Output Input:", ansInput)
	}
}

func foo(input string, N int) int {
	moves := parseInput(input)
	rope := NewRope(N)
	visitedByTail := map[[2]int]bool{}
	for _, m := range moves {
		for j := 0; j < m.times; j++ {
			switch m.direction {
			case "U":
				rope.knots[0].y++
				break
			case "R":
				rope.knots[0].x++
				break
			case "D":
				rope.knots[0].y--
				break
			case "L":
				rope.knots[0].x--
				break
			}
			for k := 1; k < rope.length; k++ {
				dx := rope.knots[k-1].x - rope.knots[k].x
				dy := rope.knots[k-1].y - rope.knots[k].y
				diff := Knot{dx, dy}
				switch diff {
				case Knot{2, 0}:
					rope.knots[k].x++
				case Knot{-2, 0}:
					rope.knots[k].x--
				case Knot{0, 2}:
					rope.knots[k].y++
				case Knot{0, -2}:
					rope.knots[k].y--
				case Knot{2, 1}, Knot{1, 2}, Knot{2, 2}:
					rope.knots[k].x++
					rope.knots[k].y++
				case Knot{2, -1}, Knot{1, -2}, Knot{2, -2}:
					rope.knots[k].x++
					rope.knots[k].y--
				case Knot{-2, -1}, Knot{-1, -2}, Knot{-2, -2}:
					rope.knots[k].x--
					rope.knots[k].y--
				case Knot{-2, 1}, Knot{-1, 2}, Knot{-2, 2}:
					rope.knots[k].x--
					rope.knots[k].y++
				}
			}
			visitedByTail[[2]int{rope.knots[N-1].x, rope.knots[N-1].y}] = true
		}
	}
	return len(visitedByTail)
}

func parseInput(input string) []Move {
	moves := []Move{}
	for _, line := range strings.Split(input, "\n") {
		tmp := strings.Split(line, " ")
		d := tmp[0]
		t, _ := strconv.Atoi(tmp[1])
		moves = append(moves, Move{d, t})
	}
	return moves
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	} else {
		return n
	}
}
