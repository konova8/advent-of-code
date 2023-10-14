package main

import (
	"C"
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

/*
 *  	X Y Z
 *  A 3 6 0
 *  B 0 3 6
 *  C 6 0 3
 */
var tablePossibleOutcame = [][]int{{3, 6, 0}, {0, 3, 6}, {6, 0, 3}}

/*
 *  	X Y Z
 *  A C A B
 *  B A B C
 *  C B C A
 */
var tablePossiblePlay = [][]byte{{'C', 'A', 'B'}, {'A', 'B', 'C'}, {'B', 'C', 'A'}}

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

func part1(input string) int {
	parsedInput := parseInput(input)
	ans := 0
	for _, pair := range parsedInput {
		sum := getValue(pair[1])
		sum += getOutcameValue(pair)
		ans += sum
	}
	return ans
}

func part2(input string) int {
	parsedInput := parseInput(input)
	ans := 0
	for _, pair := range parsedInput {
		letter := whichToWin(pair)
		sum := getValue(letter)
		sum += getOutcameValue([]byte{pair[0], letter})
		ans += sum
	}
	return ans
}

func parseInput(input string) [][]byte {
	ans := [][]byte{}
	for _, line := range strings.Split(input, "\n") {
		abc := line[0]
		xyz := line[2]
		pair := []byte{abc, xyz}
		ans = append(ans, pair)
	}
	return ans
}

func getValue(letter byte) int {
	switch letter {
	case 'X', 'A':
		return 1
	case 'Y', 'B':
		return 2
	case 'Z', 'C':
		return 3
	}
	return 0
}

func getOutcameValue(pair []byte) int {
	return tablePossibleOutcame[getValue(pair[0]-1)][getValue(pair[1]-1)]
}

func whichToWin(pair []byte) byte {
	return tablePossiblePlay[getValue(pair[0]-1)][getValue(pair[1]-1)]
}
