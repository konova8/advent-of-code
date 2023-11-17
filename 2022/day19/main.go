package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed example.txt
var example string

//go:embed input.txt
var input string

type Blueprint struct {
	costOreRobot      int
	costClayRobot     int
	costObsidianRobot [2]int
	costGeodeRobot    [2]int
}

type Rocks struct {
	ore      int
	clay     int
	obsidian int
	geode    int
}

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
			ansExample := foo(example)
			fmt.Println("Output Example:", ansExample)
		}
		if !*noInput {
			ansInput := foo(input)
			fmt.Println("Output Input:", ansInput)
		}
	}
	if part == "2" || part == "" {
		fmt.Println("--- Running part 2 ---")
		if !*noExample {
			ansExample := bar(example)
			fmt.Println("Output Example:", ansExample)
		}
		if !*noInput {
			ansInput := bar(input)
			fmt.Println("Output Input:", ansInput)
		}
	}
}

func foo(input string) int {
	ans := 0
	allBlueprintsResult := baz(input, 24, -1)
	for i, v := range allBlueprintsResult {
		ans += (i + 1) * v
	}
	return ans
}

func bar(input string) int {
	ans := 1
	allBlueprintsResult := baz(input, 32, 3)
	for _, v := range allBlueprintsResult {
		ans *= v
	}
	return ans
}

