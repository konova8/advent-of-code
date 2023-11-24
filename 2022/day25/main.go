package main

import (
	_ "embed"
	"flag"
	"fmt"
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
		fmt.Println("It's Christmas, there is no part 2!")
	}
}

func GetDecimal(r byte) int {
	switch r {
	case '2':
		return 2
	case '1':
		return 1
	case '0':
		return 0
	case '-':
		return -1
	case '=':
		return -2
	default:
		return 3
	}
}

func ToInt(str string) int {
	ans := 0
	mul := 1
	for i := len(str) - 1; i >= 0; i-- {
		ans += mul * GetDecimal(str[i])
		mul *= 5
	}
	return ans
}

var baseVal = [5]int{-2, -1, 0, 1, 2}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func ToSNAFU(n int) string {
	ans := ""
	mul := 1
	for n*2 > mul*5 {
		// for n > mul || n*2 > mul*5 {
		mul *= 5
	}
	for mul != 0 {
		val := baseVal[0]
		bestDistance := abs(n + baseVal[0]*mul)
		for _, v := range baseVal[1:] {
			distanceFromZero := abs(n + (v * mul))
			if distanceFromZero < bestDistance {
				bestDistance = distanceFromZero
				val = v
			}
		}
		n = n + (val * mul)
		switch val {
		case -2:
			ans = ans + "2"
		case -1:
			ans = ans + "1"
		case 0:
			ans = ans + "0"
		case 1:
			ans = ans + "-"
		case 2:
			ans = ans + "="
		}
		mul /= 5
	}
	return ans
}

func foo(input string) string {
	ans := 0
	for _, v := range strings.Split(input, "\n") {
		ans += ToInt(v)
	}
	return ToSNAFU(ans)
}
