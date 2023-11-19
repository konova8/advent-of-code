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
			ansExample := foo(example)
			fmt.Println("Output Example:", ansExample)
		}
		if !*noInput {
			ansInput := foo(input)
			fmt.Println("Output Input:", ansInput)
		}
	}
	if part == "2" || part == "" {
		fmt.Println("--- Running part 2 ---")
		if !*noExample {
			ansExample := bar(example)
			fmt.Println("Output Example:", ansExample)
		}
		if !*noInput {
			ansInput := bar(input)
			fmt.Println("Output Input:", ansInput)
		}
	}
}

func foo(input string) int {
	list := map[string]string{}
	for _, line := range strings.Split(input, "\n") {
		res := strings.Split(line, ": ")
		list[res[0]] = res[1]
	}
	var compute func(s string) int
	compute = func(s string) int {
		n, err := strconv.Atoi(list[s])
		if err == nil {
			return n
		}
		// it's an operation
		res := strings.Split(list[s], " ")
		switch res[1] {
		case "+":
			return compute(res[0]) + compute(res[2])
		case "-":
			return compute(res[0]) - compute(res[2])
		case "*":
			return compute(res[0]) * compute(res[2])
		case "/":
			return compute(res[0]) / compute(res[2])
		}
		panic("Something went wrong")
	}
	return compute("root")
}

func bar(input string) string {
	// Change humn line
	index := strings.Index(input, "humn: ")
	var i int
	for i = index; i < len(input); i++ {
		if input[i] == byte('\n') {
			break
		}
	}
	humnLine := input[index:i]
	input = strings.Replace(input, humnLine, "humn: x", 1)
	// Change root line
	index = strings.Index(input, "root: ")
	for i = index; i < len(input); i++ {
		if input[i] == byte('\n') {
			break
		}
	}
	oldRootLine := input[index:i]
	newRootLine := strings.Replace(oldRootLine, "+", "=", 1)
	newRootLine = strings.Replace(newRootLine, "-", "=", 1)
	newRootLine = strings.Replace(newRootLine, "*", "=", 1)
	newRootLine = strings.Replace(newRootLine, "/", "=", 1)
	input = strings.Replace(input, oldRootLine, newRootLine, 1)
	arrInput := strings.Split(input, "\n")
	maxCycle := len(arrInput)
	somethingChanged := true
	oldInput := input
	for i := 0; i < maxCycle && somethingChanged; i++ {
		somethingChanged = false
		for _, line := range arrInput {
			res := strings.Split(line, ": ")
			if res[0] != "root" {
				newInput := strings.ReplaceAll(oldInput, res[0], "("+res[1]+")")
				if newInput != oldInput {
					somethingChanged = true
					oldInput = newInput
				}
			}
		}
	}
	index = strings.Index(oldInput, "root: ")
	for i = index; i < len(oldInput); i++ {
		if oldInput[i] == byte('\n') {
			break
		}
	}
	expression := oldInput[index:i][6:]
	simplifiedExpression := simplify(expression)
	return simplifiedExpression[4:]
}

// Define Regexp
var innerParenthesis = regexp.MustCompile(`\((\d+|x)\)`)                  // (32) -> 32, (x) -> x
var solvableOperation = regexp.MustCompile(`\((\d+ . \d+)\)`)             // (32 + 2) -> = 34
var outerParenthesisLeft = regexp.MustCompile(`^\((.*)\) =`)              // (...) = -> ... =
var outerParenthesisRight = regexp.MustCompile(`= \((.*)\)$`)             // = (...) -> = ...
var operatorLeftLeft = regexp.MustCompile(`^(\d+) (.) \((.*)\) = (.*)$`)  // 3 + (...) = 6 -> ... = (6 - 3)
var operatorLeftRight = regexp.MustCompile(`^\((.*)\) (.) (\d+) = (.*)$`) // (...) + 3 = 5 -> ... = (5 - 3)
var lastOperatorLeftLeft = regexp.MustCompile(`^(\d+) (.) x = (\d+)$`)    // 3 + x = 6 -> x = 6 - 3
var lastOperatorLeftRight = regexp.MustCompile(`^x (.) (\d+) = (\d+)$`)   // x + 3 = 5 -> x = 5 - 3

