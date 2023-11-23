package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"strings"
)

type Rock []string

type Pair struct {
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
	return baz(input, 2022)
}

func bar(input string) int {
	return baz(input, 1_000_000_000_000)
}

func baz(input string, n int) int {
	chamber := [][7]rune{}
	chamber = append(chamber,
		[7]rune{'.', '.', '.', '.', '.', '.', '.'},
		[7]rune{'.', '.', '.', '.', '.', '.', '.'},
		[7]rune{'.', '.', '.', '.', '.', '.', '.'})
	rocks := [5]Rock{
		{"####"},              // Minus
		{".#.", "###", ".#."}, // Plus
		{"###", "..#", "..#"}, // L
		{"#", "#", "#", "#"},  // I
		{"##", "##"},          // Square
	}
	type GameState struct {
		indexRock         int
		indexJet          int
		shapeHeightOffset [7]int
	}
	type ValueSeen struct {
		i             int
		prevMaxHeight int
	}
	seen := map[GameState]ValueSeen{}
	k := 0
	lineRemoved := 0
	foundCycle := false
	for i := 0; i < n; i++ {
		// Type of Rock
		rock := rocks[i%5]
		// Compute Shape of Chamber
		shape := offsetHeight(chamber)
		gs := GameState{
			indexRock:         i % 5,
			indexJet:          k,
			shapeHeightOffset: shape,
		}
		if prevSeen, ok := seen[gs]; ok && !foundCycle {
			// Here if we found a Cycle
			foundCycle = true
			// Compute difference of height from previews seen state to current state
			diffHeight := maxHeight(chamber) - prevSeen.prevMaxHeight
			// Compute the cycle length for rocks
			rockCycleLength := i - prevSeen.i
			// Compute how many cycles we still need to do
			countCyclesLeft := ((n - 1) - i) % rockCycleLength
			timesCycles := ((n - 1) - i) / rockCycleLength
			// Add to lineRemoved for counting at the end additional height
			if diffHeight < 0 {
				diffHeight = -diffHeight
			}
			lineRemoved += diffHeight * timesCycles
			// Finish to compute all remaning cycles
			i = (n - 1) - countCyclesLeft
		}
		seen[gs] = ValueSeen{i: i, prevMaxHeight: lineRemoved + maxHeight(chamber)}
		// Add height to chamber
		maxH := maxHeight(chamber)
		for j := len(chamber); j < maxH+3+len(rock); j++ {
			chamber = append(chamber, [7]rune{'.', '.', '.', '.', '.', '.', '.'})
		}
		// Setup indexes
		dlCornerRock := Pair{2, maxH + 3}
		for y := maxH + 3; y >= 0; y-- {
			move := input[k%len(input)] // < or >
			// Being push by a jet of hot gas
			if move == '>' {
				if !checkCollision(chamber, rock, Pair{dlCornerRock.x + 1, dlCornerRock.y}) {
					dlCornerRock.x++
				}
			} else {
				if !checkCollision(chamber, rock, Pair{dlCornerRock.x - 1, dlCornerRock.y}) {
					dlCornerRock.x--
				}
			}
			k = (k + 1) % len(input)
			// Falling one unit down
			if !checkCollision(chamber, rock, Pair{dlCornerRock.x, dlCornerRock.y - 1}) {
				dlCornerRock.y--
			} else {
				chamber = fixRock(chamber, rock, dlCornerRock)
				break
			}
		}
	}
	// printChamber(chamber)
	return lineRemoved + maxHeight(chamber)
}

func keepLastN(chamber [][7]rune, n int) ([][7]rune, int) {
	return chamber[len(chamber)-n:], len(chamber) - n
}

func cleanChamber(chamber [][7]rune) ([][7]rune, int) {
	lineRemoved := 0
	for i := len(chamber) - 1; i >= 0; i-- {
		if chamber[i] == [7]rune{'#', '#', '#', '#', '#', '#', '#'} {
			lineRemoved = i
			break
		}
	}
	return chamber[lineRemoved:], lineRemoved
}

func printPartOfChamber(chamber [][7]rune, rows int) {
	fmt.Printf("Part of Chamber:\n")
	for i := len(chamber) - 1; i >= len(chamber)-rows; i-- {
		fmt.Printf("|")
		for j := 0; j < 7; j++ {
			fmt.Printf("%s", string(chamber[i][j]))
		}
		fmt.Printf("|\n")
	}
	fmt.Printf("+-------+\n")
}

func printChamber(chamber [][7]rune) {
	fmt.Printf("Chamber:\n")
	for i := len(chamber) - 1; i >= 0; i-- {
		fmt.Printf("|")
		for j := 0; j < 7; j++ {
			fmt.Printf("%s", string(chamber[i][j]))
		}
		fmt.Printf("|\n")
	}
	fmt.Printf("+-------+\n")
}

func fixRock(chamber [][7]rune, rock Rock, position Pair) [][7]rune {
	for y := 0; y < len(rock); y++ {
		for x := 0; x < len(rock[0]); x++ {
			if rock[y][x] == '#' { // Part of rock?
				chamber[position.y+y][position.x+x] = '#'
			}
		}
	}
	return chamber
}

func checkCollision(chamber [][7]rune, rock Rock, position Pair) bool {
	for y := 0; y < len(rock); y++ {
		for x := 0; x < len(rock[0]); x++ {
			if rock[y][x] == '#' {
				relativeX := position.x + x
				relativeY := position.y + y
				if relativeY < 0 || relativeX < 0 || relativeX > 6 || chamber[relativeY][relativeX] != '.' {
					return true
				}
			}
		}
	}
	return false
}

// Take a chamber and return the column with the max height
func maxHeight(chamber [][7]rune) int {
	hCol := []int{}
	for j := 0; j < 7; j++ {
		hCol = append(hCol, 0)
	}
	for i, line := range chamber {
		for j := 0; j < 7; j++ {
			if line[j] == '#' {
				hCol[j] = i + 1
			}
		}
	}
	return slices.Max(hCol)
}

// Return the height of each column of the chamber
func offsetHeight(chamber [][7]rune) [7]int {
	hCol := [7]int{}
	maximum := 0
	for j := 0; j < 7; j++ {
		hCol[j] = 0
	}
	for i, line := range chamber {
		for j := 0; j < 7; j++ {
			if line[j] == '#' {
				hCol[j] = i + 1
				if i+1 > maximum {
					maximum = i + 1
				}
			}
		}
	}
	for j := 0; j < 7; j++ {
		hCol[j] -= maximum
	}
	return hCol
}
