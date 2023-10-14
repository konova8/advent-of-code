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
import "sort"

//go:embed input.txt
var input string

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
	} else {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	parsedInput := parseInput(input)
	ans := 0
	for _, elf := range parsedInput {
		sum := 0
		for _, item := range elf {
			sum += item
		}
		if sum > ans {
			ans = sum
		}
	}

	return ans
}

func part2(input string) int {
	parsedInput := parseInput(input)
	ans := []int{0, 0, 0}
	for _, elf := range parsedInput {
		sum := 0
		for _, item := range elf {
			sum += item
		}
		if sum > slices.Min(ans) {
			ans = append(ans, sum)
			sort.Slice(ans, func(i, j int) bool {
				return ans[i] < ans[j]
			})
			ans = ans[1:]
		}
	}

	return ans[0] + ans[1] + ans[2]
}

func parseInput(input string) (ans [][]int) {
	// fmt.Println(input)
	ans = append(ans, []int{})
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			ans = append(ans, []int{})
		} else {
			num, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println(err)
				return
			}
			ans[len(ans)-1] = append(ans[len(ans)-1], num)
		}
	}
	// fmt.Println(ans)
	return ans
}