func simplify(str string) string {
	// We assume that the x is on the left side of the = sign
	for {
		oldStr := str
		// (3) -> 3
		str = innerParenthesis.ReplaceAllString(str, "$1")
		// (3 + 4) -> 7
		tmp := solvableOperation.FindAllString(str, -1)
		for len(tmp) > 0 {
			for _, v := range tmp {
				eval := evaluteBinaryExpression(v)
				str = strings.ReplaceAll(str, v, strconv.Itoa(eval))
			}
			tmp = solvableOperation.FindAllString(str, -1)
		}
		// Outer parenthesis
		str = outerParenthesisLeft.ReplaceAllString(str, "$1 =")
		str = outerParenthesisRight.ReplaceAllString(str, "= $1")
		// Move operator from LL to R
		arrLL := operatorLeftLeft.FindStringSubmatch(str)
		if len(arrLL) != 0 {
			str = arrLL[3] + " = "
			numL, _ := strconv.Atoi(arrLL[1])
			numR, _ := strconv.Atoi(arrLL[4])
			switch arrLL[2] {
			case "+":
				// 3 + (...) = 6 -> ... = (6 - 3)
				str += strconv.Itoa(numR - numL)
			case "-":
				// 3 - (...) = 6 -> ... = (3 - 6)
				str += strconv.Itoa(numL - numR)
			case "*":
				// 3 * (...) = 6 -> ... = (6 / 3)
				str += strconv.Itoa(numR / numL)
			case "/":
				// 3 / (...) = 6 -> ... = (3 / 6)
				str += strconv.Itoa(numL / numR)
			}
		}
		// Move operator from LR to R
		arrLR := operatorLeftRight.FindStringSubmatch(str)
		if len(arrLR) != 0 {
			str = arrLR[1] + " = "
			numL, _ := strconv.Atoi(arrLR[3])
			numR, _ := strconv.Atoi(arrLR[4])
			switch arrLR[2] {
			case "+":
				// (...) + 3 = 5 -> ... = (5 - 3)
				str += strconv.Itoa(numR - numL)
			case "-":
				// (...) - 3 = 5 -> ... = (5 + 3)
				str += strconv.Itoa(numR + numL)
			case "*":
				// (...) * 3 = 5 -> ... = (5 / 3)
				str += strconv.Itoa(numR / numL)
			case "/":
				// (...) / 3 = 5 -> ... = (5 * 3)
				str += strconv.Itoa(numR * numL)
			}
		}
		if oldStr == str {
			break
		}
	}
	// 3 + x = 6 -> x = 6 - 3
	arrLLL := lastOperatorLeftLeft.FindStringSubmatch(str)
	if len(arrLLL) > 0 {
		str = "x = "
		numL, _ := strconv.Atoi(arrLLL[1])
		numR, _ := strconv.Atoi(arrLLL[3])
		switch arrLLL[2] {
		case "+":
			// 3 + (...) = 6 -> ... = (6 - 3)
			str += strconv.Itoa(numR - numL)
		case "-":
			// 3 - (...) = 6 -> ... = (3 - 6)
			str += strconv.Itoa(numL - numR)
		case "*":
			// 3 * (...) = 6 -> ... = (6 / 3)
			str += strconv.Itoa(numR / numL)
		case "/":
			// 3 / (...) = 6 -> ... = (3 / 6)
			str += strconv.Itoa(numL / numR)
		}
	}
	// x + 3 = 5 -> x = 5 - 3
	arrLLR := lastOperatorLeftRight.FindStringSubmatch(str)
	if len(arrLLR) > 0 {
		str = "x = "
		numL, _ := strconv.Atoi(arrLLR[2])
		numR, _ := strconv.Atoi(arrLLR[3])
		switch arrLLR[1] {
		case "+":
			// (...) + 3 = 5 -> ... = (5 - 3)
			str += strconv.Itoa(numR - numL)
		case "-":
			// (...) - 3 = 5 -> ... = (5 + 3)
			str += strconv.Itoa(numR + numL)
		case "*":
			// (...) * 3 = 5 -> ... = (5 / 3)
			str += strconv.Itoa(numR / numL)
		case "/":
			// (...) / 3 = 5 -> ... = (5 * 3)
			str += strconv.Itoa(numR * numL)
		}
	}
	return str
}

func evaluteBinaryExpression(s string) int {
	s = s[1 : len(s)-1]
	r, err := regexp.Compile(`(\d+) (.) (\d+)`)
	if err != nil {
		panic("Regex not valid")
	}
	// arr := r.FindAllStringSubmatch(s, -1)
	arr := r.FindStringSubmatch(s)
	numL, _ := strconv.Atoi(arr[1])
	numR, _ := strconv.Atoi(arr[3])
	switch arr[2] {
	case "+":
		return numL + numR
	case "-":
		return numL - numR
	case "*":
		return numL * numR
	case "/":
		return numL / numR
	}
	return 0
}
