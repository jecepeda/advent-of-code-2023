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
						Source:      [2]int{98, 99},
						Destination: [2]int{50, 51},
					},
					{
						Source:      [2]int{50, 97},
						Destination: [2]int{52, 100},
					},
				},
			},
		},
	}
	testCases := []struct {
		SeedRange [2]int
		Expected  int
	}{
		// 79 14 55 13
		{
			SeedRange: [2]int{79, 79},
			Expected:  81,
		},
		{
			SeedRange: [2]int{14, 14},
			Expected:  14,
		},
		{
			SeedRange: [2]int{55, 55},
			Expected:  57,
		},
		{
			SeedRange: [2]int{13, 13},
			Expected:  13,
		},
	}
	for _, testCase := range testCases {
		almanac.Seeds = SeedRanges{testCase.SeedRange}
		result := almanac.Solve()
		require.Equal(t, testCase.Expected, result, "seed range %s", testCase.SeedRange)
	}
}
