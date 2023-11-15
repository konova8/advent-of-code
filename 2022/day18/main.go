package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
	z int
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

func foo(input string) int {
	ans := 0
	// Parse input
	r, err := regexp.Compile(`(\d+),(\d+),(\d+)`)
	if err != nil {
		panic("regex not compilable")
	}
	cubeSet := map[Position]bool{}
	for _, line := range strings.Split(input, "\n") {
		ans += 6
		v := r.FindAllStringSubmatch(line, -1)
		if v == nil {
			panic("no match found")
		}
		x, _ := strconv.Atoi(v[0][1])
		y, _ := strconv.Atoi(v[0][2])
		z, _ := strconv.Atoi(v[0][3])
		cubeSet[Position{x, y, z}] = true
		if cubeSet[Position{x - 1, y, z}] {
			ans -= 2
		}
		if cubeSet[Position{x + 1, y, z}] {
			ans -= 2
		}
		if cubeSet[Position{x, y - 1, z}] {
			ans -= 2
		}
		if cubeSet[Position{x, y + 1, z}] {
			ans -= 2
		}
		if cubeSet[Position{x, y, z - 1}] {
			ans -= 2
		}
		if cubeSet[Position{x, y, z + 1}] {
			ans -= 2
		}
	}
	return ans
}

func bar(input string) int {
	ans := 0
	// Parse input
	r, err := regexp.Compile(`(\d+),(\d+),(\d+)`)
	if err != nil {
		panic("regex not compilable")
	}
	cubeSet := map[Position]bool{}
	var minX, maxX, minY, maxY, minZ, maxZ int = 100, -100, 100, -100, 100, -100
	for _, line := range strings.Split(input, "\n") {
		v := r.FindAllStringSubmatch(line, -1)
		if v == nil {
			panic("no match found")
		}
		x, _ := strconv.Atoi(v[0][1])
		y, _ := strconv.Atoi(v[0][2])
		z, _ := strconv.Atoi(v[0][3])
		cubeSet[Position{x, y, z}] = true
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
		if z < minZ {
			minZ = z
		}
		if z > maxZ {
			maxZ = z
		}
	}
	// If is outside the box
	isOutsideBox := func(p Position) bool {
		return p.x < minX-1 || p.x > maxX+1 || p.y < minY-1 || p.y > maxY+1 || p.z < minZ-1 || p.z > maxZ+1
	}
	// We do a BFS search and append only if we are inside the max outline
	q := []Position{{minX, minY, minZ}}
	posToCheck := Position{}
	seen := map[Position]bool{}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		seen[p] = true
		posToCheck = Position{p.x - 1, p.y, p.z}
		if cubeSet[posToCheck] {
			ans++
		} else if !seen[posToCheck] && !isOutsideBox(posToCheck) && !slices.Contains(q, posToCheck) {
			q = append(q, posToCheck)
		}
		posToCheck = Position{p.x + 1, p.y, p.z}
		if cubeSet[posToCheck] {
			ans++
		} else if !seen[posToCheck] && !isOutsideBox(posToCheck) && !slices.Contains(q, posToCheck) {
			q = append(q, posToCheck)
		}
		posToCheck = Position{p.x, p.y - 1, p.z}
		if cubeSet[posToCheck] {
			ans++
		} else if !seen[posToCheck] && !isOutsideBox(posToCheck) && !slices.Contains(q, posToCheck) {
			q = append(q, posToCheck)
		}
		posToCheck = Position{p.x, p.y + 1, p.z}
		if cubeSet[posToCheck] {
			ans++
		} else if !seen[posToCheck] && !isOutsideBox(posToCheck) && !slices.Contains(q, posToCheck) {
			q = append(q, posToCheck)
		}
		posToCheck = Position{p.x, p.y, p.z - 1}
		if cubeSet[posToCheck] {
			ans++
		} else if !seen[posToCheck] && !isOutsideBox(posToCheck) && !slices.Contains(q, posToCheck) {
			q = append(q, posToCheck)
		}
		posToCheck = Position{p.x, p.y, p.z + 1}
		if cubeSet[posToCheck] {
			ans++
		} else if !seen[posToCheck] && !isOutsideBox(posToCheck) && !slices.Contains(q, posToCheck) {
			q = append(q, posToCheck)
		}
	}
	return ans
}
