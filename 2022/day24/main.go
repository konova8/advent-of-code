package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

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
	noExample := flag.Bool("no-example", false, "you don't want to check example")
	noInput := flag.Bool("no-input", false, "you don't want to check input")
	flag.Parse()

	if part == "1" || part == "" {
		fmt.Println("--- Running part 1 ---")
		if !*noExample {
			ansExample := foo(example)
			fmt.Println("Output Example:", ansExample)
		}
		if !*noInput {
			ansInput := foo(input)
			fmt.Println("Output Input:", ansInput)
		}
	}
	if part == "2" || part == "" {
		fmt.Println("--- Running part 2 ---")
		if !*noExample {
			ansExample := bar(example)
			fmt.Println("Output Example:", ansExample)
		}
		if !*noInput {
			ansInput := bar(input)
			fmt.Println("Output Input:", ansInput)
		}
	}
}

type Position struct {
	x int
	y int
}

func (p1 Position) Add(p2 Position) Position {
	return Position{p1.x + p2.x, p1.y + p2.y}
}

type Blizzard struct {
	pos Position
	dir rune
}

func parseInput(input string) (map[Blizzard]bool, int, int) {
	blizzards := map[Blizzard]bool{}
	lines := strings.Split(input, "\n")
	for i, line := range lines[1 : len(lines)-1] {
		for j, char := range line[1 : len(line)-1] {
			if char != '.' {
				b := Blizzard{}
				pos := Position{}
				pos.x = j
				pos.y = i
				b.pos = pos
				switch char {
				case '>':
					b.dir = 'R'
				case 'v':
					b.dir = 'D'
				case '<':
					b.dir = 'L'
				case '^':
					b.dir = 'U'
				}
				blizzards[b] = true
			}
		}
	}
	return blizzards, len(lines[0]) - 2, len(lines) - 2
}

var directions = [4]rune{'R', 'D', 'L', 'U'}
var offsets = [5]Position{{-1, 0}, {+1, 0}, {0, -1}, {0, +1}, {0, 0}}

func isClear(blizzards map[Blizzard]bool, pos Position) bool {
	for _, d := range directions {
		if blizzards[Blizzard{Position{pos.x, pos.y}, d}] {
			return false
		}
	}
	return true
}

func printMap(blizzards map[Blizzard]bool, me Position, width int, height int) {
	fmt.Println(me)
	grid := [][]string{}
	for y := 0; y < height; y++ {
		grid = append(grid, []string{})
		for x := 0; x < width; x++ {
			grid[y] = append(grid[y], ".")
			for _, d := range directions {
				if blizzards[Blizzard{Position{x, y}, d}] {
					grid[y][x] = string(d)
				}
			}
			if x == me.x && y == me.y {
				grid[y][x] = "E"
			}
		}
	}
	for i, v := range grid {
		fmt.Printf("%d ", i)
		for _, v2 := range v {
			fmt.Printf("%s", v2)
		}
		fmt.Println()
	}
}

func updateBlizzard(blizzards map[Blizzard]bool, width int, height int) map[Blizzard]bool {
	newBlizzardsMap := map[Blizzard]bool{}
	keys := make([]Blizzard, 0, len(blizzards))
	for b := range blizzards {
		keys = append(keys, b)
	}
	for _, b := range keys {
		newB := Blizzard{}
		switch b.dir {
		case 'R':
			newB.dir = b.dir
			newB.pos.y = b.pos.y
			if b.pos.x+1 <= width-1 {
				newB.pos.x = b.pos.x + 1
			} else {
				newB.pos.x = 0
			}
		case 'D':
			newB.dir = b.dir
			newB.pos.x = b.pos.x
			if b.pos.y+1 <= height-1 {
				newB.pos.y = b.pos.y + 1
			} else {
				newB.pos.y = 0
			}
		case 'L':
			newB.dir = b.dir
			newB.pos.y = b.pos.y
			if b.pos.x-1 >= 0 {
				newB.pos.x = b.pos.x - 1
			} else {
				newB.pos.x = width - 1
			}
		case 'U':
			newB.dir = b.dir
			newB.pos.x = b.pos.x
			if b.pos.y-1 >= 0 {
				newB.pos.y = b.pos.y - 1
			} else {
				newB.pos.y = height - 1
			}
		}
		newBlizzardsMap[newB] = true
	}
	return newBlizzardsMap
}

type GameState struct {
	pos  Position
	step int
}

