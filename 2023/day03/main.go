package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"advent-of-code/util"
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
		}
		if !*noInput {
			ansInput := part1(input)
			fmt.Println("Output Input:", ansInput)
		}
	}
	if part == "2" || part == "" {
		fmt.Println("--- Running part 2 ---")
		if !*noExample {
			ansExample := part2(example)
			fmt.Println("Output Example:", ansExample)
		}
		if !*noInput {
			ansInput := part2(input)
			fmt.Println("Output Input:", ansInput)
		}
	}
}

type Number struct {
	start util.Position
	end   util.Position
	value int
}

type Symbol struct {
	pos   util.Position
	value rune
}

const gearRune = '*'

func getPositions(str string) (map[util.Position]bool, []Number, []Symbol) {
	symbolMap := map[util.Position]bool{}
	numbers := []Number{}
	symbols := []Symbol{}
	x := 0
	y := 0
	num := ""
	startN := util.Position{X: -1, Y: -1}
	endN := util.Position{X: -1, Y: -1}
	for i := 0; i < len(str); i++ {
		r := str[i]
		if r == '\n' {
			x = 0
			y++
			continue
		} else if r == '.' {
			if startN.X == -1 && startN.Y == -1 {
				x++
				continue
			}
			n, err := strconv.Atoi(num)
			if err != nil {
				panic("error while converting number inside getPositions")
			}
			numbers = append(numbers, Number{
				start: startN,
				end:   endN,
				value: n,
			})
			num = ""
			startN = util.Position{X: -1, Y: -1}
			endN = util.Position{X: -1, Y: -1}
		} else if unicode.IsDigit(rune(r)) {
			if startN.X == -1 && startN.Y == -1 {
				startN = util.Position{X: x, Y: y}
			}
			endN = util.Position{X: x, Y: y}
			num += string(r)
		} else {
			symbolMap[util.Position{X: x, Y: y}] = true
			symbols = append(symbols, Symbol{
				pos:   util.Position{X: x, Y: y},
				value: rune(r),
			})
			if startN.X != -1 && startN.Y != -1 {
				n, err := strconv.Atoi(num)
				if err != nil {
					panic("error while converting number inside getPositions")
				}
				numbers = append(numbers, Number{
					start: startN,
					end:   endN,
					value: n,
				})
				num = ""
				startN = util.Position{X: -1, Y: -1}
				endN = util.Position{X: -1, Y: -1}
			}
		}
		x++
	}
	return symbolMap, numbers, symbols
}

func (n Number) isNear(symbols map[util.Position]bool) bool {
	for x := n.start.X - 1; x <= n.end.X+1; x++ {
		if symbols[util.Position{X: x, Y: n.start.Y - 1}] ||
			symbols[util.Position{X: x, Y: n.start.Y + 1}] {
			return true
		}
	}
	if symbols[util.Position{X: n.start.X - 1, Y: n.start.Y}] ||
		symbols[util.Position{X: n.end.X + 1, Y: n.start.Y}] {
		return true
	}
	return false
}

func part1(str string) int {
	ans := 0
	symbolMap, numbers, _ := getPositions(str)
	for _, n := range numbers {
		if n.isNear(symbolMap) {
			ans += n.value
		}
	}
	return ans
}

func generateNumberMap(numbers []Number) map[util.Position]int {
	numberMap := map[util.Position]int{}
	for _, n := range numbers {
		y := n.start.Y
		for x := n.start.X; x <= n.end.X; x++ {
			numberMap[util.Position{X: x, Y: y}] = n.value
		}
	}
	return numberMap
}

var adjOffsets = [][2]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func getAdjNum(numberMap map[util.Position]int, s Symbol) []int {
	ret := []int{}
	set := map[int]bool{}
	for _, v := range adjOffsets {
		n := numberMap[util.Position{X: s.pos.X + v[0], Y: s.pos.Y + v[1]}]
		if n != 0 {
			set[n] = true
		}
	}
	// Take unique values
	for k := range set {
		ret = append(ret, k)
	}
	return ret
}

func part2(str string) int {
	ans := 0
	_, numbers, symbols := getPositions(str)
	numberMap := generateNumberMap(numbers)
	for _, s := range symbols {
		if s.value == gearRune {
			numbers := getAdjNum(numberMap, s)
			if len(numbers) == 2 {
				ans += numbers[0] * numbers[1]
			}
		}
	}
	return ans
}
