package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type RockLine []Position

type Position struct {
	x int
	y int
}

//go:embed example.txt
var example string

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
}

func main() {
	var part string
	flag.StringVar(&part, "part", "", "part 1 or 2")
	flag.Parse()

	if part == "1" || part == "" {
		fmt.Println("--- Running part 1 ---")
		ansExample := foo(example)
		fmt.Println("Output Example:", ansExample)
		ansInput := foo(input)
		fmt.Println("Output Input:", ansInput)
	}
	if part == "2" || part == "" {
		fmt.Println("--- Running part 2 ---")
		ansExample := bar(example)
		fmt.Println("Output Example:", ansExample)
		ansInput := bar(input)
		fmt.Println("Output Input:", ansInput)
	}
}

func foo(input string) int {
	ans := 0
	rocks := parseInput(input)
	grid := map[Position]rune{}
	start := Position{500, 0}
	grid[start] = '+'

	// Draw grid
	for _, rock := range rocks {
		for j := 1; j < len(rock); j++ {
			prePos := rock[j-1]
			// grid[prePos] = '#'
			nowPos := rock[j]
			if prePos.x == nowPos.x {
				// Draw line along y axis
				for k := min(prePos.y, nowPos.y); k <= max(prePos.y, nowPos.y); k++ {
					grid[Position{prePos.x, k}] = '#'
				}
			} else if prePos.y == nowPos.y {
				// Draw line along x axis
				for k := min(prePos.x, nowPos.x); k <= max(prePos.x, nowPos.x); k++ {
					grid[Position{k, prePos.y}] = '#'
				}
			}
		}
	}

	// Output sand until resulting position is y = -1
	pos := Position{}
	for i := 1; pos.y != -1; i++ {
		pos = computePosition(start, grid)
		// Update Grid or go out
		if pos.y == -1 {
			ans = i - 1
		} else {
			grid[pos] = 'o'
		}
	}

	return ans
}

func bar(input string) int {
	ans := 0
	rocks := parseInput(input)
	grid := map[Position]rune{}
	start := Position{500, 0}
	grid[start] = '+'

	// Draw grid
	for _, rock := range rocks {
		for j := 1; j < len(rock); j++ {
			prePos := rock[j-1]
			// grid[prePos] = '#'
			nowPos := rock[j]
			if prePos.x == nowPos.x {
				// Draw line along y axis
				for k := min(prePos.y, nowPos.y); k <= max(prePos.y, nowPos.y); k++ {
					grid[Position{prePos.x, k}] = '#'
				}
			} else if prePos.y == nowPos.y {
				// Draw line along x axis
				for k := min(prePos.x, nowPos.x); k <= max(prePos.x, nowPos.x); k++ {
					grid[Position{k, prePos.y}] = '#'
				}
			}
		}
	}

	// Add bedrock layer
	minX := 100_000
	minY := 100_000
	maxX := -100_000
	maxY := -100_000
	for key := range grid {
		minX = min(minX, key.x)
		minY = min(minY, key.y)
		maxX = max(maxX, key.x)
		maxY = max(maxY, key.y)
	}
	for x := minX - maxY; x < maxX+maxY; x++ {
		grid[Position{x, maxY + 2}] = '#'
	}

	// Output sand until resulting position is y = -1
	pos := Position{}
	for i := 1; pos != start; i++ {
		pos = computePosition(start, grid)
		// Update Grid or go out
		if pos == start {
			ans = i
		} else {
			grid[pos] = 'o'
		}
	}

	return ans
}

func parseInput(input string) []RockLine {
	ret := []RockLine{}
	for _, line := range strings.Split(input, "\n") {
		tmpRock := RockLine{}
		for _, elem := range strings.Split(line, " -> ") {
			tmp := strings.Split(elem, ",")
			x, _ := strconv.Atoi(tmp[0])
			y, _ := strconv.Atoi(tmp[1])
			tmpRock = append(tmpRock, Position{x, y})
		}
		ret = append(ret, tmpRock)
	}
	return ret
}

func min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func computePosition(start Position, grid map[Position]rune) Position {
	maxHeigh := 0
	for key := range grid {
		maxHeigh = max(maxHeigh, key.y)
	}
	curr := start
	i := 0
	for i = start.y; i <= maxHeigh+1; i++ {
		if grid[Position{curr.x, curr.y + 1}] == 0 {
			curr.y++
		} else if grid[Position{curr.x - 1, curr.y + 1}] == 0 {
			curr.x--
			curr.y++
		} else if grid[Position{curr.x + 1, curr.y + 1}] == 0 {
			curr.x++
			curr.y++
		} else if i < maxHeigh {
			break
		}
	}
	if i > maxHeigh {
		curr.y = -1
		return curr
	} else {
		return curr
	}
}

func printGrid(grid map[Position]rune) {
	minX := 100_000
	minY := 100_000
	maxX := -100_000
	maxY := -100_000
	// Compute min and max
	for key := range grid {
		minX = min(minX, key.x)
		minY = min(minY, key.y)
		maxX = max(maxX, key.x)
		maxY = max(maxY, key.y)
	}
	// Print Grid
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if grid[Position{x, y}] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(string(grid[Position{x, y}]))
			}
		}
		fmt.Println(" ", y)
	}
}
