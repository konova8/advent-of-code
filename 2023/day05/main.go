package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
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

func part1(str string) int {
	groups := strings.Split(str, "\n\n")
	// Seeds
	seeds := []int{}
	tmp := strings.Split(groups[0], ": ")
	for _, v := range strings.Split(tmp[1], " ") {
		n, _ := strconv.Atoi(v)
		seeds = append(seeds, n)
	}
	// Maps
	for _, g := range groups[1:] {
		rows := strings.Split(g, "\n")[1:]
		// Generate Map
		m := [][3]int{}
		for _, r := range rows {
			tmp := strings.Split(r, " ")
			triplet := [3]int{}
			triplet[0], _ = strconv.Atoi(tmp[0])
			triplet[1], _ = strconv.Atoi(tmp[1])
			triplet[2], _ = strconv.Atoi(tmp[2])
			m = append(m, triplet)
		}
		// Update main array
		for i, v := range seeds {
			for _, t := range m {
				if v >= t[1] && v < t[1]+t[2] {
					seeds[i] = t[0] + (v - t[1])
				}
			}
		}
	}
	return slices.Min(seeds)
}

type Range struct {
	start  int
	end    int
	length int
}

type Triplet struct {
	destination int
	source      int
	length      int
}

func (r Range) toString() string {
	s := "["
	for x := r.start; x <= r.end; x++ {
		s += fmt.Sprintf("%d, ", x)
	}
	s = s[:len(s)-2] + "]"
	return s
}

func (t Triplet) toString() string {
	from := Range{t.source, t.source + t.length - 1, t.length}
	to := Range{t.destination, t.destination + t.length - 1, t.length}
	return from.toString() + " -> " + to.toString()
}

func generateNewRanges(interval Range, triplets []Triplet) []Range {
	// Order triplets. This is not necessary, I think
	slices.SortFunc(triplets, func(a, b Triplet) int { return a.source - b.source })
	needToProcess := []Range{interval}
	partOfNewRange := []Range{}
	for len(needToProcess) > 0 {
		r := needToProcess[0]
		needToProcess = needToProcess[1:]
		modified := false
		for i, t := range triplets {
			middleStart := max(t.source, r.start)
			middleEnd := min(r.end, t.source+t.length-1)
			if middleEnd-middleStart+1 > 0 {
				// Middle part (the one that changes)
				delta := t.destination - t.source
				middleRange := Range{
					start: middleStart + delta,
					end:   middleEnd + delta,
				}
				middleRange.length = middleRange.end - middleRange.start + 1
				modified = true
				partOfNewRange = append(partOfNewRange, middleRange)
				// Left part
				leftRange := Range{
					start: r.start,
					end:   min(middleStart-1, r.end),
				}
				leftRange.length = leftRange.end - leftRange.start + 1
				// fmt.Println("left: ", leftRange)
				if leftRange.length > 0 {
					needToProcess = append(needToProcess, leftRange)
				}
				// Right part
				rightRange := Range{
					start: max(middleEnd+1, r.start),
					end:   r.end,
				}
				rightRange.length = rightRange.end - rightRange.start + 1
				// fmt.Println("right: ", rightRange)
				if rightRange.length > 0 {
					needToProcess = append(needToProcess, rightRange)
				}
			} else if i == len(triplets)-1 && !modified {
				partOfNewRange = append(partOfNewRange, r)
			}
		}
	}
	ret := []Range{}
	ret = append(ret, partOfNewRange...)
	return ret
}

func part2(str string) int {
	groups := strings.Split(str, "\n\n")
	// Seeds
	seedRanges := []Range{}
	tmp := strings.Split(groups[0], ": ")
	arrSeeds := strings.Split(tmp[1], " ")
	for i := 0; i < len(arrSeeds); i += 2 {
		start, _ := strconv.Atoi(arrSeeds[i])
		length, _ := strconv.Atoi(arrSeeds[i+1])
		seedRanges = append(seedRanges, Range{
			start:  start,
			end:    start + length - 1,
			length: length,
		})
	}
	// Maps
	for _, g := range groups[1:] {
		rows := strings.Split(g, "\n")[1:]
		// Generate Map
		triplets := []Triplet{}
		for _, r := range rows {
			tmp := strings.Split(r, " ")
			triplet := Triplet{}
			triplet.destination, _ = strconv.Atoi(tmp[0])
			triplet.source, _ = strconv.Atoi(tmp[1])
			triplet.length, _ = strconv.Atoi(tmp[2])
			triplets = append(triplets, triplet)
		}
		// Get new ranges
		newSeedRanges := []Range{}
		for _, r := range seedRanges {
			newRanges := generateNewRanges(r, triplets)
			newSeedRanges = append(newSeedRanges, newRanges...)
		}
		seedRanges = newSeedRanges
	}
	return slices.MinFunc(seedRanges, func(a, b Range) int { return a.start - b.start }).start
}
