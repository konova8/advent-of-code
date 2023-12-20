package main

import (
	"advent-of-code/util"
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/atotto/clipboard"
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

var r = regexp.MustCompile(`(?m)(.*) (.*)$`)

func part1(str string) int {
	ans := 0
	tmp := r.FindAllStringSubmatch(str, -1)
	for _, v := range tmp {
		line := v[1]
		numbersStr := strings.Split(v[2], ",")
		numbers := []int{}
		for _, n := range numbersStr {
			m, _ := strconv.Atoi(n)
			numbers = append(numbers, m)
		}
		for ck := range cache {
			delete(cache, ck)
		}
		n := countAllString(line, numbers, Situation{0, 0, 0})
		ans += n
	}
	return ans
}

type Situation struct {
	index_line    int
	index_numbers int
	current       int
}

var cache = map[Situation]int{}

func countAllString(line string, numbers []int, situa Situation) int {
	if v, ok := cache[situa]; ok {
		return v
	}
	// Base cases
	if situa.index_line == len(line) {
		if situa.index_numbers == len(numbers) && situa.current == 0 {
			return 1
		} else if situa.index_numbers == len(numbers)-1 && numbers[situa.index_numbers] == situa.current {
			return 1
		} else {
			return 0
		}
	}
	// Recursion
	ans := 0
	if line[situa.index_line] == '.' || line[situa.index_line] == '?' {
		if situa.current == 0 {
			newSitua := Situation{
				index_line:    situa.index_line + 1,
				index_numbers: situa.index_numbers,
				current:       0,
			}
			ans += countAllString(line, numbers, newSitua)
		} else if situa.index_numbers < len(numbers) && numbers[situa.index_numbers] == situa.current {
			newSitua := Situation{
				index_line:    situa.index_line + 1,
				index_numbers: situa.index_numbers + 1,
				current:       0,
			}
			ans += countAllString(line, numbers, newSitua)
		}
	}
	if line[situa.index_line] == '#' || line[situa.index_line] == '?' {
		newSitua := Situation{
			index_line:    situa.index_line + 1,
			index_numbers: situa.index_numbers,
			current:       situa.current + 1,
		}
		ans += countAllString(line, numbers, newSitua)
	}
	cache[situa] = ans
	return ans
}

func part2(str string) int {
	ans := 0
	tmp := r.FindAllStringSubmatch(str, -1)
	for _, v := range tmp {
		line := v[1]
		oldNumbersStr := v[2]
		for i := 0; i < 4; i++ {
			oldNumbersStr += "," + v[2]
			line += "?" + v[1]
		}
		numbersStr := strings.Split(oldNumbersStr, ",")
		numbers := []int{}
		for _, n := range numbersStr {
			m, _ := strconv.Atoi(n)
			numbers = append(numbers, m)
		}
		for ck := range cache {
			delete(cache, ck)
		}
		n := countAllString(line, numbers, Situation{0, 0, 0})
		ans += n
	}
	return ans
}
