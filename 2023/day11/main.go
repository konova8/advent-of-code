package main

import (
	"advent-of-code/util"
	_ "embed"
	"flag"
	"fmt"
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

func parseInput(str string) ([]util.Position, map[int]bool, map[int]bool) {
	galaxies := []util.Position{}
	freeColumns := map[int]bool{}
	freeRows := map[int]bool{}
	lines := strings.Split(str, "\n")

	for y := 0; y < len(lines); y++ {
		freeRows[y] = true
	}
	for x := 0; x < len(lines[0]); x++ {
		freeColumns[x] = true
	}

	for y, line := range lines {
		for x, r := range line {
			if r == '#' {
				p := util.Position{X: x, Y: y}
				galaxies = append(galaxies, p)
				freeColumns[x] = false
				freeRows[y] = false
			}
		}
	}
	for k, v := range freeColumns {
		if !v {
			delete(freeColumns, k)
		}
	}
	for k, v := range freeRows {
		if !v {
			delete(freeRows, k)
		}
	}
	return galaxies, freeColumns, freeRows
}

func part1(str string) int {
	galaxies, freeColumns, freeRows := parseInput(str)
	ans := 0
	for i, p1 := range galaxies {
		for _, p2 := range galaxies[i+1:] {
			// Free columns
			cols := 0
			for k := range freeColumns {
				if (p1.X < k && k < p2.X) || (p2.X < k && k < p1.X) {
					cols++
				}
			}
			// Free Rows
			rows := 0
			for k := range freeRows {
				if (p1.Y < k && k < p2.Y) || (p2.Y < k && k < p1.Y) {
					rows++
				}
			}
			// Compute range
			d := util.ABS(p2.X-p1.X) + util.ABS(p2.Y-p1.Y) + rows + cols
			ans += d
		}
	}
	return ans
}

func part2(str string) int {
	galaxies, freeColumns, freeRows := parseInput(str)
	ans := 0
	for i, p1 := range galaxies {
		for _, p2 := range galaxies[i+1:] {
			// Free columns
			cols := 0
			for k := range freeColumns {
				if (p1.X < k && k < p2.X) || (p2.X < k && k < p1.X) {
					cols++
				}
			}
			// Free Rows
			rows := 0
			for k := range freeRows {
				if (p1.Y < k && k < p2.Y) || (p2.Y < k && k < p1.Y) {
					rows++
				}
			}
			// Compute range
			d := util.ABS(p2.X-p1.X) + util.ABS(p2.Y-p1.Y) + rows*999_999 + cols*999_999
			ans += d
		}
	}
	return ans
}
