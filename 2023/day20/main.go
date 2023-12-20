package main

import (
	"advent-of-code/util"
	_ "embed"
	"flag"
	"fmt"
	"slices"
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
			ansExample := fmt.Sprint(part2(example2))
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

const (
	Low = iota
	High
)

const (
	Off = iota
	On
)

type Module struct {
	name      string
	kind      rune
	to        []string
	lastPulse map[string]int
	status    int
}

func parseInput(str string) map[string]Module {
	allOutputs := map[string]bool{}
	modules := map[string]Module{}
	for _, line := range strings.Split(str, "\n") {
		tmp := strings.Split(line, " -> ")
		left := tmp[0]
		right := tmp[1]
		m := Module{}
		if left[0] == 'b' {
			m.name = left
			m.kind = 'b'
		} else {
			m.kind = rune(left[0])
			m.name = left[1:]
		}
		m.to = strings.Split(right, ", ")
		for _, v := range m.to {
			allOutputs[v] = true
		}
		m.lastPulse = map[string]int{}
		m.status = Off
		modules[m.name] = m
	}
	for _, mFrom := range modules {
		for _, v := range mFrom.to {
			mTo := modules[v]
			if mTo.kind == '&' {
				mTo.lastPulse[mFrom.name] = Low
			}
		}
	}
	for k := range allOutputs {
		_, found := modules[k]
		if !found {
			m := Module{
				name:      k,
				kind:      'o',
				to:        []string{},
				status:    Off,
				lastPulse: nil,
			}
			modules[m.name] = m
		}
	}
	return modules
}

type Pulse struct {
	kind int
	from string
	to   string
}

func computePulse(p Pulse, modules map[string]Module) []Pulse {
	ret := []Pulse{}
	m := modules[p.to]
	switch m.kind {
	case 'b':
		for _, v := range m.to {
			ret = append(ret, Pulse{
				from: m.name,
				kind: p.kind,
				to:   v,
			})
		}
		return ret
	case '%':
		if p.kind == High {
			return ret
		}
		var k int
		if m.status == On {
			old := m
			old.status = Off
			modules[p.to] = old
			k = Low
		} else {
			old := m
			old.status = On
			modules[p.to] = old
			k = High
		}
		for _, v := range m.to {
			ret = append(ret, Pulse{
				from: m.name,
				kind: k,
				to:   v,
			})
		}
		return ret
	case '&':
		m.lastPulse[p.from] = p.kind
		count := 0
		for _, v := range m.lastPulse {
			if v == High {
				count++
			}
		}
		var k int
		if count == len(m.lastPulse) {
			k = Low
		} else {
			k = High
		}
		for _, v := range m.to {
			ret = append(ret, Pulse{
				from: m.name,
				kind: k,
				to:   v,
			})
		}
		return ret
	case 'o':
		return ret
	default:
		panic("Shouldn't be here")
	}
}

func part1(str string) int {
	modules := parseInput(str)
	initialPulse := Pulse{
		from: "button",
		kind: Low,
		to:   "broadcaster",
	}
	countLow := 0
	countHigh := 0
	for i := 0; i < 1000; i++ {
		Q := []Pulse{initialPulse}
		for len(Q) > 0 {
			p := Q[0]
			Q = Q[1:]
			if p.kind == Low {
				countLow++
			} else {
				countHigh++
			}
			newPulses := computePulse(p, modules)
			Q = append(Q, newPulses...)
		}
	}
	ans := countLow * countHigh
	return ans
}

func part2(str string) int {
	modules := parseInput(str)
	var lastInverter string
	for _, m := range modules {
		if len(m.to) == 1 && m.to[0] == "rx" {
			lastInverter = m.name
		}
	}
	toCheck := []string{}
	for _, m := range modules {
		if slices.Contains(m.to, lastInverter) {
			toCheck = append(toCheck, m.name)
		}
	}
	initialPulse := Pulse{}
	initialPulse.from = "button"
	initialPulse.kind = Low
	initialPulse.to = "broadcaster"
	arr := []int{}
	for _, s := range toCheck {
		found := false
		for k := range modules {
			old := modules[k]
			old.status = Off
			for k2 := range old.lastPulse {
				old.lastPulse[k2] = Low
			}
			modules[k] = old
		}
		for i := 1; i <= 1_000_000_000 && !found; i++ {
			Q := []Pulse{initialPulse}
			for len(Q) > 0 && !found {
				p := Q[0]
				Q = Q[1:]
				newPulses := computePulse(p, modules)
				Q = append(Q, newPulses...)
				if modules[lastInverter].lastPulse[s] == High {
					found = true
					arr = append(arr, i)
				}
			}
		}
	}
	return util.LCM(arr...)
}
