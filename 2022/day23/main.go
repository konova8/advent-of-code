package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
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
			ansExample := foo(example, 10)
			fmt.Println("Output Example:", ansExample)
		}
		if !*noInput {
			ansInput := foo(input, 10)
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

const runeDot = '.'
const runeHash = '#'

type Elf struct {
	now  complex128
	want complex128
}

func parseInput(input string) []Elf {
	elves := []Elf{}
	for i, line := range strings.Split(input, "\n") {
		for j, r := range line {
			if r == runeHash {
				c := complex(float64(i), float64(j))
				elves = append(elves, Elf{c, c})
			}
		}
	}
	return elves
}

func nearElves(elf Elf, elves []Elf) []complex128 {
	ret := []complex128{}
	for _, d := range allDirectionOffsets {
		if slices.ContainsFunc(elves, func(e Elf) bool { return e.now == (elf.now + d) }) {
			ret = append(ret, elf.now+d)
		}
	}
	return ret
}

func printMap(elves []Elf) {
	minY, minX := 1_000_000, 1_000_000
	maxY, maxX := -1_000_000, -1_000_000
	for _, e := range elves {
		maxX = max(maxX, int(imag(e.now)))
		maxY = max(maxY, int(real(e.now)))
		minX = min(minX, int(imag(e.now)))
		minY = min(minY, int(real(e.now)))
	}
	// Add padding
	minX -= 2
	minY -= 2
	maxX += 2
	maxY += 2
	fmt.Printf("\t")
	for x := minX; x <= maxX; x++ {
		fmt.Printf("%d", x)
	}
	fmt.Printf("\n")
	for y := minY; y <= maxY; y++ {
		fmt.Printf("%d\t", y)
		for x := minX; x <= maxX; x++ {
			c := complex(float64(y), float64(x))
			e := Elf{c, 0}
			var r string
			if slices.ContainsFunc(elves, func(elf Elf) bool { return elf.now == e.now }) {
				r = "#"
			} else {
				r = "."
			}
			fmt.Printf("%s", r)
		}
		fmt.Printf("\n")
	}
}

const (
	complexN  = complex128(-1 + 0i)
	complexNE = complex128(-1 + 1i)
	complexNW = complex128(-1 - 1i)
	complexS  = complex128(1 + 0i)
	complexSE = complex128(1 + 1i)
	complexSW = complex128(1 + -1i)
	complexW  = complex128(0 - 1i)
	complexE  = complex128(0 + 1i)
)

var allDirectionOffsets = [8]complex128{
	complexN,
	complexNE,
	complexNW,
	complexS,
	complexSE,
	complexSW,
	complexW,
	complexE,
}

func foo(input string, roundNumbers int) int {
	// Generate set of elves
	elves := parseInput(input)
	// Play rounds
	blockingPositions := []complex128{}
	wantPositions := map[complex128]bool{}
	direction := []complex128{complexN, complexS, complexW, complexE}
	for i_round := 0; i_round < roundNumbers; i_round++ {
		blockingPositions = []complex128{}
		wantPositions = map[complex128]bool{}
		// First half
		for i_elf, elf := range elves {
			near := nearElves(elf, elves)
			if len(near) != 0 {
				// If not alone propose a direction
				// Keep a HashSet of all possible new position, so if someone else wants to move there you delete it
				nordEmpty := !slices.Contains(near, elf.now+complexN) &&
					!slices.Contains(near, elf.now+complexNW) &&
					!slices.Contains(near, elf.now+complexNE)
				southEmpty := !slices.Contains(near, elf.now+complexS) &&
					!slices.Contains(near, elf.now+complexSE) &&
					!slices.Contains(near, elf.now+complexSW)
				westEmpty := !slices.Contains(near, elf.now+complexW) &&
					!slices.Contains(near, elf.now+complexNW) &&
					!slices.Contains(near, elf.now+complexSW)
				eastEmpty := !slices.Contains(near, elf.now+complexE) &&
					!slices.Contains(near, elf.now+complexSE) &&
					!slices.Contains(near, elf.now+complexNE)
				directionEmptyBool := [4]bool{nordEmpty, southEmpty, westEmpty, eastEmpty}
				if directionEmptyBool[(0+i_round)%4] {
					newPosition := elf.now + direction[(0+i_round)%4]
					elves[i_elf].want = newPosition
					if wantPositions[newPosition] {
						// If someone else want to move here
						blockingPositions = append(blockingPositions, newPosition)
					} else {
						wantPositions[newPosition] = true
					}
				} else if directionEmptyBool[(1+i_round)%4] {
					newPosition := elf.now + direction[(1+i_round)%4]
					elves[i_elf].want = newPosition
					if wantPositions[newPosition] {
						// If someone else want to move here
						blockingPositions = append(blockingPositions, newPosition)
					} else {
						wantPositions[newPosition] = true
					}
				} else if directionEmptyBool[(2+i_round)%4] {
					newPosition := elf.now + direction[(2+i_round)%4]
					elves[i_elf].want = newPosition
					if wantPositions[newPosition] {
						// If someone else want to move here
						blockingPositions = append(blockingPositions, newPosition)
					} else {
						wantPositions[newPosition] = true
					}
				} else if directionEmptyBool[(3+i_round)%4] {
					newPosition := elf.now + direction[(3+i_round)%4]
					elves[i_elf].want = newPosition
					if wantPositions[newPosition] {
						// If someone else want to move here
						blockingPositions = append(blockingPositions, newPosition)
					} else {
						wantPositions[newPosition] = true
					}
				}
			}
		}
		// Second half
		for i := range elves {
			if slices.Contains(blockingPositions, elves[i].want) {
				// Freeze elves that want to move in a position inside blockingPoitions
				elves[i].want = elves[i].now
			}
			// Effectively move elves
			elves[i].now = elves[i].want
		}
	}
	// Compute empty tiles
	minY, minX := 1_000_000, 1_000_000
	maxY, maxX := -1_000_000, -1_000_000
	for _, e := range elves {
		maxX = max(maxX, int(real(e.now)))
		maxY = max(maxY, int(imag(e.now)))
		minX = min(minX, int(real(e.now)))
		minY = min(minY, int(imag(e.now)))
	}
	return ((maxX - minX + 1) * (maxY - minY + 1)) - len(elves)
}

func bar(input string) int {
	// Generate set of elves
	elves := parseInput(input)
	// Play rounds
	blockingPositions := []complex128{}
	wantPositions := map[complex128]bool{0 + 0i: true}
	direction := []complex128{complexN, complexS, complexW, complexE}
	i_round := 0
	for i_round = 0; len(wantPositions) != 0; i_round++ {
		blockingPositions = []complex128{}
		wantPositions = map[complex128]bool{}
		// First half
		for i_elf, elf := range elves {
			near := nearElves(elf, elves)
			if len(near) != 0 {
				// If not alone propose a direction
				// Keep a HashSet of all possible new position, so if someone else wants to move there you delete it
				nordEmpty := !slices.Contains(near, elf.now+complexN) &&
					!slices.Contains(near, elf.now+complexNW) &&
					!slices.Contains(near, elf.now+complexNE)
				southEmpty := !slices.Contains(near, elf.now+complexS) &&
					!slices.Contains(near, elf.now+complexSE) &&
					!slices.Contains(near, elf.now+complexSW)
				westEmpty := !slices.Contains(near, elf.now+complexW) &&
					!slices.Contains(near, elf.now+complexNW) &&
					!slices.Contains(near, elf.now+complexSW)
				eastEmpty := !slices.Contains(near, elf.now+complexE) &&
					!slices.Contains(near, elf.now+complexSE) &&
					!slices.Contains(near, elf.now+complexNE)
				directionEmptyBool := [4]bool{nordEmpty, southEmpty, westEmpty, eastEmpty}
				if directionEmptyBool[(0+i_round)%4] {
					newPosition := elf.now + direction[(0+i_round)%4]
					elves[i_elf].want = newPosition
					if wantPositions[newPosition] {
						// If someone else want to move here
						blockingPositions = append(blockingPositions, newPosition)
					} else {
						wantPositions[newPosition] = true
					}
				} else if directionEmptyBool[(1+i_round)%4] {
					newPosition := elf.now + direction[(1+i_round)%4]
					elves[i_elf].want = newPosition
					if wantPositions[newPosition] {
						// If someone else want to move here
						blockingPositions = append(blockingPositions, newPosition)
					} else {
						wantPositions[newPosition] = true
					}
				} else if directionEmptyBool[(2+i_round)%4] {
					newPosition := elf.now + direction[(2+i_round)%4]
					elves[i_elf].want = newPosition
					if wantPositions[newPosition] {
						// If someone else want to move here
						blockingPositions = append(blockingPositions, newPosition)
					} else {
						wantPositions[newPosition] = true
					}
				} else if directionEmptyBool[(3+i_round)%4] {
					newPosition := elf.now + direction[(3+i_round)%4]
					elves[i_elf].want = newPosition
					if wantPositions[newPosition] {
						// If someone else want to move here
						blockingPositions = append(blockingPositions, newPosition)
					} else {
						wantPositions[newPosition] = true
					}
				}
			}
		}
		// Second half
		for i := range elves {
			if slices.Contains(blockingPositions, elves[i].want) {
				// Freeze elves that want to move in a position inside blockingPoitions
				elves[i].want = elves[i].now
			}
			// Effectively move elves
			elves[i].now = elves[i].want
		}
	}
	return i_round
}
