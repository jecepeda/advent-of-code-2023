package day5

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/jecepeda/advent-of-code-2023/utils"
)

type Almanac struct {
	Seeds      []int
	SeedRanges []SeedRange
	Steps      []Step
}

type SeedRange struct {
	SourceStart int
	SourceEnd   int
}

type Step struct {
	// <source> -> <destination>
	Source      string
	Destination string
	Rules       Rules
}

func (s Step) String() string {
	return fmt.Sprintf("source=%s destination=%s", s.Source, s.Destination)
}

type Rules []Rule

func (p Rules) Len() int {
	return len(p)
}

func (p Rules) Less(i, j int) bool {
	return p[i].SourceStart < p[j].SourceStart
}

func (p Rules) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (s *Step) ReOrder() {
	sort.Sort(s.Rules)
}

func (s Step) Apply(value int) int {
	for _, rule := range s.Rules {
		if value < rule.SourceStart || value > (rule.SourceStart+(rule.Range-1)) {
			continue
		}
		return rule.DestinationStart + (value - rule.SourceStart)
	}
	return -1
}

type Rule struct {
	// <source> <destination> <range>
	// 50 100 10 would turn into
	// [50-60] [100-100]
	SourceStart      int
	DestinationStart int
	Range            int
}

func (r Rule) String() string {
	return fmt.Sprintf("start=%d destination=%d range=%d", r.SourceStart, r.DestinationStart, r.Range)
}

func Part1(filename string) (int, error) {
	almanac, err := ParseFile(filename)
	if err != nil {
		return 0, err
	}
	result := math.MaxInt
	for _, seed := range almanac.Seeds {
		partial := seed
		for _, step := range almanac.Steps {
			newPartial := step.Apply(partial)
			if newPartial != -1 {
				partial = newPartial
			}
		}
		if partial < result {
			result = partial
		}
	}
	return result, nil
}

func Part2(filename string) (int, error) {
	almanac, err := ParseFile(filename)
	if err != nil {
		return 0, err
	}
	result := math.MaxInt
	var partial, newPartial int
	for _, seedRange := range almanac.SeedRanges {
		for i := seedRange.SourceStart; i < seedRange.SourceEnd; i++ {
			partial = i
			for _, step := range almanac.Steps {
				newPartial = step.Apply(partial)
				if newPartial != -1 {
					partial = newPartial
				}
			}
			if partial < result {
				result = partial
			}
		}
	}
	return result, nil
}

func ParseFile(filename string) (*Almanac, error) {
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
			almanac.Seeds = seeds
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
			step.Rules = append(
				step.Rules,
				Rule{
					SourceStart:      numbers[1],
					DestinationStart: numbers[0],
					Range:            numbers[2], // -1 because the range is inclusive. 50 + 2 means 50-51
				},
			)
		}
	}
	almanac.Steps = append(almanac.Steps, *step)

	for i := 0; i < len(almanac.Steps); i++ {
		almanac.Steps[i].ReOrder()
	}

	for i := 0; i < len(almanac.Seeds); i += 2 {
		almanac.SeedRanges = append(almanac.SeedRanges, SeedRange{
			SourceStart: almanac.Seeds[i],
			SourceEnd:   almanac.Seeds[i] + almanac.Seeds[i+1],
		})
	}

	return &almanac, nil
}

func getSourceAndDestination(line string) (string, string) {
	line = strings.Replace(line, " map:", "", 1)
	parts := strings.Split(line, "-to-")
	return parts[0], parts[1]
}
