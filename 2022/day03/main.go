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

	for _, line := range parsedInput {
		pair := []string{line[0 : len(line)/2], line[len(line)/2:]}
		commonLetter := findFirstCommonLetter(pair[0], pair[1])
		ans += getIntValue(commonLetter)
	}

	return ans
}

func part2(input string) int {
	parsedInput := parseInput(input)
	ans := 0

	for i := 0; i < len(parsedInput); i += 3 {
		first, second, third := parsedInput[i], parsedInput[i+1], parsedInput[i+2]
		firsec := stringIntersection(first, second)
		commonLetter := findFirstCommonLetter(firsec, third)
		ans += getIntValue(commonLetter)
	}

	return ans
}

func parseInput(input string) []string {
	ans := []string{}
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, line)
	}
	return ans
}

func getIntValue(letter byte) int {
	ans := int(letter)
	if letter <= 'z' && letter >= 'a' {
		ans -= 96
	} else if letter <= 'Z' && letter >= 'A' {
		ans -= 38
	}
	return ans
}

func findFirstCommonLetter(s1, s2 string) byte {
	hashMap := make(map[rune]bool)
	for _, r := range s1 {
		hashMap[r] = true
	}
	for _, r := range s2 {
		if hashMap[r] {
			return byte(r)
		}
	}
	return 'a'
}

func stringIntersection(s1, s2 string) string {
	hashMap := make(map[rune]bool)
	res := ""
	for _, r := range s1 {
		hashMap[r] = true
	}
	for _, r := range s2 {
		if hashMap[r] {
			res += string(r)
		}
	}
	return res
}
