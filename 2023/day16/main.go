package main

import (
	"advent-of-code/util"
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/atotto/clipboard"
)

//go:embed example.txt
var example string

//go:embed example2.txt
var example2 string

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
	example2 = strings.TrimRight(example2, "\n")
	if len(example2) == 0 {
		example2 = example
	}
	p := util.Position{}
	_ = p
}

func main() {
	var part string
	flag.StringVar(&part, "part", "", "part 1 or 2")
	noExample := flag.Bool("no-example", false, "you don't want to check example")
	noInput := flag.Bool("no-input", false, "you don't want to check input")
	noE := flag.Bool("noe", false, "you don't want to check example")
	noI := flag.Bool("noi", false, "you don't want to check input")
	flag.Parse()

	if part == "1" || part == "" {
		fmt.Println("--- Running part 1 ---")
		if !*noExample && !*noE {
			s := time.Now()
			ansExample := fmt.Sprint(part1(example))
			fmt.Printf("Output Example: %s\n", ansExample)
			fmt.Printf("Computed in %v\n", time.Now().Sub(s))
			clipboard.WriteAll(ansExample)
		}
		if !*noInput && !*noI {
			s := time.Now()
			ansInput := fmt.Sprint(part1(input))
			fmt.Printf("Output Input: %s\n", ansInput)
			fmt.Printf("Computed in %v\n", time.Now().Sub(s))
			clipboard.WriteAll(ansInput)
		}
	}
	if part == "2" || part == "" {
		fmt.Println("--- Running part 2 ---")
		if !*noExample && !*noE {
			s := time.Now()
			ansExample := fmt.Sprint(part2(example))
			fmt.Printf("Output Example: %s\n", ansExample)
			fmt.Printf("Computed in %v\n", time.Now().Sub(s))
			clipboard.WriteAll(ansExample)
		}
		if !*noInput && !*noI {
			s := time.Now()
			ansInput := fmt.Sprint(part2(input))
			fmt.Printf("Output Input: %s\n", ansInput)
			fmt.Printf("Computed in %v\n", time.Now().Sub(s))
			clipboard.WriteAll(ansInput)
		}
	}
}

type Direction int

const (
	DirU = iota
	DirR
	DirD
	DirL
)

var dirToOffset = map[Direction]util.Position{
	DirU: {X: 0, Y: -1},
	DirR: {X: 1, Y: 0},
	DirD: {X: 0, Y: 1},
	DirL: {X: -1, Y: 0},
}

type Grid struct {
	grid   [][]byte
	height int
	width  int
}

type Beam struct {
	pos util.Position
	dir Direction
}

const (
	dot       = byte('.')
	mirror1   = byte('/')
	mirror2   = byte('\\')
	splitter1 = byte('|')
	splitter2 = byte('-')
)

// Mirror /
var newDirMirror1 = map[Direction]Direction{
	DirU: DirR,
	DirR: DirU,
	DirD: DirL,
	DirL: DirD,
}

// Mirror \
var newDirMirror2 = map[Direction]Direction{
	DirU: DirL,
	DirR: DirD,
	DirD: DirR,
	DirL: DirU,
}

// Splitter |
var newDirSplitter1 = map[Direction][]Direction{
	DirU: {DirU},
	DirR: {DirU, DirD},
	DirD: {DirD},
	DirL: {DirU, DirD},
}

// Splitter -
var newDirSplitter2 = map[Direction][]Direction{
	DirU: {DirR, DirL},
	DirR: {DirR},
	DirD: {DirR, DirL},
	DirL: {DirL},
}

func newPosition(g Grid, b Beam) []Beam {
	switch g.grid[b.pos.Y][b.pos.X] {
	case dot:
		return []Beam{{
			pos: b.pos.Add(dirToOffset[b.dir]),
			dir: b.dir,
		}}
	case mirror1:
		return []Beam{{
			pos: b.pos.Add(dirToOffset[newDirMirror1[b.dir]]),
			dir: newDirMirror1[b.dir],
		}}
	case mirror2:
		return []Beam{{
			pos: b.pos.Add(dirToOffset[newDirMirror2[b.dir]]),
			dir: newDirMirror2[b.dir],
		}}
	case splitter1:
		tmp := []Beam{}
		for _, d := range newDirSplitter1[b.dir] {
			tmp = append(tmp, Beam{
				pos: b.pos.Add(dirToOffset[d]),
				dir: d,
			})
		}
		return tmp
	case splitter2:
		tmp := []Beam{}
		for _, d := range newDirSplitter2[b.dir] {
			tmp = append(tmp, Beam{
				pos: b.pos.Add(dirToOffset[d]),
				dir: d,
			})
		}
		return tmp
	}
	fmt.Printf("g: %+v\n", g)
	fmt.Printf("b: %+v\n", b)
	panic("error in newPosition")
}

func printGrid(g Grid, e []util.Position) {
	for y, v := range g.grid {
		for x, b := range v {
			if slices.Contains(e, util.Position{X: x, Y: y}) {
				fmt.Printf("#")
			} else {
				fmt.Printf(string(b))
			}
		}
		fmt.Println()
	}
}

func simulation(grid Grid, start Beam) int {
	energizedTiles := []util.Position{}
	seenBeams := map[Beam]bool{}
	beams := []Beam{start}
	for len(beams) > 0 {
		b := beams[0]
		beams = beams[1:]
		seenBeams[b] = true
		energizedTiles = append(energizedTiles, b.pos)
		for _, newB := range newPosition(grid, b) {
			if newB.pos.Y >= 0 && newB.pos.X >= 0 &&
				newB.pos.Y < grid.height && newB.pos.X < grid.width &&
				!seenBeams[newB] {
				beams = append(beams, newB)
			}
		}
	}
	// printGrid(grid, energizedTiles)
	mapEnergizedTiles := map[util.Position]bool{}
	for _, p := range energizedTiles {
		mapEnergizedTiles[p] = true
	}
	return len(mapEnergizedTiles)
}

func part1(str string) int {
	grid := Grid{}
	lines := strings.Split(str, "\n")
	for _, l := range lines {
		grid.grid = append(grid.grid, []byte(l))
	}
	grid.height = len(lines)
	grid.width = len(lines[0])
	start := Beam{pos: util.Position{X: 0, Y: 0}, dir: DirR}
	return simulation(grid, start)
}

func part2(str string) int {
	grid := Grid{}
	lines := strings.Split(str, "\n")
	for _, l := range lines {
		grid.grid = append(grid.grid, []byte(l))
	}
	grid.height = len(lines)
	grid.width = len(lines[0])
	startingBeams := []Beam{}
	for x := 0; x < grid.width; x++ {
		b := Beam{}
		b.pos.X = x
		b.pos.Y = 0
		b.dir = DirD
		startingBeams = append(startingBeams, b)
		b.pos.Y = grid.height - 1
		b.dir = DirU
		startingBeams = append(startingBeams, b)
	}
	for y := 0; y < grid.width; y++ {
		b := Beam{}
		b.pos.Y = y
		b.pos.X = 0
		b.dir = DirR
		startingBeams = append(startingBeams, b)
		b.pos.X = grid.width - 1
		b.dir = DirL
		startingBeams = append(startingBeams, b)
	}
	best := 0
	for _, b := range startingBeams {
		best = max(best, simulation(grid, b))
	}
	return best
}
