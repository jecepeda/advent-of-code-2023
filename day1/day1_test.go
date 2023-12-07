package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	result, err := Part1("input.txt")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 55971, result)
}

func TestPart2Example(t *testing.T) {
	result, err := Part2("example_part_2.txt")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 281, result)
}

func TestPart2(t *testing.T) {
	result, err := Part2("input.txt")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 54719, result)
}
