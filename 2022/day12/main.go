package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

type Queue []Pair

type Pair struct {
	x int
	y int
}

type Node struct {
	height int
	edges  []Pair
}

type Graph map[Pair]Node

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
	return BFS(parseInput(input))
}

func bar(input string) int {
	minDepth := 100_000
	G, start, end := parseInput(input)
	starters := []Pair{start}
	starters = append(starters, getAllLetters(G, 'a')...)
	for _, s := range starters {
		tmp := BFS(G, s, end)
		if minDepth > tmp && tmp != -1 {
			minDepth = tmp
		}
	}
	return minDepth
}

func parseInput(input string) (Graph, Pair, Pair) {
	G := Graph{}
	start := Pair{}
	end := Pair{}

	rowInput := strings.Split(input, "\n")
	lenX := len(rowInput[0])
	lenY := len(rowInput)

	for y, line := range rowInput {
		for x, node := range strings.Split(line, "") {
			// Generate Node
			newNode := Node{}
			// Check if it's start or end
			if node == "S" {
				start = Pair{x, y}
				// Get all links
				if y > 0 {
					newNode.edges = append(newNode.edges, Pair{x, y - 1})
				}
				if x > 0 {
					newNode.edges = append(newNode.edges, Pair{x - 1, y})
				}
				if y <= lenY-2 {
					newNode.edges = append(newNode.edges, Pair{x, y + 1})
				}
				if x <= lenX-2 {
					newNode.edges = append(newNode.edges, Pair{x + 1, y})
				}
			} else if node == "E" {
				end = Pair{x, y}
			} else {
				newNode.height = getValue(node[0])
				// Get links
				if y > 0 {
					if getValue(rowInput[y][x]+1) >= getValue(rowInput[y-1][x]) {
						newNode.edges = append(newNode.edges, Pair{x, y - 1})
					}
				}
				if x > 0 {
					if getValue(rowInput[y][x]+1) >= getValue(rowInput[y][x-1]) {
						newNode.edges = append(newNode.edges, Pair{x - 1, y})
					}
				}
				if y <= lenY-2 {
					if getValue(rowInput[y][x]+1) >= getValue(rowInput[y+1][x]) {
						newNode.edges = append(newNode.edges, Pair{x, y + 1})
					}
				}
				if x <= lenX-2 {
					if getValue(rowInput[y][x]+1) >= getValue(rowInput[y][x+1]) {
						newNode.edges = append(newNode.edges, Pair{x + 1, y})
					}
				}
			}
			// Add node to Graph
			G[Pair{x, y}] = newNode
		}
	}

	return G, start, end
}

func getValue[T rune | byte](letter T) int {
	if letter == 'S' {
		return int('a') - 1
	} else if letter == 'E' {
		return int('z') + 1
	} else {
		return int(letter)
	}
}

func BFS(G Graph, start Pair, end Pair) int {
	visited := map[Pair]bool{}
	// Mark all nodes as non visited
	for key := range G {
		visited[key] = false
	}
	// Create queue for nodes
	nodeToVisit := Queue{start}
	// Mark start node as visited
	visited[start] = true
	// Start BFS
	// i = depth of BFS
	for i := 0; len(nodeToVisit) > 0; i++ {
		count := len(nodeToVisit)
		for j := 0; j < count; j++ {
			// Pop one element from queue
			currentNode := nodeToVisit[0]
			nodeToVisit = nodeToVisit[1:]
			if currentNode == end {
				return i
			}
			for _, pair := range G[currentNode].edges {
				// Add node to queue
				if !visited[pair] {
					visited[pair] = true
					nodeToVisit = append(nodeToVisit, pair)
				}
			}
		}
	}
	return -1
}

func getAllLetters[T byte | rune](G Graph, letters ...T) []Pair {
	arr := []Pair{}
	values := []int{}
	for _, l := range letters {
		values = append(values, int(l))
	}
	for k, v := range G {
		for _, n := range values {
			if v.height == n {
				arr = append(arr, k)
			}
		}
	}
	return arr
}
