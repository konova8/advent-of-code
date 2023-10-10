package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Instruction struct {
	Quantity int
	From     int
	To       int
}

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
	return foo(input, 4)
}
func part2(input string) int {
	return foo(input, 14)
}

func foo(input string, n int) int {
	buf := ""
	for i := 0; i < n-1; i++ {
		buf += string(input[i])
	}
	for i := n - 1; i < len(input); i++ {
		offset := strings.Index(buf, string(input[i]))
		if offset == -1 {
			m := map[rune]bool{}
			for _, c := range buf {
				m[c] = true
			}
			if len(m) == n-1 {
				return i + 1
			} else {
				buf = buf[1:] + string(input[i])
			}
		} else {
			buf = buf[offset+1:] + input[i:i+offset+1]
			i = i + offset
		}
	}
	return 0
}
