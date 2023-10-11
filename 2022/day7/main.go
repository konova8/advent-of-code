package main

import (
	_ "embed"
	"errors"
	"flag"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type File struct {
	parent  *File
	name    string
	size    int
	content map[string]File
}

var rootFs File = File{
	parent:  nil,
	name:    "/",
	size:    0,
	content: map[string]File{},
}

// Must have 0 as parameter when called
func (file File) print(level int) {
	if file.size == 0 {
		fmt.Printf(strings.Repeat(" ", level)+"- %s (dir)\n", file.name)
		// For printing in order
		keys := []string{}
		for k := range file.content {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			file.content[k].print(level + 1)
		}
	} else {
		fmt.Printf(strings.Repeat(" ", level)+"- %s (file)\n", file.name)
	}
}

func (file File) toRoot() File {
	for file.parent != nil {
		file = *file.parent
	}
	return file
}

func (file File) changeDir(nameDir string) File {
	if nameDir == ".." {
		file = *file.parent
	} else {
		file = file.content[nameDir]
	}
	return file
}

// Add ls output to dir file
func (file File) expandDir(outputLS []string) {
	for _, line := range outputLS {
		arr := strings.Split(line, " ")
		if arr[0] != "dir" {
			size, _ := strconv.Atoi(arr[0])
			name := arr[1]
			file.content[name] = File{
				parent:  &file,
				name:    name,
				size:    size,
				content: nil,
			}
		} else {
			name := arr[1]
			file.content[name] = File{
				parent:  &file,
				name:    name,
				size:    0,
				content: map[string]File{},
			}
		}
	}
}

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		fmt.Println("Output:", ans)
	} else if part == 2 {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	generatedFS, _ := generateFS(input)
	generatedFS = generatedFS.toRoot()
	// generatedFS.print(0)
	_, res := getSumSizeWithLimit(generatedFS, 100000, 0)
	return res
}

func part2(input string) int {
	generatedFS, _ := generateFS(input)
	generatedFS = generatedFS.toRoot()
	// generatedFS.print(0)
	usedSpace, res := getListDirSize(generatedFS, []int{})
	slices.Sort(res)
	atLeast := 30_000_000 - (70_000_000 - usedSpace)
	for _, n := range res {
		if n >= atLeast {
			return n
		}
	}
	return -1
}
func generateFS(input string) (File, error) {
	fs := rootFs
	arg := ""
	lines := strings.Split(input, "\n")
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if line[0] == '$' {
			if line[2:4] == "cd" {
				arg = line[5:]
				if arg == "/" {
					fs.toRoot()
				} else {
					fs = fs.changeDir(arg)
				}
			} else if line[2:4] == "ls" {
				outputLS := []string{}
				for ; i+1 != len(lines) && lines[i+1][0] != '$'; i++ {
					outputLS = append(outputLS, lines[i+1])
				}
				fs.expandDir(outputLS)
			} else {
				err := errors.New("command different from cd or ls")
				return fs, err
			}
			// } else if line[0] == 'd' || ('0' <= line[0] && line[0] <= '9') {
			// 	fmt.Println("FILE DESC")
		} else {
			err := errors.New("Not command nor ls output")
			return fs, err
		}
	}
	return fs, nil
}

func getSumSizeWithLimit(file File, limit int, res int) (int, int) {
	ans := 0
	if file.size == 0 {
		// for analyzing in order
		keys := []string{}
		for k := range file.content {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			dirSize, tmp := getSumSizeWithLimit(file.content[k], limit, res)
			res = tmp
			ans += dirSize
		}
		if ans <= limit {
			res += ans
		}
	} else {
		ans += file.size
	}
	return ans, res
}

func getListDirSize(file File, res []int) (int, []int) {
	ans := 0
	if file.size == 0 {
		// for analyzing in order
		keys := []string{}
		for k := range file.content {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			totalDirSize, updatedRes := getListDirSize(file.content[k], res)
			res = updatedRes
			res = append(res, totalDirSize)
			ans += totalDirSize
		}
	} else {
		ans += file.size
	}
	return ans, res
}
