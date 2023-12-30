package day10

import (
	"fmt"

	"github.com/jecepeda/advent-of-code-2023/utils"
)

const (
	NorthAndSouth    byte = '|'
	EastAndWest      byte = '-'
	NorthAndEast     byte = 'L'
	NorthAndWest     byte = 'J'
	SouthAndWest     byte = '7'
	SouthAndEast     byte = 'F'
	Ground           byte = '.'
	StartingPosition byte = 'S'
)

type Location [2]int

func (l Location) String() string {
	return fmt.Sprintf("(%d, %d)", l[0], l[1])
}

type Problem struct {
	Map              [][]byte
	StartingLocation *Location
	Visited          map[Location]bool
	MaxX             int
	MaxY             int
}

func NextDirections(previous *Location, current Location, value byte) []Location {
	var locations []Location
	switch value {
	case NorthAndSouth:
		locations = []Location{
			{current[0] - 1, current[1]},
			{current[0] + 1, current[1]},
		}
	case EastAndWest:
		locations = []Location{
			{current[0], current[1] - 1},
			{current[0], current[1] + 1},
		}
	case NorthAndEast:
		locations = []Location{
			{current[0] - 1, current[1]},
			{current[0], current[1] + 1},
		}
	case NorthAndWest:
		locations = []Location{
			{current[0] - 1, current[1]},
			{current[0], current[1] - 1},
		}
	case SouthAndWest:
		locations = []Location{
			{current[0] + 1, current[1]},
			{current[0], current[1] - 1},
		}
	case SouthAndEast:
		locations = []Location{
			{current[0] + 1, current[1]},
			{current[0], current[1] + 1},
		}
	case StartingPosition:
		locations = []Location{
			{current[0] - 1, current[1]},
			{current[0] + 1, current[1]},
			{current[0], current[1] - 1},
			{current[0], current[1] + 1},
		}
	default:
		locations = []Location{}
	}
	var result []Location
	// avoiding going back to the previous location
	for _, l := range locations {
		if previous == nil {
			result = append(result, l)
		} else if l == *previous {
			continue
		} else {
			result = append(result, l)
		}
	}
	return result
}

func (p *Problem) Prepare() {
	p.Visited = map[Location]bool{}
	for i := 0; i < len(p.Map); i++ {
		for j := 0; j < len(p.Map[i]); j++ {
			if p.Map[i][j] == StartingPosition {
				p.StartingLocation = &Location{i, j}
				break
			}
		}
		if p.StartingLocation != nil {
			break
		}
	}
	p.MaxX = len(p.Map)
	p.MaxY = len(p.Map[0])
}

func (p *Problem) Solve1() int {
	route := p.solve(nil, *p.StartingLocation)

	return len(route) / 2
}

func (p *Problem) Solve2() int {
	return 0
}

func (p Problem) solve(previous []Location, loc Location) []Location {
	v := p.Map[loc[0]][loc[1]]
	if len(previous) > 1 && v == StartingPosition {
		// We've found the starting position again, so we're done
		return previous
	}
	var prev *Location
	if len(previous) > 0 {
		prev = &previous[len(previous)-1]
	}
	nextMoves := NextDirections(prev, loc, v)
	for _, nextMove := range nextMoves {
		// fmt.Printf("Going from %v (%c) to %v (%c)\n", loc, v, nextMove, p.Map[nextMove[0]][nextMove[1]])
		if !p.Visited[nextMove] {
			p.Visited[nextMove] = true
			if nextMove[0] < 0 || nextMove[0] >= p.MaxX || nextMove[1] < 0 || nextMove[1] >= p.MaxY {
				continue
			}
			result := p.solve(append(previous, loc), nextMove)
			if result != nil {
				return result
			}
		} else {
			continue
		}
	}
	return nil
}

func Part1(filename string) (int, error) {
	s, err := utils.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	lines := [][]byte{}
	for s.Scan() {
		lines = append(lines, []byte(s.Text()))
	}
	problem := Problem{
		Map: lines,
	}
	problem.Prepare()
	return problem.Solve1(), nil
}
