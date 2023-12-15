package main

import (
	"advent-of-code/util"
	_ "embed"
	"flag"
	"fmt"
	"slices"
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

func hash(s string) int {
	ret := 0
	for _, r := range s {
		ret = ((ret + int(r)) * 17) % 256
	}
	return ret
}

func part1(str string) int {
	ans := 0
	for _, s := range strings.Split(str, ",") {
		ans += hash(s)
	}
	return ans
}

type Box struct {
	label string
	num   int
}

func part2(str string) int {
	lensSlots := [256][]Box{}
	for _, s := range strings.Split(str, ",") {
		i_equal := strings.IndexByte(s, '=')
		box := Box{}
		if i_equal == -1 {
			box.label = s[:len(s)-1]
			box.num = -1
		} else {
			box.label = s[:i_equal]
			box.num, _ = strconv.Atoi(s[i_equal+1:])
		}
		i := hash(box.label)
		if box.num == -1 {
			lensSlots[i] = slices.DeleteFunc(lensSlots[i], func(b Box) bool {
				return b.label == box.label
			})
		} else {
			j := slices.IndexFunc(lensSlots[i], func(b Box) bool {
				return b.label == box.label
			})
			if j == -1 {
				lensSlots[i] = append(lensSlots[i], box)
			} else {
				lensSlots[i][j].num = box.num
			}
		}
	}
	ans := 0
	for i, ls := range lensSlots {
		for j, box := range ls {
			ans += (i + 1) * (j + 1) * (box.num)
		}
	}
	return ans
}
