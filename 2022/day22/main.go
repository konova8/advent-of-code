package main

import (
	_ "embed"
	"flag"
	"fmt"
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
			fmt.Println("Output Example: NOT IMPLEMENTED")
		}
		if !*noInput {
			ansInput := bar(input)
			fmt.Println("Output Input:", ansInput)
		}
	}
}

type Position struct {
	row    uint8
	column uint8
	facing uint8
}

const hashByteValue = byte('#')
const spaceByteValue = byte(' ')

func Move(grid []string, pos Position, move interface{}) Position {
	str, isString := move.(string)
	if isString {
		if str == "R" {
			pos.facing = (pos.facing + 1) % 4
		} else {
			pos.facing = (pos.facing - 1) % 4
		}
		return pos
	}
	num, isInt := move.(int)
	if isInt {
	Loop:
		for i := 0; i < num; i++ {
			switch pos.facing {
			case 0:
				newRow := pos.row
				newColumn := pos.column + 1
				if newColumn >= uint8(len(grid[newRow])) {
					newColumn = firstNonWhiteCharRow(grid, newRow)
				}
				if grid[newRow][newColumn] == hashByteValue {
					break Loop
				} else {
					pos.column = newColumn
				}
			case 1:
				newRow := pos.row + 1
				newColumn := pos.column
				if newRow >= uint8(len(grid)) || newColumn >= uint8(len(grid[newRow])) || grid[newRow][newColumn] == spaceByteValue {
					newRow = firstNonWhiteCharCol(grid, newColumn)
				}
				if grid[newRow][newColumn] == hashByteValue {
					break Loop
				} else {
					pos.row = newRow
				}
			case 2:
				newRow := pos.row
				newColumn := pos.column - 1
				if newColumn == 255 || grid[newRow][newColumn] == spaceByteValue {
					newColumn = lastNonWhiteCharRow(grid, newRow)
				}
				if grid[newRow][newColumn] == hashByteValue {
					break Loop
				} else {
					pos.column = newColumn
				}
			case 3:
				newRow := pos.row - 1
				newColumn := pos.column
				if newRow == 255 || newColumn >= uint8(len(grid[newRow])) || grid[newRow][newColumn] == spaceByteValue {
					newRow = lastNonWhiteCharCol(grid, newColumn)
				}
				if grid[newRow][newColumn] == hashByteValue {
					break Loop
				} else {
					pos.row = newRow
				}
			}
		}
		return pos
	}
	panic("`move` is not string nor number")
}

func firstNonWhiteCharCol(grid []string, col uint8) uint8 {
	for i := 0; i < len(grid); i++ {
		v := grid[i]
		if col < uint8(len(v)) && v[col] != ' ' {
			return uint8(i)
		}
	}
	return uint8(255)
}

func lastNonWhiteCharCol(grid []string, col uint8) uint8 {
	for i := len(grid) - 1; i >= 0; i-- {
		v := grid[i]
		if col < uint8(len(v)) && v[col] != ' ' {
			return uint8(i)
		}
	}
	return uint8(255)
}

func firstNonWhiteCharRow(grid []string, row uint8) uint8 {
	return uint8(strings.IndexFunc(grid[row], func(r rune) bool { return r != ' ' }))
}

func lastNonWhiteCharRow(grid []string, row uint8) uint8 {
	return uint8(strings.LastIndexFunc(grid[row], func(r rune) bool { return r != ' ' }))
}

func foo(input string) int {
	grid, moves := parseInput(input)
	// Find initial position
	pos := Position{
		row:    0,
		column: firstNonWhiteCharRow(grid, 0),
		facing: 0,
	}
	for _, m := range moves {
		pos = Move(grid, pos, m)
	}
	return 1000*(int(pos.row)+1) + 4*(int(pos.column)+1) + int(pos.facing)
}

func parseInput(input string) (grid []string, moves []interface{}) {
	arr := strings.Split(input, "\n\n")
	grid = append(grid, strings.Split(arr[0], "\n")...)
	num := 0
	for i := 0; i < len(arr[1]); i++ {
		c := arr[1][i]
		switch c {
		case byte('L'):
			moves = append(moves, num)
			moves = append(moves, "L")
			num = 0
		case byte('R'):
			moves = append(moves, num)
			moves = append(moves, "R")
			num = 0
		default:
			num = num*10 + (int(c) - 48)
		}
	}
	if num != 0 {
		moves = append(moves, num)
	}
	return
}

func bar(input string) int {
	grid, moves := parseInput(input)
	warps := generateWarps(grid)
	// Find initial position
	pos := Position{
		row:    0,
		column: firstNonWhiteCharRow(grid, 0),
		facing: 0,
	}
	for _, m := range moves {
		pos = Move2(grid, warps, pos, m)
	}
	return 1000*(int(pos.row)+1) + 4*(int(pos.column)+1) + int(pos.facing)
}

