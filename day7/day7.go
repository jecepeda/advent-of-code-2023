package day7

import (
	"sort"
	"strconv"
	"strings"

	"github.com/jecepeda/advent-of-code-2023/utils"
)

var cardRelationShip1 = map[byte]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

var cardRelationShip2 = map[byte]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'J': 1,
}

type HandType int

const (
	HighCard     HandType = 1
	OnePair      HandType = 2
	TwoPairs     HandType = 3
	ThreeOfAKind HandType = 4
	FullHouse    HandType = 5
	FourOfAKind  HandType = 6
	FiveOfAKind  HandType = 7
)

type Hand struct {
	Cards    string
	Type     HandType
	TypeJack HandType
	Bid      int
}

func (h *Hand) InferType() {
	cases := map[rune]int{}
	for _, c := range h.Cards {
		cases[c]++
	}
	switch len(cases) {
	case 5:
		h.Type = HighCard
	case 4:
		h.Type = OnePair
	case 3:
		for _, count := range cases {
			if count == 3 {
				h.Type = ThreeOfAKind
				return
			}
		}
		h.Type = TwoPairs
	case 2:
		for _, v := range cases {
			if v == 4 {
				h.Type = FourOfAKind
				return
			}
			h.Type = FullHouse
		}
	case 1:
		h.Type = FiveOfAKind
	}
}
func (h *Hand) InferTypeJack() {
	cases := map[rune]int{}
	numberOfJacks := 0
	for _, c := range h.Cards {
		if c == 'J' {
			numberOfJacks++
		} else {
			cases[c]++
		}
	}
	highestKey := rune(0)
	highestValue := 0
	for k, v := range cases {
		if v > highestValue {
			highestKey = k
			highestValue = v
		}
	}
	cases[highestKey] = cases[highestKey] + numberOfJacks
	switch len(cases) {
	case 5:
		h.TypeJack = HighCard
	case 4:
		h.TypeJack = OnePair
	case 3:
		for _, count := range cases {
			if count == 3 {
				h.TypeJack = ThreeOfAKind
				return
			}
		}
		h.TypeJack = TwoPairs
	case 2:
		for _, v := range cases {
			if v == 4 {
				h.TypeJack = FourOfAKind
				return
			}
			h.TypeJack = FullHouse
		}
	case 1:
		h.TypeJack = FiveOfAKind
	}
}

func HandOrder1(h1, h2 Hand) bool {
	if h1.Type != h2.Type {
		return h1.Type < h2.Type
	}
	for x := 0; x < len(h1.Cards); x++ {
		if cardRelationShip1[h1.Cards[x]] == cardRelationShip1[h2.Cards[x]] {
			continue
		}
		return cardRelationShip1[h1.Cards[x]] < cardRelationShip1[h2.Cards[x]]
	}
	return false
}

func HandOrder2(h1, h2 Hand) bool {
	if h1.TypeJack != h2.TypeJack {
		return h1.TypeJack < h2.TypeJack
	}
	for x := 0; x < len(h1.Cards); x++ {
		if cardRelationShip2[h1.Cards[x]] == cardRelationShip2[h2.Cards[x]] {
			continue
		}
		return cardRelationShip2[h1.Cards[x]] < cardRelationShip2[h2.Cards[x]]
	}
	return false
}

func Part1(filename string) (int, error) {
	hands, err := ReadFile(filename)
	if err != nil {
		return 0, err
	}

	sort.Slice(hands, func(i, j int) bool {
		return HandOrder1(hands[i], hands[j])
	})

	result := 0
	for rank := 0; rank < len(hands); rank++ {
		result = result + (hands[rank].Bid * (rank + 1))
	}
	return result, nil
}

func Part2(filename string) (int, error) {
	hands, err := ReadFile(filename)
	if err != nil {
		return 0, err
	}

	sort.Slice(hands, func(i, j int) bool {
		return HandOrder2(hands[i], hands[j])
	})

	result := 0
	for rank := 0; rank < len(hands); rank++ {
		result = result + (hands[rank].Bid * (rank + 1))
	}
	return result, nil
}

func ReadFile(filename string) ([]Hand, error) {
	hands := make([]Hand, 0)

	s, err := utils.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	for s.Scan() {
		values := strings.Fields(s.Text())
		bid, err := strconv.Atoi(values[1])
		if err != nil {
			return nil, err
		}
		hands = append(hands, Hand{
			Cards: values[0],
			Type:  HighCard,
			Bid:   bid,
		})
	}
	for i := range hands {
		hands[i].InferType()
		hands[i].InferTypeJack()
	}

	return hands, nil
}
