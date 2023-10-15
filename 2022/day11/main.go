package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Operation struct {
	operator string
	second   string
}

type Test struct {
	number  int
	ifTrue  int
	ifFalse int
}

type Monkey struct {
	items          []int
	operation      Operation
	test           Test
	countInspected int
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
	monkeys := parseInput(input)
	for i := 1; i <= 20; i++ {
		for j := 0; j < len(monkeys); j++ {
			for _, e := range monkeys[j].items {
				monkeys[j].countInspected++
				e = doOperation(e, monkeys[j].operation)
				e = e / 3
				monkeys[j].items = monkeys[j].items[1:]
				if e%monkeys[j].test.number == 0 {
					monkeys[monkeys[j].test.ifTrue].items = append(monkeys[monkeys[j].test.ifTrue].items, e)
				} else {
					monkeys[monkeys[j].test.ifFalse].items = append(monkeys[monkeys[j].test.ifFalse].items, e)
				}
			}
		}
	}
	count := []int{}
	for _, m := range monkeys {
		count = append(count, m.countInspected)
	}
	n1 := slices.Max(count)
	i_n1 := slices.Index(count, n1)
	count = slices.Delete(count, i_n1, i_n1+1)
	n2 := slices.Max(count)
	return n1 * n2
}

func bar(input string) int {
	monkeys := parseInput(input)
	mcm := 1
	for _, m := range monkeys {
		mcm *= m.test.number
	}
	for i := 1; i <= 10000; i++ {
		for j := 0; j < len(monkeys); j++ {
			for _, e := range monkeys[j].items {
				monkeys[j].countInspected++
				e = doOperation(e, monkeys[j].operation) % mcm
				monkeys[j].items = monkeys[j].items[1:]
				if e%monkeys[j].test.number == 0 {
					monkeys[monkeys[j].test.ifTrue].items = append(monkeys[monkeys[j].test.ifTrue].items, e)
				} else {
					monkeys[monkeys[j].test.ifFalse].items = append(monkeys[monkeys[j].test.ifFalse].items, e)
				}
			}
		}
	}
	count := []int{}
	for _, m := range monkeys {
		count = append(count, m.countInspected)
	}
	n1 := slices.Max(count)
	i_n1 := slices.Index(count, n1)
	count = slices.Delete(count, i_n1, i_n1+1)
	n2 := slices.Max(count)
	return n1 * n2
}

func parseInput(input string) []Monkey {
	monkeyRet := []Monkey{}
	for _, m := range strings.Split(input, "\n\n") {
		monkey := Monkey{}
		monkeyArr := strings.Split(m, "\n")
		// Starting items
		eArr := strings.Split((strings.Split(monkeyArr[1], ": "))[1], ", ")
		arrItems := []int{}
		for _, e := range eArr {
			n, _ := strconv.Atoi(e)
			arrItems = append(arrItems, n)
		}
		monkey.items = arrItems
		// Operation
		arrOp := strings.Split(strings.Split(monkeyArr[2], "= ")[1], " ")
		monkey.operation = Operation{arrOp[1], arrOp[2]}
		// Test
		num, _ := strconv.Atoi(strings.Split(monkeyArr[3], "by ")[1])
		monkey.test.number = num
		num, _ = strconv.Atoi(strings.Split(monkeyArr[4], "monkey ")[1])
		monkey.test.ifTrue = num
		num, _ = strconv.Atoi(strings.Split(monkeyArr[5], "monkey ")[1])
		monkey.test.ifFalse = num
		monkeyRet = append(monkeyRet, monkey)
	}
	return monkeyRet
}

func doOperation(e int, o Operation) int {
	switch o.operator {
	case "+":
		v, err := strconv.Atoi(o.second)
		if err != nil {
			return e + e
		} else {
			return e + v
		}
	case "-":
		v, err := strconv.Atoi(o.second)
		if err != nil {
			return e - e
		} else {
			return e - v
		}
	case "*":
		v, err := strconv.Atoi(o.second)
		if err != nil {
			return e * e
		} else {
			return e * v
		}
	case "/":
		v, err := strconv.Atoi(o.second)
		if err != nil {
			return e / e
		} else {
			return e / v
		}
	}
	return -1
}
