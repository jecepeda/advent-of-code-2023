package day11

import "github.com/jecepeda/advent-of-code-2023/utils"

type Problem struct {
	increment   int
	lines       [][]byte
	horizontals []int
	verticals   []int
	galaxies    [][2]int
}

func NewProblem(lines [][]byte, increment int) *Problem {
	return &Problem{
		lines:     lines,
		increment: increment,
	}
}

func (p *Problem) Process() {
	p.SaveEmptyLines()
	p.SaveGalaxies()
}

func (p *Problem) SaveEmptyLines() {
	horizontals := []int{}
	verticals := []int{}

	for i := range p.lines {
		galaxy := false
		for j := range p.lines[i] {
			if p.lines[i][j] == '#' {
				galaxy = true
				break
			}
		}
		if !galaxy {
			horizontals = append(horizontals, i)
		}
	}
	for i := range p.lines[0] {
		galaxy := false
		for j := range p.lines {
			if p.lines[j][i] == '#' {
				galaxy = true
				break
			}
		}
		if !galaxy {
			verticals = append(verticals, i)
		}
	}
	p.horizontals = horizontals
	p.verticals = verticals
}

func (p *Problem) SaveGalaxies() {
	galaxies := [][2]int{}
	for i := range p.lines {
		for j := range p.lines[i] {
			if p.lines[i][j] == '#' {
				galaxies = append(galaxies, [2]int{i, j})
			}
		}
	}
	p.galaxies = galaxies
}

func (p *Problem) getShortestPathSum() int {
	sum := 0

	for i := 0; i < len(p.galaxies); i++ {
		for j := i + 1; j < len(p.galaxies); j++ {
			a, b := p.galaxies[i], p.galaxies[j]
			sum += distance(a[0], b[0], p.horizontals, p.increment)
			sum += distance(a[1], b[1], p.verticals, p.increment)
		}
	}

	return sum
}

func distance(a, b int, between []int, amount int) int {
	distance := 0
	for i := range between {
		if (between[i] > a && between[i] < b) || (between[i] < a && between[i] > b) {
			distance += (amount - 1)
		}
	}
	return utils.Abs(a-b) + distance
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

	problem := NewProblem(lines, 2)
	problem.Process()

	return problem.getShortestPathSum(), nil
}

func Part2(filename string) (int, error) {
	s, err := utils.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	lines := [][]byte{}
	for s.Scan() {
		lines = append(lines, []byte(s.Text()))
	}

	problem := NewProblem(lines, 1_000_000)
	problem.Process()

	return problem.getShortestPathSum(), nil
}
