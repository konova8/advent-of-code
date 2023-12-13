package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExample(t *testing.T) {
	s := "A23A4"
	expected := "AA23A4"
	actual := "A" + s
	assert.Equal(t, expected, actual)
}
