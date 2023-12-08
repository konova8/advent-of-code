package util_test

import (
	"testing"

	"github.com/konova8/advent-of-code/util"
)

func TestStrReverse(t *testing.T) {
	s := "ABCD"
	actual := util.StrReverse(s)
	expected := "DCBA"
	if actual != expected {
		t.Fatalf(`Original = %s, should be %s, obtained %s`, s, expected, actual)
	}
}

func TestGCD(t *testing.T) {
	n := []int{6, 7, 10, 15}
	actual := util.GCD(n...)
	expected := 1
	if actual != expected {
		t.Fatalf(`Should be %d, obtained %d`, expected, actual)
	}
}

func TestGCD2(t *testing.T) {
	n := []int{}
	actual := util.GCD(n...)
	expected := -1
	if actual != expected {
		t.Fatalf(`Should be %d, obtained %d`, expected, actual)
	}
}

func TestLCM(t *testing.T) {
	n := []int{6, 7, 10, 15, 11}
	actual := util.LCM(n...)
	expected := 2 * 3 * 5 * 7 * 11
	if actual != expected {
		t.Fatalf(`Should be %d, obtained %d`, expected, actual)
	}
}

func TestLCM2(t *testing.T) {
	n := []int{}
	actual := util.LCM(n...)
	expected := -1
	if actual != expected {
		t.Fatalf(`Should be %d, obtained %d`, expected, actual)
	}
}
