package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type Move struct {
	name string
	val  int
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
	flag.Parse()

	if part == "1" || part == "" {
		fmt.Println("--- Running part 1 ---")
		ansExample := foo(example)
		fmt.Println("Output Example:", ansExample)
		ansInput := foo(input)
		fmt.Println("Output Input:", ansInput)
	}
	if part == "2" || part == "" {
		fmt.Println("--- Running part 2 ---")
		fmt.Println("Output Example:")
		bar(example)
		fmt.Println("Output Input:")
		bar(input)
	}
}

func foo(input string) int {
	XregHistory := []int{}
	Xreg := 1
	for _, line := range strings.Split(input, "\n") {
		if line[0] == 'a' {
			tmp := strings.Split(line, " ")
			v, _ := strconv.Atoi(tmp[1])
			XregHistory = append(XregHistory, Xreg)
			XregHistory = append(XregHistory, Xreg)
			Xreg += v
		} else if line[0] == 'n' {
			XregHistory = append(XregHistory, Xreg)
		} else {
			panic("not addx nor noop")
		}
	}
	ans := 0
	for i := 0; i < (len(XregHistory)+20)/40; i++ {
		ans += XregHistory[20+i*40-1] * (20 + i*40)
	}
	return ans
}

func bar(input string) {
	XregHistory := []int{}
	Xreg := 1
	for _, line := range strings.Split(input, "\n") {
		if line[0] == 'a' {
			tmp := strings.Split(line, " ")
			v, _ := strconv.Atoi(tmp[1])
			XregHistory = append(XregHistory, Xreg)
			XregHistory = append(XregHistory, Xreg)
			Xreg += v
		} else if line[0] == 'n' {
			XregHistory = append(XregHistory, Xreg)
		} else {
			panic("not addx nor noop")
		}
	}
	for i, e := range XregHistory {
		if e-1 <= i%40 && i%40 <= e+1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
		if (i+1)%40 == 0 {
			fmt.Println()
		}
	}
}