func generateWarps(grid []string) map[Position]Position {
	warps := map[Position]Position{}
	var dim uint8
	var i_row, i_col, j_row, j_col uint8
	if len(grid)%50 == 0 {
		dim = 50
		i_row = (dim * 0)
		i_col = (dim * 1)
		j_row = (dim * 3)
		j_col = (dim * 0)
		for i := 0; i < int(dim); i++ {
			// U -> B
			warps[Position{i_row - 1, i_col, 3}] = Position{j_row, j_col, 0}
			// B -> U
			warps[Position{j_row, j_col - 1, 2}] = Position{i_row, i_col, 1}
			i_col++
			j_row++
		}
		i_row = (dim * 0)
		i_col = (dim * 2)
		j_row = (dim * 4) - 1
		j_col = (dim * 0)
		for i := 0; i < int(dim); i++ {
			// R -> B
			warps[Position{i_row - 1, i_col, 3}] = Position{j_row, j_col, 3}
			// B -> R
			warps[Position{j_row + 1, j_col, 1}] = Position{i_row, i_col, 1}
			i_col++
			j_col++
		}
		i_row = (dim * 0)
		i_col = (dim * 1)
		j_row = (dim * 3) - 1
		j_col = (dim * 0)
		for i := 0; i < int(dim); i++ {
			// U -> L
			warps[Position{i_row, i_col - 1, 2}] = Position{j_row, j_col, 0}
			// L -> U
			warps[Position{j_row, j_col - 1, 2}] = Position{i_row, i_col, 0}
			i_row++
			j_row--
		}
		i_row = (dim * 0)
		i_col = (dim * 3) - 1
		j_row = (dim * 3) - 1
		j_col = (dim * 2) - 1
		for i := 0; i < int(dim); i++ {
			// R -> D
			warps[Position{i_row, i_col + 1, 0}] = Position{j_row, j_col, 2}
			// D -> R
			warps[Position{j_row, j_col + 1, 0}] = Position{i_row, i_col, 2}
			i_row++
			j_row--
		}
		i_row = (dim * 1)
		i_col = (dim * 1)
		j_row = (dim * 2)
		j_col = (dim * 0)
		for i := 0; i < int(dim); i++ {
			// F -> L
			warps[Position{i_row, i_col - 1, 2}] = Position{j_row, j_col, 1}
			// L -> F
			warps[Position{j_row - 1, j_col, 3}] = Position{i_row, i_col, 0}
			i_row++
			j_col++
		}
		i_row = (dim * 1) - 1
		i_col = (dim * 2)
		j_row = (dim * 1)
		j_col = (dim * 2) - 1
		for i := 0; i < int(dim); i++ {
			// R -> F
			warps[Position{i_row + 1, i_col, 1}] = Position{j_row, j_col, 2}
			// F -> R
			warps[Position{j_row, j_col + 1, 0}] = Position{i_row, i_col, 3}
			i_col++
			j_row++
		}
		i_row = (dim * 3) - 1
		i_col = (dim * 1)
		j_row = (dim * 3)
		j_col = (dim * 1) - 1
		for i := 0; i < int(dim); i++ {
			// D -> B
			warps[Position{i_row + 1, i_col, 1}] = Position{j_row, j_col, 2}
			// B -> D
			warps[Position{j_row, j_col + 1, 0}] = Position{i_row, i_col, 3}
			i_col++
			j_row++
		}
	} else if len(grid)%4 == 0 {
		dim = 4
		// U -> L
		// U <- L
		// U -> B
		// U <- B
		// F -> R
		// F <- R
		// R -> B
		// R <- B
		// L -> D
		// L <- D
		// B -> D
		// B <- D
	}
	return warps
}

func Move2(grid []string, warps map[Position]Position, pos Position, move interface{}) Position {
	str, isString := move.(string)
	if isString {
		if str == "R" {
			pos.facing = (pos.facing + 1) % 4
		} else {
			pos.facing = (pos.facing - 1) % 4
		}
		return pos
	}
	num, isInt := move.(int)
	if isInt {
	Loop:
		for i := 0; i < num; i++ {
			newRow := pos.row
			newColumn := pos.column
			newFacing := pos.facing
			switch pos.facing {
			case 0:
				newColumn = pos.column + 1
			case 1:
				newRow = pos.row + 1
			case 2:
				newColumn = pos.column - 1
			case 3:
				newRow = pos.row - 1
			}
			newPos := warps[Position{newRow, newColumn, pos.facing}]
			if !(newPos.row == 0 && newPos.column == 0) {
				newRow = newPos.row
				newColumn = newPos.column
				newFacing = newPos.facing
			}
			if grid[newRow][newColumn] == hashByteValue {
				break Loop
			} else {
				pos.row = newRow
				pos.column = newColumn
				pos.facing = newFacing
			}
		}
		return pos
	}
	panic("`move` is not string nor number")
}
