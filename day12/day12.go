package day12

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jecepeda/advent-of-code-2023/utils"
)

const (
	Operational byte = '.'
	Damaged     byte = '#'
	Unknown     byte = '?'
)

type Problem struct {
	input        string
	arrangements []int
	cache        map[[2]int]int
}

func NewProblem(input string, arrangements []int) *Problem {
	p := &Problem{
		input:        input,
		arrangements: arrangements,
		cache:        make(map[[2]int]int, 1_000), // a bit of size tuning to avoid reallocations
	}
	return p
}

func (p *Problem) Unfold(times int) {
	var newInput strings.Builder
	for i := 0; i < times; i++ {
		newInput.WriteString(p.input)
		if i < times-1 {
			newInput.WriteByte(Unknown)
		}
	}
	p.input = newInput.String()
	newArrangements := make([]int, len(p.arrangements)*times)
	for i := 0; i < times; i++ {
		for j := 0; j < len(p.arrangements); j++ {
			newArrangements[i*len(p.arrangements)+j] = p.arrangements[j]
		}
	}
	p.arrangements = newArrangements
}

func (p *Problem) Solve() int {
	return p.solve(0, 0)
}

func (p *Problem) solve(i, j int) int {
	if i >= len(p.input) {
		if j < len(p.arrangements) {
			return 0
		}
		return 1
	}
	if v, ok := p.cache[[2]int{i, j}]; ok {
		return v
	}
	var result int
	if p.input[i] == Operational {
		result = p.solve(i+1, j)
	} else {
		if p.input[i] == Unknown {
			result = p.solve(i+1, j)
		}
		if j < len(p.arrangements) {
			var count int
			for x := i; x < len(p.input); x++ {
				if p.input[x] == Operational {
					break
				} else if count > p.arrangements[j] {
					break
				} else if count == p.arrangements[j] && p.input[x] == Unknown {
					break
				}
				count++
			}
			if count == p.arrangements[j] {
				if i+count < len(p.input) && p.input[i+count] != Damaged {
					// if not damaged, means that it's either operational or unknown
					// so we can skip the next one
					// as there should be a difference of at least 1
					result += p.solve(i+count+1, j+1)
				} else {
					result += p.solve(i+count, j+1)
				}
			}
		}
	}

	p.cache[[2]int{i, j}] = result
	return result
}

func Part1(filename string) (int, error) {
	s, err := utils.ReadFile(filename)
	if err != nil {
		return -1, err
	}
	var result int
	for s.Scan() {
		var p *Problem
		p, err = readLine(s.Text())
		if err != nil {
			return -1, err
		}
		result += p.Solve()
	}
	if s.Err() != nil {
		return -1, err
	}

	return result, nil
}

func Part2(filename string) (int, error) {
	s, err := utils.ReadFile(filename)
	if err != nil {
		return -1, err
	}
	var result int
	for s.Scan() {
		var p *Problem
		p, err = readLine(s.Text())
		p.Unfold(5)
		if err != nil {
			return -1, err
		}
		result += p.Solve()
	}
	if s.Err() != nil {
		return -1, err
	}

	return result, nil
}

func readLine(line string) (*Problem, error) {
	var pInput string
	var arrangements []int
	split := strings.Split(line, " ")
	if len(split) != 2 {
		return nil, fmt.Errorf("invalid input: %q", line)
	}
	pInput = split[0]
	for _, s := range strings.Split(split[1], ",") {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		arrangements = append(arrangements, n)
	}
	return NewProblem(pInput, arrangements), nil
}
