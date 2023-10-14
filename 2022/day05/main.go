package main

import (
	"C"
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Instruction struct {
	Quantity int
	From     int
	To       int
}

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		fmt.Println("Output:", ans)
	} else if part == 2 {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

func part1(input string) string {
	situation, moves := parseInput(input)
	ans := ""

	for _, m := range moves {
		for i := 0; i < m.Quantity; i++ {
			situation = doMove(situation, m.From, m.To, 1)
		}
	}

	for i := 0; i < len(situation); i++ {
		ans += string(situation[i][len(situation[i])-1])
	}

	return ans
}

func part2(input string) string {
	situation, moves := parseInput(input)
	ans := ""

	for _, m := range moves {
		situation = doMove(situation, m.From, m.To, m.Quantity)
	}

	for i := 0; i < len(situation); i++ {
		ans += string(situation[i][len(situation[i])-1])
	}

	return ans
}

func parseInput(input string) ([][]byte, []Instruction) {
	initSituation := [][]byte{}
	moves := []Instruction{}
	lines := strings.Split(input, "\n")
	separator := slices.Index(lines, "")
	linesInitialSituation := lines[:separator-1]
	linesMoves := lines[separator+1:]
	slices.Reverse(linesInitialSituation)
	numberOfColumns := (len(linesInitialSituation[0]) + 1) / 4
	for i := 0; i < numberOfColumns; i++ {
		initSituation = append(initSituation, []byte{})
	}
	for _, line := range linesInitialSituation {
		for i := 0; i < numberOfColumns; i++ {
			e := line[1+4*i]
			if e != ' ' {
				initSituation[i] = append(initSituation[i], e)
			}
		}
	}
	for _, line := range linesMoves {
		arr := strings.Split(line, " ")
		m := Instruction{}
		m.Quantity, _ = strconv.Atoi(arr[1])
		m.From, _ = strconv.Atoi(arr[3])
		m.To, _ = strconv.Atoi(arr[5])
		moves = append(moves, m)
	}
	return initSituation, moves
}

func doMove(situation [][]byte, from int, to int, quantity int) [][]byte {
	buffer := []byte{}
	lenFrom := len(situation[from-1])

	// Taking buffer
	buffer = situation[from-1][lenFrom-quantity:]
	situation[from-1] = situation[from-1][:lenFrom-quantity]

	// Landing buffer
	situation[to-1] = append(situation[to-1], buffer...)

	return situation
}