func baz(input string, minutes int, maxBlueprints int) []int {
	ans := []int{}
	r, err := regexp.Compile(`Blueprint \d+: Each ore robot costs (\d+) ore. Each clay robot costs (\d+) ore. Each obsidian robot costs (\d+) ore and (\d+) clay. Each geode robot costs (\d+) ore and (\d+) obsidian.`)
	if err != nil {
		panic("regex is wrong")
	}
	blueprints := []Blueprint{}
	for _, line := range strings.Split(input, "\n") {
		res := r.FindAllStringSubmatch(line, -1)
		tmp1, _ := strconv.Atoi(res[0][1])
		tmp2, _ := strconv.Atoi(res[0][2])
		tmp3, _ := strconv.Atoi(res[0][3])
		tmp4, _ := strconv.Atoi(res[0][4])
		tmp5, _ := strconv.Atoi(res[0][5])
		tmp6, _ := strconv.Atoi(res[0][6])
		newBlueprint := Blueprint{
			costOreRobot:      tmp1,
			costClayRobot:     tmp2,
			costObsidianRobot: [2]int{tmp3, tmp4},
			costGeodeRobot:    [2]int{tmp5, tmp6},
		}
		blueprints = append(blueprints, newBlueprint)
	}
	type GameState struct {
		b      Blueprint
		m      int
		bin    Rocks
		robots Rocks
	}
	cache := map[GameState]Rocks{}
	emptyBin := Rocks{}
	var bestOverall int = 0
	var simulateBlueprint func(b Blueprint, m int, bin, robots, maxRobots Rocks) Rocks
	simulateBlueprint = func(b Blueprint, m int, bin, robots, maxRobots Rocks) Rocks {
		if m == 0 {
			return bin
		}
		gs := GameState{b, m, bin, robots}
		if cache[gs] != emptyBin {
			return bin
		}
		bestBin := Rocks{bin.ore, bin.clay, bin.obsidian, bin.geode}
		maxGeode := bin.geode
		possibleGeodeRobots := robots.geode
		for i := 0; i < m; i++ {
			maxGeode += possibleGeodeRobots
			possibleGeodeRobots++
		}
		if bestOverall >= maxGeode {
			// Cut Branch
			return bin
		}
		if bin.ore >= b.costGeodeRobot[0] && bin.obsidian >= b.costGeodeRobot[1] {
			// Can build a Geode Robot
			newBin := Rocks{bin.ore - b.costGeodeRobot[0], bin.clay, bin.obsidian - b.costGeodeRobot[1], bin.geode}
			newBin.ore += robots.ore
			newBin.clay += robots.clay
			newBin.obsidian += robots.obsidian
			newBin.geode += robots.geode
			newRobots := Rocks{robots.ore, robots.clay, robots.obsidian, robots.geode + 1}
			resultSimulation := simulateBlueprint(b, m-1, newBin, newRobots, maxRobots)
			if resultSimulation.geode > bestOverall {
				bestOverall = resultSimulation.geode
			}
			if bestBin.geode < resultSimulation.geode {
				bestBin = resultSimulation
			}
		} else {
			if bin.ore >= b.costOreRobot && robots.ore < maxRobots.ore {
				// Can build a Ore Robot
				newBin := Rocks{bin.ore - b.costOreRobot, bin.clay, bin.obsidian, bin.geode}
				newBin.ore += robots.ore
				newBin.clay += robots.clay
				newBin.obsidian += robots.obsidian
				newBin.geode += robots.geode
				newRobots := Rocks{robots.ore + 1, robots.clay, robots.obsidian, robots.geode}
				resultSimulation := simulateBlueprint(b, m-1, newBin, newRobots, maxRobots)
				if resultSimulation.geode > bestOverall {
					bestOverall = resultSimulation.geode
				}
				if bestBin.geode < resultSimulation.geode {
					bestBin = resultSimulation
				}
			}
			if bin.ore >= b.costClayRobot && robots.clay < maxRobots.clay {
				// Can build a Clay Robot
				newBin := Rocks{bin.ore - b.costClayRobot, bin.clay, bin.obsidian, bin.geode}
				newBin.ore += robots.ore
				newBin.clay += robots.clay
				newBin.obsidian += robots.obsidian
				newBin.geode += robots.geode
				newRobots := Rocks{robots.ore, robots.clay + 1, robots.obsidian, robots.geode}
				resultSimulation := simulateBlueprint(b, m-1, newBin, newRobots, maxRobots)
				if resultSimulation.geode > bestOverall {
					bestOverall = resultSimulation.geode
				}
				if bestBin.geode < resultSimulation.geode {
					bestBin = resultSimulation
				}
			}
			if bin.ore >= b.costObsidianRobot[0] && bin.clay >= b.costObsidianRobot[1] && robots.obsidian < maxRobots.obsidian {
				// Can build a Obsidian Robot
				newBin := Rocks{bin.ore - b.costObsidianRobot[0], bin.clay - b.costObsidianRobot[1], bin.obsidian, bin.geode}
				newBin.ore += robots.ore
				newBin.clay += robots.clay
				newBin.obsidian += robots.obsidian
				newBin.geode += robots.geode
				newRobots := Rocks{robots.ore, robots.clay, robots.obsidian + 1, robots.geode}
				resultSimulation := simulateBlueprint(b, m-1, newBin, newRobots, maxRobots)
				if resultSimulation.geode > bestOverall {
					bestOverall = resultSimulation.geode
				}
				if bestBin.geode < resultSimulation.geode {
					bestBin = resultSimulation
				}
			}
		}
		bin.ore += robots.ore
		bin.clay += robots.clay
		bin.obsidian += robots.obsidian
		bin.geode += robots.geode
		resultSimulation := simulateBlueprint(b, m-1, bin, robots, maxRobots)
		if resultSimulation.geode > bestOverall {
			bestOverall = resultSimulation.geode
		}
		if bestBin.geode < resultSimulation.geode {
			bestBin = resultSimulation
		}
		cache[gs] = bestBin
		return bestBin
	}
	if maxBlueprints >= len(blueprints) {
		maxBlueprints = len(blueprints)
	} else if maxBlueprints != -1 {
		blueprints = blueprints[:maxBlueprints]
	}
	for _, b := range blueprints {
		// Compute maximum meaningful value for robots
		maxObsidian := b.costGeodeRobot[1]
		maxClay := b.costObsidianRobot[1]
		maxOre := max(b.costOreRobot, b.costClayRobot, b.costObsidianRobot[0], b.costGeodeRobot[0])
		maxRobots := Rocks{
			ore:      maxOre,
			clay:     maxClay,
			obsidian: maxObsidian,
			geode:    10000,
		}
		bestOverall = 0
		cache = map[GameState]Rocks{}
		tmp := simulateBlueprint(b, minutes, Rocks{}, Rocks{ore: 1}, maxRobots)
		ans = append(ans, tmp.geode)
	}
	return ans
}
