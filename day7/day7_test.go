package day7

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPriority(t *testing.T) {
}

func TestHand_InferType(t *testing.T) {
	tests := []struct {
		name         string
		Cards        string
		expectedType HandType
	}{
		{
			name:         "HighCard",
			Cards:        "A2345",
			expectedType: HighCard,
		},
		{
			name:         "OnePair",
			Cards:        "AA345",
			expectedType: OnePair,
		},
		{
			name:         "TwoPairs",
			Cards:        "AA335",
			expectedType: TwoPairs,
		},
		{
			name:         "ThreeOfAKind",
			Cards:        "AAA35",
			expectedType: ThreeOfAKind,
		},
		{
			name:         "FullHouse",
			Cards:        "AAA33",
			expectedType: FullHouse,
		},
		{
			name:         "FourOfAKind",
			Cards:        "AAAA3",
			expectedType: FourOfAKind,
		},
		{
			name:         "FiveOfAKind",
			Cards:        "AAAAA",
			expectedType: FiveOfAKind,
		},
		{
			name:         "test",
			Cards:        "22223",
			expectedType: FourOfAKind,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Hand{
				Cards: tt.Cards,
			}
			h.InferType()
			require.Equal(t, tt.expectedType, h.Type)
		})
	}
}

func TestOrdering(t *testing.T) {
	testCases := []struct {
		name     string
		hands    []Hand
		expected []Hand
	}{
		{
			name: "ThreeOfAKind",
			hands: []Hand{
				{
					Cards: "T55J5",
					Type:  ThreeOfAKind,
				},
				{
					Cards: "QQQJA",
					Type:  ThreeOfAKind,
				},
			},
			expected: []Hand{
				{
					Cards: "T55J5",
					Type:  ThreeOfAKind,
				},
				{
					Cards: "QQQJA",
					Type:  ThreeOfAKind,
				},
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			sort.SliceStable(tt.hands, func(i, j int) bool {
				return HandOrder1(tt.hands[i], tt.hands[j])
			})
			require.Equal(t, tt.expected, tt.hands)
		})
	}
}

func TestExamplePart1(t *testing.T) {
	result, err := Part1("example.txt")
	require.NoError(t, err)
	require.Equal(t, 6440, result)
}

func TestPart1(t *testing.T) {
	result, err := Part1("input.txt")
	require.NoError(t, err)
	require.Equal(t, 249483956, result)
}

func TestExamplePart2(t *testing.T) {
	result, err := Part2("example.txt")
	require.NoError(t, err)
	require.Equal(t, 5905, result)
}

func TestPart2(t *testing.T) {
	result, err := Part2("input.txt")
	require.NoError(t, err)
	require.Equal(t, 252137472, result)
}
