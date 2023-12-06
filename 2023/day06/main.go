package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
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
			ansExample := part2(example)
			fmt.Println("Output Example:", ansExample)
			clipboard.WriteAll(fmt.Sprint(ansExample))
		}
		if !*noInput {
			ansInput := part2(input)
			fmt.Println("Output Input:", ansInput)
			clipboard.WriteAll(fmt.Sprint(ansInput))
		}
	}
}

var r = regexp.MustCompile(` \d+`)

type Pair struct {
	time int
	dist int
}

func part1(str string) int {
	lines := strings.Split(str, "\n")
	pairs := []Pair{}
	// Time
	times := r.FindAllString(lines[0], -1)
	for _, v := range times {
		v = strings.TrimSpace(v)
		t, _ := strconv.Atoi(v)
		pairs = append(pairs, Pair{time: t})
	}
	// Distance
	dists := r.FindAllString(lines[1], -1)
	for i, v := range dists {
		v = strings.TrimSpace(v)
		d, _ := strconv.Atoi(v)
		pairs[i].dist = d
	}
	// Compute Pairs
	ans := 1
	for _, p := range pairs {
		t := float64(p.time)
		d := float64(p.dist)
		delta := math.Sqrt(math.Pow(t, 2) - 4*d)
		upperBound := int(math.Ceil(((t + delta) / 2) - 1))
		lowerBound := int(math.Floor(((t - delta) / 2) + 1))
		ans *= upperBound - lowerBound + 1
	}
	return ans
}

func part2(str string) int {
	lines := strings.Split(str, "\n")
	p := Pair{}
	// Time
	timeNum := strings.Split(lines[0], ":")
	timeStr := strings.Replace(timeNum[1], " ", "", -1)
	time, _ := strconv.Atoi(timeStr)
	p.time = time
	// Distance
	distNum := strings.Split(lines[1], ":")
	distStr := strings.Replace(distNum[1], " ", "", -1)
	dist, _ := strconv.Atoi(distStr)
	p.dist = dist
	// Compute Pairs
	t := float64(p.time)
	d := float64(p.dist)
	delta := math.Sqrt(math.Pow(t, 2) - 4*d)
	upperBound := int(math.Ceil(((t + delta) / 2) - 1))
	lowerBound := int(math.Floor(((t - delta) / 2) + 1))
	return upperBound - lowerBound + 1
}
