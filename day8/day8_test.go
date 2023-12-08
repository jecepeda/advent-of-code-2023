package day8

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1Example(t *testing.T) {
	result, err := Part1("example.txt")
	require.NoError(t, err)
	require.Equal(t, 2, result)
}

func TestPart1Example2(t *testing.T) {
	result, err := Part1("example2.txt")
	require.NoError(t, err)
	require.Equal(t, 6, result)
}

func TestPart1(t *testing.T) {
	result, err := Part1("input.txt")
	require.NoError(t, err)
	require.Equal(t, 11911, result)
}

func TestPart2Example3(t *testing.T) {
	result, err := Part2("example3.txt")
	require.NoError(t, err)
	require.Equal(t, 6, result)
}

func TestPart2(t *testing.T) {
	result, err := Part2("input.txt")
	require.NoError(t, err)
	require.Equal(t, 10151663816849, result)
}
