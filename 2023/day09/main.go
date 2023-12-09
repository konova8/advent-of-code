package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/konova8/advent-of-code/util"
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
		example2 = example
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
			clipboard.WriteAll(fmt.Sprint(ansExample))
		}
		if !*noInput {
			ansInput := part1(input)
			fmt.Println("Output Input:", ansInput)
			clipboard.WriteAll(fmt.Sprint(ansInput))
		}
	}
	if part == "2" || part == "" {
		fmt.Println("--- Running part 2 ---")
		if !*noExample {
			ansExample := part2(example2)
			fmt.Println("Output Example:", ansExample)
			clipboard.WriteAll(fmt.Sprint(ansExample))
		}
		if !*noInput {
			ansInput := part2(input)
			fmt.Println("Output Input:", ansInput)
			clipboard.WriteAll(fmt.Sprint(ansInput))
		}
	}
	p := util.Position{}
	_ = p
}

var r = regexp.MustCompile(`-?\d+`)

func findNextNumber(history []int) int {
	allZero := true
	for _, v := range history {
		if v != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		return 0
	}
	newHistory := []int{}
	for i := 0; i < len(history)-1; i++ {
		newHistory = append(newHistory, history[i+1]-history[i])
	}
	return findNextNumber(newHistory) + history[len(history)-1]
}

func part1(str string) int {
	ans := 0
	for _, line := range strings.Split(str, "\n") {
		tmp := r.FindAllStringSubmatch(line, -1)
		history := []int{}
		for _, v := range tmp {
			n, _ := strconv.Atoi(v[0])
			history = append(history, n)
		}
		ans += findNextNumber(history)
	}
	return ans
}

func findPrevNumber(history []int) int {
	allZero := true
	for _, v := range history {
		if v != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		return 0
	}
	newHistory := []int{}
	for i := 0; i < len(history)-1; i++ {
		newHistory = append(newHistory, history[i+1]-history[i])
	}
	return history[0] - findPrevNumber(newHistory)
}

func part2(str string) int {
	ans := 0
	for _, line := range strings.Split(str, "\n") {
		tmp := r.FindAllStringSubmatch(line, -1)
		history := []int{}
		for _, v := range tmp {
			n, _ := strconv.Atoi(v[0])
			history = append(history, n)
		}
		ans += findPrevNumber(history)
	}
	return ans
}
