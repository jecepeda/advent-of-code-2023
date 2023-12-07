package day2

import (
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/jecepeda/advent-of-code-2023/utils"
)

type Game struct {
	ID          int
	BallSubsets []BallSubset
}

func (g Game) IsValid(maxReds, maxGreens, maxBluees int) bool {
	for _, subset := range g.BallSubsets {
		if subset.Reds > maxReds || subset.Greens > maxGreens || subset.Blues > maxBluees {
			return false
		}
	}
	return true
}

func (g Game) Power() int {
	minReds := math.MinInt
	minBlues := math.MinInt
	minGreens := math.MinInt
	for _, subset := range g.BallSubsets {
		if subset.Reds > minReds {
			minReds = subset.Reds
		}
		if subset.Greens > minGreens {
			minGreens = subset.Greens
		}
		if subset.Blues > minBlues {
			minBlues = subset.Blues
		}
	}
	return minReds * minGreens * minBlues
}

type BallSubset struct {
	Reds   int
	Greens int
	Blues  int
}

func Part1(filename string) (int, error) {
	r, err := utils.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	result := 0
	for r.Scan() {
		game, err := readGame(r.Text())
		if err != nil {
			return 0, err
		}
		if game.IsValid(12, 13, 14) {
			result += game.ID
		}
	}
	return result, nil
}

func Part2(filename string) (int, error) {
	r, err := utils.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	result := 0
	for r.Scan() {
		game, err := readGame(r.Text())
		if err != nil {
			return 0, err
		}
		result += game.Power()
	}
	return result, nil
}

var subsetRegex = regexp.MustCompile(`(\d+) (blue|red|green)`)

func readGame(line string) (Game, error) {
	var err error
	g := Game{}
	splitted := strings.Split(line, ":")
	g.ID, err = strconv.Atoi(strings.Replace(splitted[0], "Game ", "", 1))
	if err != nil {
		return g, err
	}
	for _, strSubset := range strings.Split(splitted[1], ";") {
		subset := BallSubset{}
		for _, match := range subsetRegex.FindAllStringSubmatch(strSubset, -1) {
			switch match[2] {
			case "blue":
				subset.Blues, err = strconv.Atoi(match[1])
				if err != nil {
					return g, err
				}
			case "red":
				subset.Reds, err = strconv.Atoi(match[1])
				if err != nil {
					return g, err
				}
			case "green":
				subset.Greens, err = strconv.Atoi(match[1])
				if err != nil {
					return g, err
				}
			}
			if err != nil {
				return g, err
			}
		}
		g.BallSubsets = append(g.BallSubsets, subset)
	}

	return g, nil
}
