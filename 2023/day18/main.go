package main

import (
	"advent-of-code/util"
	_ "embed"
	"flag"
	"fmt"
	"strconv"
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

func solve(vertexes []util.Position) int {
	area := 0
	for i := range vertexes {
		before := i % len(vertexes)
		after := (i + 1) % len(vertexes)
		area += vertexes[before].X * vertexes[after].Y
		area -= vertexes[after].X * vertexes[before].Y
	}
	return area / 2
}

func part1(str string) int {
	vertexes := []util.Position{}
	posNow := util.Position{X: 0, Y: 0}
	perimeter := 0
	for _, line := range strings.Split(str, "\n") {
		tmp := strings.Split(line, " ")
		dir := tmp[0]
		n := util.Atoi(tmp[1])
		perimeter += n
		switch dir {
		case "R":
			posNow = posNow.Add(util.Position{X: n, Y: 0})
			vertexes = append(vertexes, posNow)
		case "U":
			posNow = posNow.Add(util.Position{X: 0, Y: -n})
			vertexes = append(vertexes, posNow)
		case "D":
			posNow = posNow.Add(util.Position{X: 0, Y: n})
			vertexes = append(vertexes, posNow)
		case "L":
			posNow = posNow.Add(util.Position{X: -n, Y: 0})
			vertexes = append(vertexes, posNow)
		}
	}
	ans := solve(vertexes)
	return ans + (perimeter / 2) + 1
}

func part2(str string) int {
	vertexes := []util.Position{}
	posNow := util.Position{X: 0, Y: 0}
	perimeter := 0
	for _, line := range strings.Split(str, "\n") {
		tmp := strings.Split(line, " ")
		dir := tmp[2][len(tmp[2])-2]
		dec, _ := strconv.ParseInt(tmp[2][2:len(tmp[2])-2], 16, 64)
		n := int(dec)
		perimeter += n
		switch dir {
		// case "R":
		case '0':
			posNow = posNow.Add(util.Position{X: n, Y: 0})
			vertexes = append(vertexes, posNow)
		// case "U":
		case '3':
			posNow = posNow.Add(util.Position{X: 0, Y: -n})
			vertexes = append(vertexes, posNow)
		// case "D":
		case '1':
			posNow = posNow.Add(util.Position{X: 0, Y: n})
			vertexes = append(vertexes, posNow)
		// case "L":
		case '2':
			posNow = posNow.Add(util.Position{X: -n, Y: 0})
			vertexes = append(vertexes, posNow)
		}
	}
	ans := solve(vertexes)
	return ans + (perimeter / 2) + 1
}
