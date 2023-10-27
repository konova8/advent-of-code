package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

type Pair struct {
	x int
	y int
}

type Grid map[Pair]rune

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
		ansExample := foo(example, 10)
		fmt.Println("Output Example:", ansExample)
		ansInput := foo(input, 2_000_000)
		fmt.Println("Output Input:", ansInput)
	}
	if part == "2" || part == "" {
		fmt.Println("--- Running part 2 ---")
		ansExample := bar(example, 0, 20)
		fmt.Println("Output Example:", ansExample)
		ansInput := bar(input, 0, 4_000_000)
		fmt.Println("Output Input:", ansInput)
	}
}

func foo(input string, yToCheck int) int {
	pairOfPairs := parseInput(input)
	ans := countNonBeaconPairs(pairOfPairs, yToCheck)
	return ans
}

func bar(input string, lower int, upper int) int {
	pairOfPairs := parseInput(input)
	// Print Graphical Grid (works only for example)
	/*
		grid := Grid{}
		for _, pp := range pairOfPairs {
			grid = markArea(grid, pp[0], pp[1])
		}

		minX, maxX, minY, maxY := getBoundariesGrid(pairOfPairs)
		for y := minY; y <= maxY; y++ {
			for x := minX; x <= maxX; x++ {
				if grid[Pair{x, y}] == 0 {
					fmt.Print(".")
				} else {
					fmt.Print(string(grid[Pair{x, y}]))
				}
			}
			fmt.Println(" ", y)
		}

		for y := lower; y <= upper; y++ {
			for x := lower; x <= upper; x++ {
				if grid[Pair{x, y}] == 0 {
					fmt.Print(".")
				} else {
					fmt.Print(string(grid[Pair{x, y}]))
				}
			}
			fmt.Println(" ", y)
		}
	*/
	ans := getUnknownPairs(pairOfPairs, lower, upper)
	if len(ans) != 1 {
		panic("Need exactly one")
	}
	return ans[0].x*4_000_000 + ans[0].y
}

func parseInput(input string) [][2]Pair {
	res := [][2]Pair{}
	for _, line := range strings.Split(input, "\n") {
		x1 := 0
		y1 := 0
		x2 := 0
		y2 := 0
		n, err := fmt.Sscanf(line,
			"Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&x1, &y1, &x2, &y2)
		if err != nil || n != 4 {
			panic("Sscanf error")
		}
		res = append(res, [2]Pair{{x1, y1}, {x2, y2}})
	}
	return res
}

func generateGrid(pairOfPairs [][2]Pair) Grid {
	grid := Grid{}
	for _, pp := range pairOfPairs {
		sensor := pp[0]
		beacon := pp[1]
		grid = markArea(grid, sensor, beacon)
	}
	return grid
}

func countNonBeaconPairs(pairOfPairs [][2]Pair, y int) int {
	ret := 0
	minX, maxX, _, _ := getBoundariesGrid(pairOfPairs)
	for x := minX; x <= maxX; x++ {
		if !notBeaconSpot(pairOfPairs, Pair{x, y}) {
			ret++
		}
	}
	return ret
}

func getUnknownPairs(pairOfPairs [][2]Pair, lower int, upper int) []Pair {
	ret := []Pair{}
	for y := lower; y <= upper; y++ {
		for x := lower; x <= upper; x++ {
			possible := true
			for _, pp := range pairOfPairs {
				sensor := pp[0]
				beacon := pp[1]
				distanceMax := getDistance(sensor, beacon)
				distanceSpot := getDistance(sensor, Pair{x, y})
				if distanceSpot <= distanceMax {
					x = sensor.x + distanceMax - abs(sensor.y-y)
					if x > upper {
						break
					}
					possible = false
				}
			}
			if possible && x <= upper {
				ret = append(ret, Pair{x, y})
			}
		}
	}
	return ret
}

func notBeaconSpot(pairOfPairs [][2]Pair, spot Pair) bool {
	for _, pp := range pairOfPairs {
		sensor := pp[0]
		beacon := pp[1]
		distanceMax := getDistance(sensor, beacon)
		distanceSpot := getDistance(sensor, spot)
		if distanceSpot <= distanceMax && spot != beacon {
			return false
		}
	}
	return true
}

func getBoundariesGrid(pairOfPairs [][2]Pair) (minX, maxX, minY, maxY int) {
	minX = pairOfPairs[0][0].x
	maxX = pairOfPairs[0][0].x
	minY = pairOfPairs[0][0].y
	maxY = pairOfPairs[0][0].y
	for _, pp := range pairOfPairs {
		sensor := pp[0]
		beacon := pp[1]
		distance := abs(sensor.x-beacon.x) + abs(sensor.y-beacon.y)
		minX = min(minX, distance-sensor.x)
		maxX = max(maxX, distance+sensor.x)
		minY = min(minY, distance-sensor.y)
		maxY = max(maxY, distance+sensor.y)
	}
	return
}

func getDistance(a Pair, b Pair) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func markArea(grid Grid, sensor Pair, beacon Pair) Grid {
	grid[sensor] = 'S'
	grid[beacon] = 'B'
	distance := getDistance(sensor, beacon)
	minX := sensor.x - distance
	maxX := sensor.x + distance
	minY := sensor.y - distance
	maxY := sensor.y + distance
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			diffX := abs(sensor.x - x)
			diffY := abs(sensor.y - y)
			if diffX+diffY <= distance && grid[Pair{x, y}] == 0 {
				grid[Pair{x, y}] = '#'
			}
		}
	}
	return grid
}

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
