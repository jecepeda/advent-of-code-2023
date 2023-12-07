package day5

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseFile(t *testing.T) {
	almanac, err := ParseFile("example.txt")
	require.NoError(t, err)
	require.Equal(t, []int{79, 14, 55, 13}, almanac.Seeds)
}

func TestParseFileRealExample(t *testing.T) {
	almanac, err := ParseFile("input.txt")
	require.NoError(t, err)
	require.NotNil(t, almanac)
}

func TestStep_Apply(t *testing.T) {
	type fields struct {
		Source      string
		Destination string
		Rules       []Rule
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "1",
			fields: fields{
				Source:      "A",
				Destination: "B",
				Rules: []Rule{
					{
						SourceStart:      98,
						DestinationStart: 50,
						Range:            2,
					},
				},
			},
			args: args{value: 98},
			want: 50,
		},
		{
			name: "second",
			fields: fields{
				Source:      "A",
				Destination: "B",
				Rules: []Rule{
					{
						SourceStart:      98,
						DestinationStart: 50,
						Range:            2,
					},
				},
			},
			args: args{value: 99},
			want: 51,
		},
		{
			name: "3",
			fields: fields{
				Source:      "A",
				Destination: "B",
				Rules: []Rule{
					{
						SourceStart:      98,
						DestinationStart: 50,
						Range:            2,
					},
					{
						SourceStart:      100,
						DestinationStart: 50,
						Range:            1,
					},
				},
			},
			args: args{value: 100},
			want: 50,
		},
		{
			name: "error case",
			fields: fields{
				Source:      "A",
				Destination: "B",
				Rules: []Rule{
					{
						SourceStart:      98,
						DestinationStart: 50,
						Range:            2,
					},
					{
						SourceStart:      100,
						DestinationStart: 50,
						Range:            1,
					},
				},
			},
			args: args{value: 97},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Step{
				Source:      tt.fields.Source,
				Destination: tt.fields.Destination,
				Rules:       tt.fields.Rules,
			}
			got := s.Apply(tt.args.value)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestPart1Example(t *testing.T) {
	result, err := Part1("example.txt")
	require.NoError(t, err)
	require.Equal(t, 35, result)
}

func TestPart2Example(t *testing.T) {
	result, err := Part2("example.txt")
	require.NoError(t, err)
	require.Equal(t, 46, result)
}

func TestPart1(t *testing.T) {
	result, err := Part1("input.txt")
	require.NoError(t, err)
	require.Equal(t, 579439039, result)
}

func TestPart2(t *testing.T) {
	t.Skip("This test takes too long to run. Fix later")
	result, err := Part2("input.txt")
	require.NoError(t, err)
	require.Equal(t, 35, result)
}
