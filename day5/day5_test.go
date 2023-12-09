package day5

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1Example(t *testing.T) {
	result, err := Part1("example.txt")
	require.NoError(t, err)
	require.Equal(t, 35, result)
}

func TestPart1(t *testing.T) {
	result, err := Part1("input.txt")
	require.NoError(t, err)
	require.Equal(t, 579439039, result)
}

func TestPart2Example(t *testing.T) {
	result, err := Part2("example.txt")
	require.NoError(t, err)
	require.Equal(t, 46, result)
}

func TestPart2(t *testing.T) {
	result, err := Part2("input.txt")
	require.NoError(t, err)
	require.Equal(t, 7873084, result)
}

func TestSolve(t *testing.T) {
	almanac := Almanac{
		Steps: []Step{
			{
				Source:      "A",
				Destination: "B",
				// 50 98 2
				// 52 50 48
				Rules: []Rule{
					{
						Source:      Range{98, 99},
						Destination: Range{50, 51},
					},
					{
						Source:      Range{50, 97},
						Destination: Range{52, 100},
					},
				},
			},
		},
	}
	testCases := []struct {
		SeedRange Range
		Expected  int
	}{
		// 79 14 55 13
		{
			SeedRange: Range{79, 79},
			Expected:  81,
		},
		{
			SeedRange: Range{14, 14},
			Expected:  14,
		},
		{
			SeedRange: Range{55, 55},
			Expected:  57,
		},
		{
			SeedRange: Range{13, 13},
			Expected:  13,
		},
	}
	for _, testCase := range testCases {
		almanac.Seeds = []Range{testCase.SeedRange}
		result := almanac.Solve()
		require.Equal(t, testCase.Expected, result, "seed range %s", testCase.SeedRange)
	}
}
