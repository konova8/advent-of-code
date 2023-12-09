package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertHand(t *testing.T) {
	s := "A23A4"
	expected := Hand{
		'A': 2,
		'2': 1,
		'3': 1,
		'4': 1,
	}
	actual := convertHand(s)
	assert.Equal(t, expected, actual)
}

func TestConvertHand2(t *testing.T) {
	s := "QJJQ2"
	expected := Hand{
		'Q': 4,
		'2': 1,
	}
	actual := convertHand2(s)
	assert.Equal(t, expected, actual)
}

func TestComputeHand(t *testing.T) {
	s := "A23A4"
	expected := One
	h := convertHand(s)
	actual := computeHand(h)
	assert.Equal(t, expected, actual)
}

func TestCompareHand2(t *testing.T) {
	s1 := "QJJQ2"
	s2 := "JKKK2"
	hwb1 := HandWithBid{
		str:  s1,
		hand: convertHand2(s1),
		bid:  1,
	}
	hwb2 := HandWithBid{
		str:  s2,
		hand: convertHand2(s2),
		bid:  2,
	}
	expected := 1
	actual := CompareHand(hwb1, hwb2)
	assert.Equal(t, expected, actual)
}
