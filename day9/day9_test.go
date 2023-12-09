package day9

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1Example(t *testing.T) {
	result, err := Part1("example.txt")
	require.NoError(t, err)
	require.Equal(t, 114, result)
}

func TestPart1(t *testing.T) {
	result, err := Part1("input.txt")
	require.NoError(t, err)
	require.Equal(t, 1637452029, result)
}

func TestPart2Example(t *testing.T) {
	result, err := Part2("example.txt")
	require.NoError(t, err)
	require.Equal(t, 2, result)
}

func TestPart2(t *testing.T) {
	result, err := Part2("input.txt")
	require.NoError(t, err)
	require.Equal(t, 908, result)
}
