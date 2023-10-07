package main

import (
	"C"
	_ "embed"
	"flag"
	"fmt"
	"strings"
)
import "strconv"

//go:embed input.txt
var input string

type PairPair struct {
	first  IntPair
	second IntPair
}

type IntPair struct {
	first  int
	second int
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

func part1(input string) int {
	parsedInput := parseInput(input)
	ans := 0

	for _, pp := range parsedInput {
		if isFullyOverlapped(pp) {
			ans++
		}
	}

	return ans
}

func part2(input string) int {
	parsedInput := parseInput(input)
	ans := 0

	for _, pp := range parsedInput {
		if isOverlapped(pp) {
			ans++
		}
	}

	return ans
}

func parseInput(input string) []PairPair {
	ans := []PairPair{}
	for _, line := range strings.Split(input, "\n") {
		firstAndSecond := strings.Split(line, ",")
		first := strings.Split(firstAndSecond[0], "-")
		second := strings.Split(firstAndSecond[1], "-")
		ff, err := strconv.Atoi(first[0])
		if err != nil {
			return []PairPair{}
		}
		fs, err := strconv.Atoi(first[1])
		if err != nil {
			return []PairPair{}
		}
		sf, err := strconv.Atoi(second[0])
		if err != nil {
			return []PairPair{}
		}
		ss, err := strconv.Atoi(second[1])
		if err != nil {
			return []PairPair{}
		}
		ans = append(ans, PairPair{
			IntPair{ff, fs},
			IntPair{sf, ss},
		})
	}
	return ans
}

func isFullyOverlapped(pp PairPair) bool {
	return ((pp.first.first >= pp.second.first && pp.first.second <= pp.second.second) ||
		(pp.first.first <= pp.second.first && pp.first.second >= pp.second.second))
}

func isOverlapped(pp PairPair) bool {
	return ((pp.first.second >= pp.second.first && pp.first.first <= pp.second.second) ||
		(pp.second.second >= pp.first.first && pp.first.second >= pp.second.first))
}
