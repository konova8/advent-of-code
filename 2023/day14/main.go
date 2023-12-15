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

var offsets = []util.Position{
	{X: 0, Y: -1},
	{X: -1, Y: 0},
	{X: 0, Y: 1},
	{X: 1, Y: 0},
}

func part1(str string) int {
	ans := 0
	spheres := map[util.Position]bool{}
	rocks := map[util.Position]bool{}
	lines := strings.Split(str, "\n")
	height := len(lines)
	for y, line := range lines {
		for x, r := range line {
			if r == 'O' {
				spheres[util.Position{X: x, Y: y}] = true
			} else if r == '#' {
				rocks[util.Position{X: x, Y: y}] = true
			}
		}
	}
	modified := true
	for modified {
		modified = false
		for p := range spheres {
			pN := p.Add(offsets[0])
			for p.Y > 0 && !spheres[pN] && !rocks[pN] {
				delete(spheres, p)
				spheres[pN] = true
				pN = p.Add(offsets[0])
				modified = true
			}
		}
	}
	for p := range spheres {
		ans += height - p.Y
	}
	return ans
}

func printGrid(spheres, rocks map[util.Position]bool, width, height int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			p := util.Position{X: x, Y: y}
			if spheres[p] {
				fmt.Printf("O")
			} else if rocks[p] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}

const SPHERES int = 2031 // Maximum number of spheres

func part2(str string) int {
	ans := 0
	spheres := map[util.Position]bool{}
	rocks := map[util.Position]bool{}
	lines := strings.Split(str, "\n")
	width := len(lines[0])
	height := len(lines)
	for y, line := range lines {
		for x, r := range line {
			if r == 'O' {
				spheres[util.Position{X: x, Y: y}] = true
			} else if r == '#' {
				rocks[util.Position{X: x, Y: y}] = true
			}
		}
	}
	N := 1_000_000_000
	count := 0
	alreadyDone := false
	cache := map[[SPHERES]util.Position]bool{}
	startCycle := -1
	endCycle := -1
	sphereCycle := [SPHERES]util.Position{}
	for count < N {
		keys := []util.Position{}
		i := 0
		for p := range spheres {
			keys = append(keys, p)
			i++
		}
		slices.SortFunc(keys, func(p1, p2 util.Position) int {
			if p1.X == p2.X {
				return p1.Y - p2.Y
			}
			return p1.X - p2.X
		})
		k := [SPHERES]util.Position{}
		for j, p := range keys {
			k[j] = p
		}
		if startCycle != -1 && !alreadyDone {
			if k == sphereCycle {
				alreadyDone = true
				endCycle = count
				withoutPrefix := N - startCycle + 1
				toAdd := withoutPrefix % (endCycle - startCycle)
				count = N - toAdd + 1
			}
		} else if _, ok := cache[k]; ok {
			sphereCycle = k
			startCycle = count
		}
		cache[k] = true
		for _, o := range offsets {
			modified := true
			for modified {
				modified = false
				newSpheres := map[util.Position]bool{}
				for p := range spheres {
					pN := p.Add(o)
					pBak := p
					for pN.X >= 0 && pN.Y >= 0 && pN.X < width && pN.Y < height &&
						!spheres[pN] && !rocks[pN] {
						modified = true
						pBak = pN
						pN = pN.Add(o)
					}
					newSpheres[pBak] = true
				}
				spheres = newSpheres
			}
		}
		count++
	}
	for p := range spheres {
		ans += height - p.Y
	}
	return ans
}
