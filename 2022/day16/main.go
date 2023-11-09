package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"strings"
)

type Valve struct {
	flowRate     int
	leadToValves []string
}

type Distances map[string]int

type Matrix map[string]Distances

type Pair struct {
	from string
	to   string
}

type Dist map[Pair]int

const N = 100

type GameState struct {
	time   int
	cur    string
	closed [N]string
}

type GameStateBis struct {
	time   int
	cur    string
	closed uint
}

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
	flag.Parse()

	if part == "1" || part == "" {
		fmt.Println("--- Running part 1 ---")
		ansExample := foo(example)
		fmt.Println("Output Example:", ansExample)
		ansInput := foo(input)
		fmt.Println("Output Input:", ansInput)
	}
	if part == "2" || part == "" {
		fmt.Println("--- Running part 2 ---")
		ansExample := bar(example)
		fmt.Println("Output Example:", ansExample)
		ansInput := bar(input)
		fmt.Println("Output Input:", ansInput)
	}
}

func foo(input string) int {
	list := parseInput(input)
	matrix := computeDistancePath(list)
	positive := []string{}
	for k, v := range list {
		if v.flowRate != 0 {
			positive = append(positive, k)
		}
	}
	dist := sanitizeMatrix(matrix, positive)
	dp := map[GameState]int{}
	var bestPossibleFlow func(gs GameState) int
	bestPossibleFlow = func(gs GameState) int {
		if dp[gs] != 0 {
			return dp[gs]
		}
		// Get all closed valves
		closed := clearArrStr(gs.closed)
		// Branch on every possible next open valve
		best := 0
		for _, valve := range closed {
			cost := dist[Pair{gs.cur, valve}] + 1
			if cost < gs.time {
				// Update Closed Valves
				tmp := []string{}
				for _, v := range gs.closed {
					if v != "" && v != valve {
						tmp = append(tmp, v)
					}
				}
				newClosed := [N]string{}
				for i, v := range tmp {
					newClosed[i] = v
				}
				// Generate New State
				newState := GameState{
					time:   gs.time - cost,
					cur:    valve,
					closed: newClosed,
				}
				// Compute New Result
				newStateResult := (gs.time-cost)*list[valve].flowRate +
					bestPossibleFlow(newState)
				// Compare to best (for now)
				best = max(best, newStateResult)
			}
		}
		// Add to Dynamic Programming map
		dp[gs] = best
		return best
	}
	// Initialize positive valve array
	pos := [N]string{}
	for i, v := range positive {
		pos[i] = v
	}
	ans := bestPossibleFlow(GameState{30, "AA", pos})
	return ans
}

func bar(input string) int {
	list := parseInput(input)
	matrix := computeDistancePath(list)
	positive := []string{}
	for k, v := range list {
		if v.flowRate != 0 {
			positive = append(positive, k)
		}
	}
	dist := sanitizeMatrix(matrix, positive)
	dp := map[GameStateBis]int{}
	var bestPossibleFlow func(GameStateBis) int
	bestPossibleFlow = func(gs GameStateBis) int {
		if dp[gs] != 0 {
			return dp[gs]
		}
		// Branch on every possible next open valve
		best := 0
		// Cycle through every positive valve that is still open
		// for p := 1; p < len(positive); p++ {
		for i, valve := range positive {
			var bitPositive uint = 1 << i
			// If positive[p] is still closed
			if (bitPositive & gs.closed) != 0 {
				cost := dist[Pair{gs.cur, valve}] + 1
				if cost < gs.time {
					// Generate New State
					newState := GameStateBis{
						time: gs.time - cost,
						cur:  valve,
						// Update Closed Valves, turning off the bit of the corrisponding valve
						closed: gs.closed & notBit(bitPositive, len(positive)),
					}
					// Compute New Result
					newStateResult := (gs.time-cost)*list[valve].flowRate +
						bestPossibleFlow(newState)
					// Compare to best (for now)
					best = max(best, newStateResult)
				}
			}
		}
		// Add to Dynamic Programming map
		dp[gs] = best
		return best
	}
	// Initialize positive valve bitset
	var bitset uint = (1 << len(positive)) - 1
	best := 0
	// Cycle through each possible combination of closed valves (from me and elephant)
	var i uint
	for i = 0; i <= bitset; i++ {
		var meMask uint = i
		var elephantMask uint = notBit(i, len(positive))
		tmp := bestPossibleFlow(GameStateBis{26, "AA", meMask})
		tmp += bestPossibleFlow(GameStateBis{26, "AA", elephantMask})
		if best < tmp {
			best = tmp
		}
	}
	return best
}

func sanitizeMatrix(m Matrix, positive []string) Dist {
	d := Dist{}
	for from, value := range m {
		for to, dist := range value {
			if from == "AA" || (slices.Contains(positive, from) && slices.Contains(positive, to)) {
				d[Pair{from, to}] = dist
			}
		}
	}
	return d
}

func computeDistancePath(list map[string]Valve) Matrix {
	matrix := Matrix{}
	for k := range list {
		matrix[k] = BFS(list, k)
	}
	return matrix
}

func parseInput(input string) map[string]Valve {
	scanOutput := map[string]Valve{}
	for _, line := range strings.Split(input, "\n") {
		var valve string
		var flowRate int
		var strValves string
		var valves []string
		lineArr := strings.Split(line, "; ")
		fmt.Sscanf(lineArr[0], "Valve %s has flow rate=%d", &valve, &flowRate)
		strValves = strings.SplitN(lineArr[1], " to ", 2)[1]
		strValves = strings.SplitN(strValves, " ", 2)[1]
		for _, v := range strings.Split(strValves, ", ") {
			valves = append(valves, v)
		}
		scanOutput[valve] = Valve{flowRate, valves}
	}
	return scanOutput
}

func BFS(list map[string]Valve, start string) Distances {
	dist := Distances{}
	queue := []string{}
	seen := []string{}
	queue = append(queue, start)
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		seen = append(seen, current)
		for _, v := range list[current].leadToValves {
			if !slices.Contains(seen, v) {
				dist[v] = dist[current] + 1
				queue = append(queue, v)
			}
		}
	}
	return dist
}

func clearArrStr(arr [N]string) []string {
	ret := []string{}
	for _, v := range arr {
		if v != "" {
			ret = append(ret, v)
		}
	}
	return ret
}

// Generic functions for simple math operation
func min[T int | float64](a T, b T) T {
	if a <= b {
		return a
	} else {
		return b
	}
}

func max[T int | float64](a T, b T) T {
	if a >= b {
		return a
	} else {
		return b
	}
}

func abs[T int | float64](n T) T {
	if n < 0 {
		return -n
	}
	return n
}

func notBit(n uint, l int) uint {
	return ^n & ((1 << int(l)) - 1)
}

func getPermutations[T int | string](arr []T) [][]T {
	var helper func([]T, int)
	res := [][]T{}

	helper = func(arr []T, n int) {
		if n == 1 {
			tmp := make([]T, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
