package util

import (
	"fmt"
	"strconv"
	"strings"
)

func Atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

type Position struct {
	X int
	Y int
}

func (p1 Position) Add(p2 Position) Position {
	return Position{p1.X + p2.X, p1.Y + p2.Y}
}

func (p1 Position) Sub(p2 Position) Position {
	return Position{p1.X - p2.X, p1.Y - p2.Y}
}

func StrReverse(s string) string {
	var ret string
	for _, v := range s {
		ret = string(v) + ret
	}
	return ret
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}

func LCM(numbers ...int) int {
	if len(numbers) == 0 {
		return -1
	}
	if len(numbers) == 1 {
		return numbers[0]
	}
	return lcm(numbers[0], numbers[1], numbers[2:]...)
}

func GCD(numbers ...int) int {
	if len(numbers) == 0 {
		return -1
	}
	if len(numbers) == 1 {
		return numbers[0]
	}
	D := gcd(numbers[0], numbers[1])
	newNumbers := []int{D}
	newNumbers = append(newNumbers, numbers[2:]...)
	return GCD(newNumbers...)
}

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func ABS[T Number](n T) T {
	if n < 0 {
		return -n
	}
	return n
}

type Direction int

const (
	DirU = iota
	DirR
	DirD
	DirL
)

var Offsets = [4]Position{
	{X: 0, Y: -1},
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: -1, Y: 0},
}

func (d Direction) Rotate(n int) Direction {
	return Direction((int(d) + (4 + (n % 4))) % 4)
}

func (d Direction) Clockwise() Direction {
	return d.Rotate(1)
}

func (d Direction) AntiClockwise() Direction {
	return d.Rotate(-1)
}

var DirToOffset = map[Direction]Position{
	DirU: {X: 0, Y: -1},
	DirR: {X: 1, Y: 0},
	DirD: {X: 0, Y: 1},
	DirL: {X: -1, Y: 0},
}

type Grid struct {
	Grid   [][]byte
	Height int
	Width  int
}

func ParseGrid(str string) Grid {
	lines := strings.Split(str, "\n")
	g := Grid{}
	for _, l := range lines {
		g.Grid = append(g.Grid, []byte(l))
	}
	g.Height = len(lines)
	g.Width = len(lines[0])
	return g
}

func (g Grid) Print() {
	fmt.Println("Printing Grid...")
	for _, v := range g.Grid {
		for _, b := range v {
			fmt.Printf(string(b))
		}
		fmt.Println()
	}
}
