package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strings"

	"advent-of-code/util"
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
		panic("empty example2.txt file")
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

var r = regexp.MustCompile(`(.*) = \((.*), (.*)\)`)

func part1(str string) int {
	tmp := strings.Split(str, "\n\n")
	instructions := tmp[0]
	lines := r.FindAllStringSubmatch(tmp[1], -1)
	nodes := map[string][2]string{}
	for _, v := range lines {
		nodes[v[1]] = [2]string{v[2], v[3]}
	}
	nowNode := "AAA"
	ans := 0
	for i := 0; nowNode != "ZZZ"; i = (i + 1) % len(instructions) {
		if instructions[i] == 'R' {
			nowNode = nodes[nowNode][1]
		} else {
			nowNode = nodes[nowNode][0]
		}
		ans++
	}
	return ans
}

func part2(str string) int {
	tmp := strings.Split(str, "\n\n")
	instructions := tmp[0]
	lines := r.FindAllStringSubmatch(tmp[1], -1)
	nodes := map[string][2]string{}
	startingNodes := []string{}
	pointsNodes := []int{}
	for _, v := range lines {
		nodes[v[1]] = [2]string{v[2], v[3]}
		if v[1][2] == 'A' {
			startingNodes = append(startingNodes, v[1])
			pointsNodes = append(pointsNodes, 0)
		}
	}
	for i, v := range startingNodes {
		nowNode := v
		count := 0
		for i := 0; nowNode[2] != 'Z'; i = (i + 1) % len(instructions) {
			if instructions[i] == 'R' {
				nowNode = nodes[nowNode][1]
			} else {
				nowNode = nodes[nowNode][0]
			}
			count++
		}
		pointsNodes[i] = count
	}
	return util.LCM(pointsNodes...)
}
