package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Element struct {
	value  int
	left   int
	top    int
	right  int
	bottom int
}

func NewElement(value int) *Element {
	e := Element{
		value:  value,
		left:   -1,
		top:    -1,
		right:  -1,
		bottom: -1,
	}
	return &e
}

type Grid [][]Element

func (grid Grid) print(args ...bool) {
	full := len(args) > 0
	for _, v := range grid {
		for _, e := range v {
			if full {
				fmt.Print(e)
			} else {
				fmt.Print(e.value)
			}
		}
		fmt.Println()
	}
}

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		fmt.Println("Output:", ans)
	} else if part == 2 {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	grid := parseInput(input)
	sum := solveP1(grid)
	return sum
}

func part2(input string) int {
	grid := parseInput(input)
	sum := solveP2(grid)
	return sum
}

func solveP1(grid Grid) int {
	dim := len(grid)
	// View from left
	for i := 0; i < dim; i++ {
		max := grid[i][0].value
		for j := 1; j < dim; j++ {
			grid[i][j].left = max
			if max < grid[i][j].value {
				max = grid[i][j].value
			}
		}
	}

	// View from right
	for i := 0; i < dim; i++ {
		max := grid[i][dim-1].value
		for j := dim - 2; j >= 0; j-- {
			grid[i][j].right = max
			if max < grid[i][j].value {
				max = grid[i][j].value
			}
		}
	}

	// View from top
	for i := 0; i < dim; i++ {
		max := grid[0][i].value
		for j := 1; j < dim; j++ {
			grid[j][i].top = max
			if max < grid[j][i].value {
				max = grid[j][i].value
			}
		}
	}

	// View from bottom
	for i := 0; i < dim; i++ {
		max := grid[dim-1][i].value
		for j := dim - 2; j >= 0; j-- {
			grid[j][i].bottom = max
			if max < grid[j][i].value {
				max = grid[j][i].value
			}
		}
	}

	// Check how many are visibile
	ans := 0
	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {
			if isVisible(grid[i][j]) {
				ans++
			}
		}
	}
	return ans
}

func solveP2(grid Grid) int {
	ans := 0
	n := len(grid)
	for i := 1; i < n-1; i++ {
		for j := 1; j < n-1; j++ {
			vd := getViewingDistance(grid, i, j)
			if vd > ans {
				ans = vd
			}
		}
	}
	return ans
}

func parseInput(input string) Grid {
	grid := Grid{}
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, []Element{})
		lastRow := &grid[len(grid)-1]
		for _, v := range strings.Split(line, "") {
			val, _ := strconv.Atoi(v)
			e := *NewElement(val)
			*lastRow = append(*lastRow, e)
		}
	}
	return grid
}

func isVisible(e Element) bool {
	if e.left < e.value {
		return true
	} else if e.top < e.value {
		return true
	} else if e.right < e.value {
		return true
	} else if e.bottom < e.value {
		return true
	} else {
		return false
	}
}

func getViewingDistance(grid Grid, i int, j int) int {
	count, ans := 1, 1
	n := len(grid)
	value := grid[i][j].value
	// Left to Right
	for k := j + 1; k < n && value > grid[i][k].value; k++ {
		count++
	}
	if j+count > n-1 {
		count--
	}
	ans *= count
	// Right to Left
	count = 1
	for k := j - 1; k >= 0 && value > grid[i][k].value; k-- {
		count++
	}
	if j-count < 0 {
		count--
	}
	ans *= count
	// Top to Bottom
	count = 1
	for k := i + 1; k < n && value > grid[k][j].value; k++ {
		count++
	}
	if i+count > n-1 {
		count--
	}
	ans *= count
	// Bottom to Top
	count = 1
	for k := i - 1; k >= 0 && value > grid[k][j].value; k-- {
		count++
	}
	if i-count < 0 {
		count--
	}
	ans *= count
	return ans
}
