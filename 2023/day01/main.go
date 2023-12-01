package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
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
			ansExample := part2(example2)
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
	first := 0
	second := 0
	for _, line := range strings.Split(str, "\n") {
		for _, r := range line {
			if '0' <= r && r <= '9' {
				first = int(r) - 48
				break
			}
		}
		for j := len(line) - 1; j >= 0; j-- {
			r := line[j]
			if '0' <= r && r <= '9' {
				second = int(r) - 48
				break
			}
		}
		ans += first*10 + second
	}
	return ans
}

var digit = [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func part2(str string) int {
	ans := 0
	for _, line := range strings.Split(str, "\n") {
		first := 0
		first_i := len(line)
		second := 0
		second_i := -1
		for j, d := range digit {
			if n := strings.Index(line, d); n != -1 {
				if first_i > n {
					first_i = n
					first = j
				}
			}
		}
		for j, r := range line {
			if '0' <= r && r <= '9' && j < first_i {
				first = int(r) - 48
				break
			}
		}
		for j, d := range digit {
			if n := strings.LastIndex(line, d); n != -1 {
				if second_i < n {
					second_i = n
					second = j
				}
			}
		}
		for j := len(line) - 1; j >= 0; j-- {
			r := line[j]
			if '0' <= r && r <= '9' && j > second_i {
				second = int(r) - 48
				break
			}
		}
		ans += first*10 + second
	}
	return ans
}
