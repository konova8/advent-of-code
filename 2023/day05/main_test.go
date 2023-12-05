package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func SortRange(a, b Range) int {
	if a.start != b.start {
		return a.start - b.start
	} else if a.end != b.end {
		return a.end - b.end
	}
	return 0
}

func TestGenerateNewRange1(tt *testing.T) {
	r := Range{3, 7, 5}
	t := Triplet{11, 1, 9}
	fmt.Println(r.toString())
	fmt.Println(t.toString())
	fmt.Println(r, t)
	actual := generateNewRanges(r, []Triplet{t})
	expected := []Range{
		{13, 17, 5},
	}
	slices.SortFunc(expected, SortRange)
	slices.SortFunc(actual, SortRange)
	assert.Equal(tt, expected, actual)
}

func TestGenerateNewRange2(tt *testing.T) {
	r := Range{1, 9, 9}
	t := Triplet{13, 3, 5}
	fmt.Println(r.toString())
	fmt.Println(t.toString())
	fmt.Println(r, t)
	actual := generateNewRanges(r, []Triplet{t})
	expected := []Range{
		{1, 2, 2},
		{13, 17, 5},
		{8, 9, 2},
	}
	slices.SortFunc(expected, SortRange)
	slices.SortFunc(actual, SortRange)
	assert.Equal(tt, expected, actual)
}

func TestGenerateNewRange3(tt *testing.T) {
	r := Range{3, 7, 5}
	t := Triplet{11, 1, 5}
	fmt.Println(r.toString())
	fmt.Println(t.toString())
	fmt.Println(r, t)
	actual := generateNewRanges(r, []Triplet{t})
	expected := []Range{
		{13, 15, 3},
		{6, 7, 2},
	}
	slices.SortFunc(expected, SortRange)
	slices.SortFunc(actual, SortRange)
	assert.Equal(tt, expected, actual)
}

func TestGenerateNewRange4(tt *testing.T) {
	r := Range{3, 7, 5}
	t := Triplet{15, 5, 5}
	fmt.Println(r.toString())
	fmt.Println(t.toString())
	fmt.Println(r, t)
	actual := generateNewRanges(r, []Triplet{t})
	expected := []Range{
		{3, 4, 2},
		{15, 17, 3},
	}
	slices.SortFunc(expected, SortRange)
	slices.SortFunc(actual, SortRange)
	assert.Equal(tt, expected, actual)
}

func TestGenerateNewRange5(tt *testing.T) {
	r := Range{1, 4, 4}
	t := Triplet{17, 7, 2}
	fmt.Println(r.toString())
	fmt.Println(t.toString())
	fmt.Println(r, t)
	actual := generateNewRanges(r, []Triplet{t})
	expected := []Range{
		{1, 4, 4},
	}
	slices.SortFunc(expected, SortRange)
	slices.SortFunc(actual, SortRange)
	assert.Equal(tt, expected, actual)
}

func TestGenerateNewRange6(tt *testing.T) {
	r := Range{5, 8, 4}
	t := Triplet{12, 2, 2}
	fmt.Println(r.toString())
	fmt.Println(t.toString())
	fmt.Println(r, t)
	actual := generateNewRanges(r, []Triplet{t})
	expected := []Range{
		{5, 8, 4},
	}
	slices.SortFunc(expected, SortRange)
	slices.SortFunc(actual, SortRange)
	assert.Equal(tt, expected, actual)
}

func TestGenerateNewRange7(tt *testing.T) {
	r := Range{5, 8, 4}
	t1 := Triplet{17, 7, 2}
	t2 := Triplet{12, 2, 2}
	fmt.Println(r.toString())
	fmt.Println(t1.toString())
	fmt.Println(t2.toString())
	fmt.Println(r, t1, t2)
	actual := generateNewRanges(r, []Triplet{t1, t2})
	expected := []Range{
		{5, 6, 2},
		{17, 18, 2},
	}
	slices.SortFunc(expected, SortRange)
	slices.SortFunc(actual, SortRange)
	assert.Equal(tt, expected, actual)
}
