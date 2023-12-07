package day3

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1Example(t *testing.T) {
	result, err := Part1("example.txt")
	require.NoError(t, err)
	require.Equal(t, 4361, result)
}

func TestPart1(t *testing.T) {
	fmt.Println("Part 1")

	result, err := Part1("input.txt")
	require.NoError(t, err)
	require.Equal(t, 532445, result)
}

func TestPart2(t *testing.T) {
	result, err := Part2("input.txt")
	require.NoError(t, err)
	require.Equal(t, 79842967, result)
}

func TestPart2Example(t *testing.T) {
	result, err := Part2("example.txt")
	require.NoError(t, err)
	require.Equal(t, 467835, result)
}
