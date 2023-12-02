package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
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

type Hand struct {
	blue  int
	red   int
	green int
}

func part1(str string) int {
	r, err := regexp.Compile("([0-9]+) (blue|red|green)")
	if err != nil {
		panic("regex compile error")
	}
	ans := 0
	for i, line := range strings.Split(str, "\n") {
		ans += i + 1
		tmp := r.FindAllStringSubmatch(line, -1)
		for _, t := range tmp {
			n, err := strconv.Atoi(t[1])
			if err != nil {
				panic("strconv atoi compile error")
			}
			switch t[2] {
			case "red":
				if n > 12 {
					ans -= i + 1
					goto BREAK
				}
			case "green":
				if n > 13 {
					ans -= i + 1
					goto BREAK
				}
			case "blue":
				if n > 14 {
					ans -= i + 1
					goto BREAK
				}
			}
		}
	BREAK:
	}
	return ans
}

func part2(str string) int {
	r, err := regexp.Compile("([0-9]+) (blue|red|green)")
	if err != nil {
		panic("regex compile error")
	}
	ans := 0
	for _, line := range strings.Split(str, "\n") {
		minHand := Hand{}
		tmp := r.FindAllStringSubmatch(line, -1)
		for _, t := range tmp {
			n, err := strconv.Atoi(t[1])
			if err != nil {
				panic("strconv atoi compile error")
			}
			switch t[2] {
			case "red":
				minHand.red = max(minHand.red, n)
			case "green":
				minHand.green = max(minHand.green, n)
			case "blue":
				minHand.blue = max(minHand.blue, n)
			}
		}
		ans += minHand.red * minHand.green * minHand.blue
	}
	return ans
}
