package day5

import (
	"sort"
	"strings"

	"github.com/jecepeda/advent-of-code-2023/utils"
)

type Almanac struct {
	Seeds SeedRanges
	Steps []Step
}

func (a *Almanac) Solve() int {
	projections := SeedRanges{}
	for _, s := range a.Seeds {
		seed := SeedRanges{s}
		for _, step := range a.Steps {
			seed = step.Apply(seed)
		}
		projections = append(projections, seed...)
	}
	sort.Sort(projections)
	return projections[0][0]
}

type SeedRange [2]int

type SeedRanges []SeedRange

func (p SeedRanges) Len() int {
	return len(p)
}

func (p SeedRanges) Less(i, j int) bool {
	return p[i][0] < p[j][0]
}

func (p SeedRanges) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type Step struct {
	Source      string
	Destination string
	Rules       Rules
}

func getOverlapping(a, b [2]int) *[2]int {
	left, right := max(a[0], b[0]), min(a[1], b[1])
	if left <= right {
		return &[2]int{left, right}
	}
	return nil
}

func getMapping(seed SeedRange, rule Rule) SeedRange {
	offset := rule.Destination[0] - rule.Source[0]

	return SeedRange{seed[0] + offset, seed[1] + offset}
}

func (s Step) Apply(seeds SeedRanges) SeedRanges {
	result := SeedRanges{}
	for _, seed := range seeds {
		partial := SeedRanges{}
		for _, rule := range s.Rules {
			overlapping := getOverlapping(seed, rule.Source)
			if overlapping != nil {
				l, r := overlapping[0], overlapping[1]
				mapping := getMapping(SeedRange{l, r}, rule)
				partial = append(partial, mapping)
				if seed[0] < l {
					partial = append(partial, SeedRange{seed[0], l})
				}
				seed = SeedRange{r, seed[1]}
			}
		}
		if seed[0] < seed[1] || len(partial) == 0 {
			partial = append(partial, seed)
		}
		result = append(result, partial...)
	}
	return result
}

type Rules []Rule

func (p Rules) Len() int {
	return len(p)
}

func (p Rules) Less(i, j int) bool {
	return p[i].Source[0] < p[j].Source[0]
}

func (p Rules) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type Rule struct {
	Source      [2]int
	Destination [2]int
}

func (s *Step) ReOrder() {
	sort.Sort(s.Rules)
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
					almanac.Seeds = append(almanac.Seeds, SeedRange{seeds[i], seeds[i] + seeds[i+1] - 1})
				}
			} else {
				for i := 0; i < len(seeds); i++ {
					almanac.Seeds = append(almanac.Seeds, SeedRange{seeds[i], seeds[i]})
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
					Source:      [2]int{src, src + rng - 1},
					Destination: [2]int{dst, dst + rng - 1},
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
