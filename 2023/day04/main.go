package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"strconv"
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
			ansExample := part1(example)
			fmt.Println("Output Example:", ansExample)
		}
		if !*noInput {
			ansInput := part1(input)
			fmt.Println("Output Input:", ansInput)
		}
	}
	if part == "2" || part == "" {
		fmt.Println("--- Running part 2 ---")
		if !*noExample {
			ansExample := part2(example)
			fmt.Println("Output Example:", ansExample)
		}
		if !*noInput {
			ansInput := part2(input)
			fmt.Println("Output Input:", ansInput)
		}
	}
}

func part1(str string) int {
	ans := 0
	for _, line := range strings.Split(str, "\n") {
		points := 1
		tmp1 := strings.Split(line, ": ")
		tmp2 := strings.Split(tmp1[1], " | ")
		winningNumbers := []int{}
		for _, n := range strings.Split(tmp2[0], " ") {
			if n != "" {
				num, err := strconv.Atoi(n)
				if err != nil {
					s := fmt.Sprintf("`%s` error with atoi", n)
					panic(s)
				}
				winningNumbers = append(winningNumbers, num)
			}
		}
		for _, n := range strings.Split(tmp2[1], " ") {
			if n != "" {
				num, err := strconv.Atoi(n)
				if err != nil {
					s := fmt.Sprintf("`%s` error with atoi", n)
					panic(s)
				}
				if slices.Contains(winningNumbers, num) {
					points *= 2
				}
			}
		}
		ans += (points / 2)
	}
	return ans
}

func part2(str string) int {
	ans := 0
	numCards := []int{}
	lines := strings.Split(str, "\n")
	for i := 0; i < len(lines); i++ {
		numCards = append(numCards, 1)
	}
	for i, line := range lines {
		points := 0
		tmp1 := strings.Split(line, ": ")
		tmp2 := strings.Split(tmp1[1], " | ")
		winningNumbers := []int{}
		for _, n := range strings.Split(tmp2[0], " ") {
			if n != "" {
				num, err := strconv.Atoi(n)
				if err != nil {
					s := fmt.Sprintf("`%s` error with atoi", n)
					panic(s)
				}
				winningNumbers = append(winningNumbers, num)
			}
		}
		for _, n := range strings.Split(tmp2[1], " ") {
			if n != "" {
				num, err := strconv.Atoi(n)
				if err != nil {
					s := fmt.Sprintf("`%s` error with atoi", n)
					panic(s)
				}
				if slices.Contains(winningNumbers, num) {
					points++
				}
			}
		}
		for j := i + 1; j < i+1+points && j < len(numCards); j++ {
			numCards[j] += numCards[i]
		}
	}
	for _, v := range numCards {
		ans += v
	}
	return ans
}
