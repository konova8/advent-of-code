package main

import (
	"advent-of-code/util"
	_ "embed"
	"flag"
	"fmt"
	"github.com/atotto/clipboard"
	"strings"
	"time"
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
	p := util.Position{}
	_ = p
}

func main() {
	var part string
	flag.StringVar(&part, "part", "", "part 1 or 2")
	noExample := flag.Bool("no-example", false, "you don't want to check example")
	noInput := flag.Bool("no-input", false, "you don't want to check input")
	noE := flag.Bool("noe", false, "you don't want to check example")
	noI := flag.Bool("noi", false, "you don't want to check input")
	flag.Parse()

	if part == "1" || part == "" {
		fmt.Println("--- Running part 1 ---")
		if !*noExample && !*noE {
			s := time.Now()
			ansExample := fmt.Sprint(part1(example))
			fmt.Printf("Output Example: %s\n", ansExample)
			fmt.Printf("Computed in %v\n", time.Now().Sub(s))
			clipboard.WriteAll(ansExample)
		}
		if !*noInput && !*noI {
			s := time.Now()
			ansInput := fmt.Sprint(part1(input))
			fmt.Printf("Output Input: %s\n", ansInput)
			fmt.Printf("Computed in %v\n", time.Now().Sub(s))
			clipboard.WriteAll(ansInput)
		}
	}
	if part == "2" || part == "" {
		fmt.Println("--- Running part 2 ---")
		if !*noExample && !*noE {
			s := time.Now()
			ansExample := fmt.Sprint(part2(example))
			fmt.Printf("Output Example: %s\n", ansExample)
			fmt.Printf("Computed in %v\n", time.Now().Sub(s))
			clipboard.WriteAll(ansExample)
		}
		if !*noInput && !*noI {
			s := time.Now()
			ansInput := fmt.Sprint(part2(input))
			fmt.Printf("Output Input: %s\n", ansInput)
			fmt.Printf("Computed in %v\n", time.Now().Sub(s))
			clipboard.WriteAll(ansInput)
		}
	}
}

func part1(str string) int {
	ans := 0
	gridsStr := strings.Split(str, "\n\n")
	for _, gridStr := range gridsStr {
		lines := strings.Split(gridStr, "\n")
		// Check for vertical mirror
		width := len(lines[0])
		for i := 0; i <= width-2; i++ {
			badColumn := false
			for _, line := range lines {
				lineMirrored := true
				for j := 0; i-j >= 0 && i+1+j <= len(line)-1 && lineMirrored; j++ {
					if line[i-j] != line[i+1+j] {
						lineMirrored = false
					}
				}
				if !lineMirrored {
					badColumn = true
					break
				}
			}
			if badColumn {
				continue
			} else {
				ans += i + 1
			}
		}
		// Check for horizontal mirror
		height := len(lines)
		for i := 0; i <= height-2; i++ {
			badRow := false
			for j := 0; i-j >= 0 && i+1+j <= len(lines)-1 && !badRow; j++ {
				if lines[i-j] != lines[i+1+j] {
					badRow = true
				}
			}
			if badRow {
				continue
			} else {
				ans += 100 * (i + 1)
			}
		}
	}
	return ans
}

func part2(str string) int {
	ans := 0
	gridsStr := strings.Split(str, "\n\n")
	for _, gridStr := range gridsStr {
		lines := strings.Split(gridStr, "\n")
		// Check for vertical mirror
		width := len(lines[0])
		for i := 0; i <= width-2; i++ {
			badColumn := false
			smudge := 0
			for _, line := range lines {
				smudgeLine := 0
				for j := 0; i-j >= 0 && i+1+j <= len(line)-1 && smudgeLine <= 1; j++ {
					if line[i-j] != line[i+1+j] {
						smudgeLine++
					}
				}
				smudge += smudgeLine
				if smudge > 1 {
					badColumn = true
					break
				}
			}
			if badColumn || smudge != 1 {
				continue
			} else {
				ans += i + 1
				break
			}
		}
		// Check for horizontal mirror
		height := len(lines)
		for i := 0; i <= height-2; i++ {
			smudge := 0
			badRow := false
			for j := 0; i-j >= 0 && i+1+j <= len(lines)-1 && !badRow; j++ {
				smudgeRow := 0
				l1 := lines[i-j]
				l2 := lines[i+1+j]
				for k := 0; k <= len(l1)-1 && smudgeRow <= 1; k++ {
					if l1[k] != l2[k] {
						smudgeRow++
					}
				}
				if smudgeRow > 1 {
					badRow = true
					break
				} else {
					smudge += smudgeRow
				}
			}
			if badRow || smudge != 1 {
				continue
			} else {
				ans += 100 * (i + 1)
				break
			}
		}
	}
	return ans
}
