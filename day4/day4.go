package day4

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/jecepeda/advent-of-code-2023/utils"
)

var cardRegex = regexp.MustCompile(`Card\s+(\d+):`)

type Game struct {
	Cards []ScratchCard
}

func (g Game) Day1() int {
	result := 0
	for _, card := range g.Cards {
		result += card.Score()
	}
	return result
}

func (g Game) Day2() int {
	copies := make(map[int]int)
	// initialization, at least 1 card for ID
	for _, card := range g.Cards {
		copies[card.ID] = 1
	}

	for _, card := range g.Cards {
		matches := card.MatchingNumbers()
		for i := card.ID + 1; i <= card.ID+matches; i++ {
			copies[i] += copies[card.ID]
		}
	}

	result := 0
	for _, v := range copies {
		result += v
	}
	return result
}

type ScratchCard struct {
	ID             int
	WinningNumbers map[int]int
	ScratchNumbers []int
}

func NewScratchCard() ScratchCard {
	return ScratchCard{
		WinningNumbers: make(map[int]int),
	}
}

func (s ScratchCard) MatchingNumbers() int {
	result := 0
	ok := false
	for _, num := range s.ScratchNumbers {
		if _, ok = s.WinningNumbers[num]; ok {
			result++
		}

	}
	return result
}

func (s ScratchCard) Score() int {
	result := 0
	for _, num := range s.ScratchNumbers {
		_, ok := s.WinningNumbers[num]
		if ok && result == 0 {
			result = 1
		} else if ok {
			result *= 2
		}
	}

	return result
}

func Part1(filename string) (int, error) {
	s, err := utils.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	game := Game{}
	for s.Scan() {
		card, err := readScratchCard(s.Text())
		if err != nil {
			return 0, err
		}
		game.Cards = append(game.Cards, card)
	}

	return game.Day1(), nil
}

func Part2(filename string) (int, error) {
	s, err := utils.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	game := Game{}
	for s.Scan() {
		card, err := readScratchCard(s.Text())
		if err != nil {
			return 0, err
		}
		game.Cards = append(game.Cards, card)
	}

	return game.Day2(), nil
}

func readScratchCard(line string) (ScratchCard, error) {
	card := NewScratchCard()
	var err error

	cardMatch := cardRegex.FindStringSubmatch(line)
	card.ID, err = strconv.Atoi(cardMatch[1])
	if err != nil {
		return ScratchCard{}, err
	}

	line = cardRegex.ReplaceAllLiteralString(line, "")
	numbers := strings.Split(line, "|")

	for i, strNum := range strings.Split(numbers[0], " ") {
		if strNum == "" {
			continue
		}
		num, err := strconv.Atoi(strNum)
		if err != nil {
			return ScratchCard{}, err
		}
		card.WinningNumbers[num] = i
	}

	for _, strNum := range strings.Split(numbers[1], " ") {
		if strNum == "" {
			continue
		}
		num, err := strconv.Atoi(strNum)
		if err != nil {
			return ScratchCard{}, err
		}
		card.ScratchNumbers = append(card.ScratchNumbers, num)
	}

	return card, nil
}
