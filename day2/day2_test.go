package day2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExample1(t *testing.T) {
	res, err := Part1("example_1.txt")
	require.NoError(t, err)
	require.Equal(t, 8, res)
}

func TestExample2(t *testing.T) {
	res, err := Part2("example_1.txt")
	require.NoError(t, err)
	require.Equal(t, 2286, res)
}

func TestPart1(t *testing.T) {
	res, err := Part1("input.txt")
	require.NoError(t, err)
	require.Equal(t, 2278, res)
}

func TestPart2(t *testing.T) {
	res, err := Part2("input.txt")
	require.NoError(t, err)
	require.Equal(t, 67953, res)
}

func TestGameCreation(t *testing.T) {
	game, err := readGame("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
	require.NoError(t, err)
	expectedGame := Game{
		ID: 1,
		BallSubsets: []BallSubset{
			{Blues: 3, Reds: 4},
			{Reds: 1, Greens: 2, Blues: 6},
			{Greens: 2},
		},
	}
	require.Equal(t, expectedGame, game)
}
