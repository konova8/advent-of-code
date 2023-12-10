package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"advent-of-code/util"
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
	p := util.Position{}
	_ = p
}

var valueCard = map[rune]int{
	'2': 0,
	'3': 1,
	'4': 2,
	'5': 3,
	'6': 4,
	'7': 5,
	'8': 6,
	'9': 7,
	'T': 8,
	'J': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}

type Hand map[rune]int

func convertHand(handStr string) Hand {
	hand := Hand{}
	for _, r := range handStr {
		v, found := hand[r]
		if found {
			hand[r] = v + 1
		} else {
			hand[r] = 1
		}
	}
	return hand
}

const (
	High = iota
	One
	Two
	Three
	Full
	Four
	Five
)

func computeHand(hand Hand) int {
	if len(hand) == 1 {
		return Five
	}
	if len(hand) == 2 {
		for _, v := range hand {
			if v == 4 {
				return Four
			}
		}
		return Full
	}
	if len(hand) == 3 {
		for _, v := range hand {
			if v == 3 {
				return Three
			}
		}
		return Two
	}
	if len(hand) == 4 {
		return One
	}
	if len(hand) == 5 {
		return High
	}
	return -1
}

func CompareHand(h1, h2 HandWithBid) int {
	n1 := computeHand(h1.hand)
	n2 := computeHand(h2.hand)
	if n1 != n2 {
		return n1 - n2
	}
	for i, r1 := range h1.str {
		for j, r2 := range h2.str {
			if i == j {
				v1 := valueCard[r1]
				v2 := valueCard[r2]
				if v1 != v2 {
					return v1 - v2
				}
			}
		}
	}
	return 0
}

type HandWithBid struct {
	str  string
	hand Hand
	bid  int
}

func part1(str string) int {
	valueCard['J'] = 9
	pairs := []HandWithBid{}
	for _, v := range strings.Split(str, "\n") {
		tmp := strings.Split(v, " ")
		b, _ := strconv.Atoi(tmp[1])
		p := HandWithBid{
			str:  tmp[0],
			hand: convertHand(tmp[0]),
			bid:  b,
		}
		pairs = append(pairs, p)
	}
	slices.SortFunc(pairs, CompareHand)
	ans := 0
	for i, hwb := range pairs {
		ans += (i + 1) * hwb.bid
	}
	return ans
}

func convertHand2(handStr string) Hand {
	hand := Hand{}
	jollyCount := 0
	for _, r := range handStr {
		if r == 'J' {
			jollyCount++
			continue
		}
		v, found := hand[r]
		if found {
			hand[r] = v + 1
		} else {
			hand[r] = 1
		}
	}
	maxK := ' '
	maxV := 0
	for k, v := range hand {
		if v > maxV {
			maxK = k
			maxV = v
		}
	}
	for i := 0; i < jollyCount; i++ {
		hand[maxK] += 1
	}
	return hand
}

func part2(str string) int {
	valueCard['J'] = -1
	pairs := []HandWithBid{}
	for _, v := range strings.Split(str, "\n") {
		tmp := strings.Split(v, " ")
		b, _ := strconv.Atoi(tmp[1])
		p := HandWithBid{
			str:  tmp[0],
			hand: convertHand2(tmp[0]),
			bid:  b,
		}
		pairs = append(pairs, p)
	}
	slices.SortFunc(pairs, CompareHand)
	ans := 0
	for i, hwb := range pairs {
		ans += (i + 1) * hwb.bid
	}
	return ans
}
