package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"strconv"
	"strings"
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
	encryptedList := []int{}
	for _, line := range strings.Split(input, "\n") {
		n, err := strconv.Atoi(line)
		if err != nil {
			panic("Atoi conversion falied with line `" + line + "`")
		}
		encryptedList = append(encryptedList, n)
	}
	// fmt.Println(encryptedList)
	unencryptedList := baz(encryptedList)
	// fmt.Println(unencryptedList)
	zeroPosition := slices.Index(unencryptedList, 0)
	if zeroPosition == -1 {
		panic("0 not found in unencripted list")
	}
	ans := unencryptedList[(1000+zeroPosition)%len(unencryptedList)] +
		unencryptedList[(2000+zeroPosition)%len(unencryptedList)] +
		unencryptedList[(3000+zeroPosition)%len(unencryptedList)]
	return ans
}

func baz(list []int) []int {
	indexList := make([]int, len(list))
	for i := range list {
		indexList[i] = i
	}
	newList := []int{}
	for i, v := range list {
		jStart := slices.Index(indexList, i)
		// Cycle for v times, module necessary for counting spaces between number of list
		// We want the number always >= 0, so we add the module
		jStop := (jStart + v + (len(list) - 1)) % (len(list) - 1)
		// fmt.Printf("%d from %d to %d\n", v, jStart, jStop)
		if jStart > jStop {
			// Cycle from jStart to jStop, but backwords
			for j := jStart; j > jStop; j-- {
				indexList[(j+len(list))%len(list)], indexList[(j-1+len(list))%len(list)] =
					indexList[(j-1+len(list))%len(list)], indexList[(j+len(list))%len(list)]
			}
		} else {
			// Cycle from jStart to jStop
			for j := jStart; j < jStop; j++ {
				indexList[(j+len(list))%len(list)], indexList[(j+1+len(list))%len(list)] =
					indexList[(j+1+len(list))%len(list)], indexList[(j+len(list))%len(list)]
			}
		}
		// fmt.Println("indexList: ", indexList)
		newList = []int{}
		for _, v := range indexList {
			newList = append(newList, list[v])
		}
		// fmt.Println("newList: ", newList)
	}
	return newList
}

func bar(input string) int {
	encryptedList := []int{}
	for _, line := range strings.Split(input, "\n") {
		n, err := strconv.Atoi(line)
		if err != nil {
			panic("Atoi conversion falied with line `" + line + "`")
		}
		encryptedList = append(encryptedList, n)
	}
	// fmt.Println(encryptedList)
	unencryptedList := baz2(encryptedList)
	// fmt.Println(unencryptedList)
	zeroPosition := slices.Index(unencryptedList, 0)
	if zeroPosition == -1 {
		panic("0 not found in unencripted list")
	}
	ans := unencryptedList[(1000+zeroPosition)%len(unencryptedList)] +
		unencryptedList[(2000+zeroPosition)%len(unencryptedList)] +
		unencryptedList[(3000+zeroPosition)%len(unencryptedList)]
	return ans
}

func baz2(list []int) []int {
	decryptionKey := 811589153
	for i := range list {
		list[i] *= decryptionKey
	}
	indexList := make([]int, len(list))
	for i := range list {
		indexList[i] = i
	}
	for k := 0; k < 10; k++ {
		for i, v := range list {
			jStart := slices.Index(indexList, i)
			// Cycle for v times, module necessary for counting spaces between number of list
			// We want the number always >= 0, so we add the module
			jStop := (jStart + v) % (len(list) - 1)
			if jStop < 0 {
				jStop += len(list) - 1
			}
			// fmt.Printf("%d from %d to %d\n", v, jStart, jStop)
			if jStart > jStop {
				// Cycle from jStart to jStop, but backwords
				for j := jStart; j > jStop; j-- {
					jIndex := j % len(list)
					if jIndex < 0 {
						jIndex += len(list)
					}
					jm1Index := (j - 1) % len(list)
					if jm1Index < 0 {
						jm1Index += len(list)
					}
					indexList[jIndex], indexList[jm1Index] =
						indexList[jm1Index], indexList[jIndex]
				}
			} else {
				// Cycle from jStart to jStop
				for j := jStart; j < jStop; j++ {
					jIndex := j % len(list)
					if jIndex < 0 {
						jIndex += len(list)
					}
					jp1Index := (j + 1) % len(list)
					if jp1Index < 0 {
						jp1Index += len(list)
					}
					indexList[jIndex], indexList[jp1Index] =
						indexList[jp1Index], indexList[jIndex]
				}
			}
			// fmt.Println("indexList: ", indexList)
		}
	}
	newList := []int{}
	for _, v := range indexList {
		newList = append(newList, list[v])
	}
	// fmt.Println("newList: ", newList)
	return newList
}
