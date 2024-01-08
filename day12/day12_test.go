package day12

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblem_Solve(t *testing.T) {
	type fields struct {
		input        string
		arrangements []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "example",
			fields: fields{
				input:        "???.###",
				arrangements: []int{1, 1, 3},
			},
			want: 1,
		},
		{
			name: "example 2",
			fields: fields{
				input:        ".??..??...?##.",
				arrangements: []int{1, 1, 3},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewProblem(tt.fields.input, tt.fields.arrangements)
			require.Equal(t, tt.want, p.Solve())
		})
	}
}

func TestProblemUnfold(t *testing.T) {
	type fields struct {
		input        string
		arrangements []int
	}

	tests := []struct {
		name   string
		fields fields
		want   *Problem
	}{
		{
			name: "example",
			fields: fields{
				input:        ".#",
				arrangements: []int{1},
			},
			want: &Problem{
				input:        ".#?.#",
				arrangements: []int{1, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewProblem(tt.fields.input, tt.fields.arrangements)
			p.Unfold(2)
			require.Equal(t, tt.want.input, p.input)
			require.Equal(t, tt.want.arrangements, p.arrangements)
		})
	}
}

func TestPart1Example(t *testing.T) {
	result, err := Part1("example.txt")
	require.NoError(t, err)
	require.Equal(t, 21, result)
}

func TestPart1Input(t *testing.T) {
	result, err := Part1("input.txt")
	require.NoError(t, err)
	require.Equal(t, 7084, result)
}

func TestPart2Example(t *testing.T) {
	result, err := Part2("example.txt")
	require.NoError(t, err)
	require.Equal(t, 525152, result)
}

func TestPart2Input(t *testing.T) {
	result, err := Part2("input.txt")
	require.NoError(t, err)
	require.Equal(t, 8414003326821, result)
}
