package day6

import (
	"strings"

	"github.com/jecepeda/advent-of-code-2023/utils"
)

type Equation struct {
	RaceTime int
	Distance int
}

func solve(equations []Equation) int {
	result := 1
	var wins int
	for _, e := range equations {
		wins = 0
		for buttonTime := 1; buttonTime <= e.RaceTime; buttonTime++ {
			remainingTime := e.RaceTime - buttonTime
			distance := buttonTime * remainingTime
			if distance > e.Distance {
				wins++
			}
		}
		result *= wins
	}
	return result
}

func Part1(filename string) (int, error) {
	equations, err := ReadFile(filename)
	if err != nil {
		return 0, err
	}
	return solve(equations), nil
}

func Part2(filename string) (int, error) {
	eq, err := Read2(filename)

	if err != nil {
		return 0, err
	}
	return solve([]Equation{eq}), nil
}

func Read2(filename string) (Equation, error) {
	s, err := utils.ReadFile(filename)
	if err != nil {
		return Equation{}, err
	}
	eq := Equation{}
	for s.Scan() {
		t := s.Text()
		if strings.Contains(t, "Time:") {
			t = strings.Replace(t, "Time: ", "", -1)
			t = strings.ReplaceAll(t, " ", "")
			lines, err := utils.ReadNumbers(t)
			if err != nil {
				return Equation{}, err
			}
			eq.RaceTime = lines[0]
		} else {
			t = strings.Replace(t, "Distance: ", "", -1)
			t = strings.ReplaceAll(t, " ", "")
			lines, err := utils.ReadNumbers(t)
			if err != nil {
				return Equation{}, err
			}
			eq.Distance = lines[0]
		}
	}
	return eq, nil
}

func ReadFile(filename string) ([]Equation, error) {
	s, err := utils.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var equations []Equation
	for s.Scan() {
		t := s.Text()
		if strings.Contains(t, "Time:") {
			t = strings.Replace(t, "Time: ", "", -1)
			lines, err := utils.ReadNumbers(t)
			if err != nil {
				return nil, err
			}
			equations = make([]Equation, len(lines))
			for i, l := range lines {
				equations[i].RaceTime = l
			}
		} else {
			t = strings.Replace(t, "Distance: ", "", -1)
			lines, err := utils.ReadNumbers(t)
			if err != nil {
				return nil, err
			}
			for i, l := range lines {
				equations[i].Distance = l
			}
		}
	}
	return equations, nil
}
