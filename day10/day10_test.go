package day10

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1Example1(t *testing.T) {
	result, err := Part1("example1.txt")
	require.NoError(t, err)
	require.Equal(t, 4, result)
}

func TestPart1Example2(t *testing.T) {
	result, err := Part1("example2.txt")
	require.NoError(t, err)
	require.Equal(t, 8, result)
}

func TestPart1(t *testing.T) {
	result, err := Part1("input.txt")
	require.NoError(t, err)
	require.Equal(t, 6867, result)
}

// Part 2 requires more research
