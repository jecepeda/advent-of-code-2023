package day11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExamplePart1(t *testing.T) {
	result, err := Part1("example.txt")

	assert.NoError(t, err)
	assert.Equal(t, 374, result)
}

func TestPart1(t *testing.T) {
	result, err := Part1("input.txt")

	assert.NoError(t, err)
	assert.Equal(t, 10165598, result)
}

func TestPart2(t *testing.T) {
	result, err := Part2("input.txt")

	assert.NoError(t, err)
	assert.Equal(t, 678728808158, result)
}
