package day5

import (
	"sort"
	"strings"

	"github.com/jecepeda/advent-of-code-2023/utils"
)

type Almanac struct {
	Seeds []Range
	Steps []Step
}

func (a *Almanac) Solve() int {
	projections := []Range{}
	for _, s := range a.Seeds {
		seed := []Range{s}
		for _, step := range a.Steps {
			seed = step.Apply(seed)
		}
		projections = append(projections, seed...)
	}
	sort.Slice(projections, func(i, j int) bool {
		return projections[i][0] < projections[j][0]
	})
	return projections[0][0]
}

type Range [2]int

type Step struct {
	Source      string
	Destination string
	Rules       []Rule
}

func getOverlapping(a, b Range) *Range {
	left, right := max(a[0], b[0]), min(a[1], b[1])
	if left <= right {
		return &Range{left, right}
	}
	return nil
}

func getMapping(seed Range, rule Rule) Range {
	offset := rule.Destination[0] - rule.Source[0]

	return Range{seed[0] + offset, seed[1] + offset}
}

func (s Step) Apply(seeds []Range) []Range {
	result := []Range{}
	for _, seed := range seeds {
		partial := []Range{}
		for _, rule := range s.Rules {
			overlapping := getOverlapping(seed, rule.Source)
			if overlapping != nil {
				l, r := overlapping[0], overlapping[1]
				mapping := getMapping(Range{l, r}, rule)
				partial = append(partial, mapping)
				if seed[0] < l {
					partial = append(partial, Range{seed[0], l})
				}
				seed = Range{r, seed[1]}
			}
		}
		if seed[0] < seed[1] || len(partial) == 0 {
			partial = append(partial, seed)
		}
		result = append(result, partial...)
	}
	return result
}

type Rule struct {
	Source      Range
	Destination Range
}

func (s *Step) ReOrder() {
	sort.Slice(s.Rules, func(i, j int) bool {
		return s.Rules[i].Source[0] < s.Rules[j].Source[0]
	})
}

func Part1(filename string) (int, error) {
	almanac, err := ParseFile(filename, false)
	if err != nil {
		return 0, err
	}
	return almanac.Solve(), nil
}

func Part2(filename string) (int, error) {
	almanac, err := ParseFile(filename, true)
	if err != nil {
		return 0, err
	}
	return almanac.Solve(), nil
}

func ParseFile(filename string, isRange bool) (*Almanac, error) {
	almanac := Almanac{}

	s, err := utils.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var step *Step = nil
	for s.Scan() {
		line := s.Text()
		if strings.Contains(line, "seeds") {
			rawSeeds := strings.Replace(line, "seeds: ", "", 1)
			seeds, err := utils.ReadNumbers(rawSeeds)
			if err != nil {
				return nil, err
			}
			if isRange {
				for i := 0; i < len(seeds); i += 2 {
					almanac.Seeds = append(almanac.Seeds, Range{seeds[i], seeds[i] + seeds[i+1] - 1})
				}
			} else {
				for i := 0; i < len(seeds); i++ {
					almanac.Seeds = append(almanac.Seeds, Range{seeds[i], seeds[i]})
				}
			}
		} else if strings.Contains(line, "map:") {
			if step == nil {
				step = &Step{}
				step.Source, step.Destination = getSourceAndDestination(line)
			} else {
				almanac.Steps = append(almanac.Steps, *step)
				step = &Step{}
				step.Source, step.Destination = getSourceAndDestination(line)
			}
		} else if line != "" {
			numbers, err := utils.ReadNumbers(line)
			if err != nil {
				return nil, err
			}
			dst, src, rng := numbers[0], numbers[1], numbers[2]
			step.Rules = append(
				step.Rules,
				Rule{
					Source:      Range{src, src + rng - 1},
					Destination: Range{dst, dst + rng - 1},
				},
			)
		}
	}
	almanac.Steps = append(almanac.Steps, *step)

	for i := 0; i < len(almanac.Steps); i++ {
		almanac.Steps[i].ReOrder()
	}

	return &almanac, nil
}

func getSourceAndDestination(line string) (string, string) {
	line = strings.Replace(line, " map:", "", 1)
	parts := strings.Split(line, "-to-")
	return parts[0], parts[1]
}
