package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"github.com/konova8/advent-of-code/util"
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
			ansExample := fmt.Sprint(part2(example2))
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

type Node util.Position

type Graph map[Node][]Node

var openLeft []rune = []rune{
	'-',
	'J',
	'7',
	'S',
}
var openRight []rune = []rune{
	'-',
	'L',
	'F',
	'S',
}
var openTop []rune = []rune{
	'|',
	'J',
	'L',
	'S',
}
var openBottom []rune = []rune{
	'|',
	'F',
	'7',
	'S',
}

func parseInput(str string) (Graph, Node) {
	graph := Graph{}
	start := Node{}
	lines := strings.Split(str, "\n")
	emptyRow := strings.Repeat(".", len(lines[0]))
	lines = append([]string{emptyRow}, lines...)
	lines = append(lines, emptyRow)
	for i, v := range lines {
		lines[i] = "." + v + "."
	}
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[0]); x++ {
			r := rune(lines[y][x])
			n := Node{X: x, Y: y}
			if r == 'S' {
				start = n
			}
			graph[n] = []Node{}
			if r != '.' {
				if slices.Contains(openRight, r) &&
					slices.Contains(openLeft, rune(lines[y][x+1])) {
					newN := Node{X: x + 1, Y: y}
					graph[n] = append(graph[n], newN)
				}
				if slices.Contains(openLeft, r) &&
					slices.Contains(openRight, rune(lines[y][x-1])) {
					newN := Node{X: x - 1, Y: y}
					graph[n] = append(graph[n], newN)
				}
				if slices.Contains(openBottom, r) &&
					slices.Contains(openTop, rune(lines[y+1][x])) {
					newN := Node{X: x, Y: y + 1}
					graph[n] = append(graph[n], newN)
				}
				if slices.Contains(openTop, r) &&
					slices.Contains(openBottom, rune(lines[y-1][x])) {
					newN := Node{X: x, Y: y - 1}
					graph[n] = append(graph[n], newN)
				}
			}
		}
	}
	return graph, start
}

func part1(str string) int {
	graph, start := parseInput(str)
	depth := 0
	seen := map[Node]bool{}
	invalidNode := Node{-1, -1}
	q := []Node{start, invalidNode}
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		if n == invalidNode {
			q = append(q, invalidNode)
			if q[0] == invalidNode {
				break
			} else {
				depth++
				continue
			}
		}
		for _, newN := range graph[n] {
			if !seen[newN] {
				seen[newN] = true
				q = append(q, newN)
			}
		}
	}
	return depth
}

func newMap(str string) []string {
	modified := true
	for modified {
		modified = false
		lines := strings.Split(str, "\n")
		for y := 0; y < len(lines); y++ {
			for x := 0; x < len(lines[y])-1; x++ {
				r := rune(lines[y][x])
				if r != '.' {
					count := 0
					connectedRight := (x+1) < len(lines[y]) && slices.Contains(openRight, r) && slices.Contains(openLeft, rune(lines[y][x+1]))
					connectedLeft := (x-1) >= 0 && slices.Contains(openLeft, r) && slices.Contains(openRight, rune(lines[y][x-1]))
					connectedBottom := (y+1) < len(lines) && slices.Contains(openBottom, r) && slices.Contains(openTop, rune(lines[y+1][x]))
					connectedTop := (y-1) >= 0 && slices.Contains(openTop, r) && slices.Contains(openBottom, rune(lines[y-1][x]))
					if connectedRight {
						count++
					}
					if connectedLeft {
						count++
					}
					if connectedBottom {
						count++
					}
					if connectedTop {
						count++
					}
					if count < 2 {
						lines[y] = lines[y][:x] + "." + lines[y][x+1:]
						modified = true
					}
				}
			}
		}
		str = strings.Join(lines, "\n")
	}
	lines := strings.Split(str, "\n")
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y])-1; x++ {
			r1 := rune(lines[y][x])
			r2 := rune(lines[y][x+1])
			if slices.Contains(openRight, r1) &&
				slices.Contains(openLeft, r2) {
				lines[y] = lines[y][:x+1] + "-" + lines[y][x+1:]
				x++
			} else {
				lines[y] = lines[y][:x+1] + "," + lines[y][x+1:]
				x++
			}
		}
	}
	emptyRow := strings.Repeat(",", len(lines[0]))
	newLines := []string{}
	for i, v := range lines {
		if i != len(lines)-1 {
			newLines = append(newLines, v, emptyRow)
		} else {
			newLines = append(newLines, v)
		}
	}
	lines = newLines
	for y := 0; y < len(lines)-1; y += 2 {
		for x := 0; x < len(lines[0]); x++ {
			r1 := rune(lines[y][x])
			r2 := rune(lines[y+2][x])
			if slices.Contains(openBottom, r1) &&
				slices.Contains(openTop, r2) {
				lines[y+1] = lines[y+1][:x] + "|" + lines[y+1][x+1:]
			}
		}
	}
	emptyRow = strings.Repeat("0", len(lines[0]))
	lines = append([]string{emptyRow}, lines...)
	lines = append(lines, emptyRow)
	for i, v := range lines {
		lines[i] = "0" + v + "0"
	}
	return lines
}

var offsets []util.Position = []util.Position{
	{X: -1, Y: +0},
	{X: +1, Y: +0},
	{X: +0, Y: -1},
	{X: +0, Y: +1},
}

func part2(str string) int {
	lines := newMap(str)
	modified := true
	for modified {
		modified = false
		for y := 1; y < len(lines)-1; y++ {
			l := lines[y]
			for x := 1; x < len(l)-1; x++ {
				r := l[x]
				for _, o := range offsets {
					if lines[y+o.Y][x+o.X] == '0' && (r == '.' || r == ',') {
						lines[y] = lines[y][:x] + "0" + lines[y][x+1:]
						modified = true
						break
					}
				}
			}
		}
	}
	count := 0
	for _, l := range lines {
		for _, r := range l {
			if r == '.' {
				count++
			}
		}
	}
	return count
}
