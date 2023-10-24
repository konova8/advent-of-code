package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"slices"
	"strings"
)

type Pair struct {
	first        []interface{}
	firstString  string
	second       []interface{}
	secondString string
}

type Packet struct {
	elem       []interface{}
	elemString string
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
		ansExample := bar(example)
		fmt.Println("Output Example:", ansExample)
		ansInput := bar(input)
		fmt.Println("Output Input:", ansInput)
	}
}

func foo(input string) int {
	ans := 0
	pairs := parseInput(input)
	for i, p := range pairs {
		if Compare(p.first, p.second) == -1 {
			ans += i + 1
		}
	}
	return ans
}

func bar(input string) int {
	ans := 1
	pairs := parseInput(input)
	packets := []Packet{{
		elem:       getRawString("[[2]]"),
		elemString: "[[2]]",
	}, {
		elem:       getRawString("[[6]]"),
		elemString: "[[6]]",
	}}

	for _, p := range pairs {
		packets = append(packets, Packet{
			elem:       p.first,
			elemString: p.firstString,
		})
		packets = append(packets, Packet{
			elem:       p.second,
			elemString: p.secondString,
		})
	}

	slices.SortFunc(packets, func(a, b Packet) int {
		return Compare(a.elem, b.elem)
	})

	for i, p := range packets {
		if p.elemString == "[[2]]" {
			ans *= i + 1
		} else if p.elemString == "[[6]]" {
			ans *= i + 1
		}
	}

	return ans
}

func parseInput(input string) []Pair {
	ret := []Pair{}
	for _, twolines := range strings.Split(input, "\n\n") {
		newPair := Pair{}
		tmp := strings.Split(twolines, "\n")
		firstLine := tmp[0]
		secondLine := tmp[1]
		newPair.first = getRawString(firstLine)
		newPair.firstString = firstLine
		newPair.second = getRawString(secondLine)
		newPair.secondString = secondLine
		ret = append(ret, newPair)
	}
	return ret
}

func getRawString(line string) []interface{} {
	ans := []interface{}{}
	json.Unmarshal([]byte(line), &ans)
	return ans
}

func Compare(left interface{}, right interface{}) int {
	leftNum, isLeftNum := left.(float64)
	rightNum, isRightNum := right.(float64)

	leftList, isLeftList := left.([]interface{})
	rightList, isRightList := right.([]interface{})

	if isLeftNum && isRightNum {
		if leftNum < rightNum {
			return -1
		} else if leftNum > rightNum {
			return 1
		} else if leftNum == rightNum {
			return 0
		}
	} else if isLeftList && isRightList {
		for i := 0; ; i++ {
			if i == len(leftList) && i == len(rightList) {
				// Same length
				return 0
			} else if i == len(leftList) && i != len(rightList) {
				// Left without more elements
				return -1
			} else if i != len(leftList) && i == len(rightList) {
				// Right without more elements
				return 1
			} else {
				// Both have still elements
				a := Compare(leftList[i], rightList[i])
				if a != 0 {
					return a
				}
			}
		}
	} else if isLeftNum && isRightList {
		a := Compare([]interface{}{leftNum}, rightList)
		if a != 0 {
			return a
		}
	} else if isLeftList && isRightNum {
		a := Compare(leftList, []interface{}{rightNum})
		if a != 0 {
			return a
		}
	}
	return 0
}