func foo(input string) int {
	// Initial setup
	blizzards, width, height := parseInput(input)
	start := Position{x: 0, y: -1}
	end := Position{x: width - 1, y: height}
	// Start Simulation
	step := 1
	seen := map[GameState]bool{}
	Q := []GameState{{start, 0}}
	for {
		// Q.pop()
		gs := Q[0]
		Q = Q[1:]
		if gs.pos == end {
			// If reached end return step
			return gs.step - 2
		}
		if gs.step > step {
			// fmt.Println("updating Blizzard with step: ", step)
			// If change number Compute Blizzard
			blizzards = updateBlizzard(blizzards, width, height)
			step++
		}
		// Compute all possible new positions
		for _, o := range offsets {
			p := gs.pos.Add(o)
			// If ouside of box ignore (unless start or end)
			validPosition := (p.y >= 0 && p.y <= height-1 && p.x >= 0 && p.x <= width) || p == start || p == end
			if isClear(blizzards, p) && validPosition && !seen[GameState{p, (step + 1) % (width * height)}] {
				seen[GameState{p, (step + 1) % (width * height)}] = true
				Q = append(Q, GameState{p, step + 1})
			}
		}
	}
}

func bar(input string) int {
	ans := 0
	// Initial setup
	blizzards, width, height := parseInput(input)
	start := Position{x: 0, y: -1}
	end := Position{x: width - 1, y: height}
	// start -> end
	step := 1
	seen := map[GameState]bool{}
	Q := []GameState{{start, 0}}
	for {
		// Q.pop()
		gs := Q[0]
		Q = Q[1:]
		if gs.pos == end {
			// If reached end return step
			ans += gs.step - 2
			break
		}
		if gs.step > step {
			// fmt.Println("updating Blizzard with step: ", step)
			// If change number Compute Blizzard
			blizzards = updateBlizzard(blizzards, width, height)
			step++
		}
		// Compute all possible new positions
		for _, o := range offsets {
			p := gs.pos.Add(o)
			// If ouside of box ignore (unless start or end)
			validPosition := (p.y >= 0 && p.y <= height-1 && p.x >= 0 && p.x <= width) || p == start || p == end
			if isClear(blizzards, p) && validPosition && !seen[GameState{p, (step + 1) % (width * height)}] {
				seen[GameState{p, (step + 1) % (width * height)}] = true
				Q = append(Q, GameState{p, step + 1})
			}
		}
	}
	// end -> start
	step = 1
	seen = map[GameState]bool{}
	Q = []GameState{{end, 0}}
	for {
		// Q.pop()
		gs := Q[0]
		Q = Q[1:]
		if gs.pos == start {
			// If reached end return step
			ans += gs.step - 1
			break
		}
		if gs.step > step {
			// fmt.Println("updating Blizzard with step: ", step)
			// If change number Compute Blizzard
			blizzards = updateBlizzard(blizzards, width, height)
			step++
		}
		// Compute all possible new positions
		for _, o := range offsets {
			p := gs.pos.Add(o)
			// If ouside of box ignore (unless start or end)
			validPosition := (p.y >= 0 && p.y <= height-1 && p.x >= 0 && p.x <= width) || p == start || p == end
			if isClear(blizzards, p) && validPosition && !seen[GameState{p, (step + 1) % (width * height)}] {
				seen[GameState{p, (step + 1) % (width * height)}] = true
				Q = append(Q, GameState{p, step + 1})
			}
		}
	}
	// start -> end
	step = 1
	seen = map[GameState]bool{}
	Q = []GameState{{start, 0}}
	for {
		// Q.pop()
		gs := Q[0]
		Q = Q[1:]
		if gs.pos == end {
			// If reached end return step
			ans += gs.step - 1
			break
		}
		if gs.step > step {
			// fmt.Println("updating Blizzard with step: ", step)
			// If change number Compute Blizzard
			blizzards = updateBlizzard(blizzards, width, height)
			step++
		}
		// Compute all possible new positions
		for _, o := range offsets {
			p := gs.pos.Add(o)
			// If ouside of box ignore (unless start or end)
			validPosition := (p.y >= 0 && p.y <= height-1 && p.x >= 0 && p.x <= width) || p == start || p == end
			if isClear(blizzards, p) && validPosition && !seen[GameState{p, (step + 1) % (width * height)}] {
				seen[GameState{p, (step + 1) % (width * height)}] = true
				Q = append(Q, GameState{p, step + 1})
			}
		}
	}
	return ans
}
