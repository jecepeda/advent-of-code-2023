package day3

import (
	"github.com/jecepeda/advent-of-code-2023/utils"
)

type CodeMap struct {
	codeMap   [][]rune
	NumberMap map[[2]int]int
}

func NewCodeMap(codeMap [][]rune) *CodeMap {
	return &CodeMap{
		codeMap:   codeMap,
		NumberMap: make(map[[2]int]int),
	}
}

func (c *CodeMap) Part1() (int, error) {
	c.processNumbers()
	result := 0
	for i := range c.codeMap {
		for j := range c.codeMap[i] {
			if !c.isNumber(i, j) && !c.isDot(i, j) {
				adjacents := c.getAdjacents(i, j)
				for _, adjacent := range adjacents {
					result += adjacent
				}
			}
		}
	}
	return result, nil
}

func (c *CodeMap) Part2() (int, error) {
	c.processNumbers()
	result := 0
	for i := range c.codeMap {
		for j := range c.codeMap[i] {
			if c.isGear(i, j) {
				adjacents := c.getAdjacents(i, j)
				if len(adjacents) != 2 {
					continue
				}
				result += adjacents[0] * adjacents[1]
			}
		}
	}
	return result, nil
}

func (c *CodeMap) getAdjacents(x, y int) []int {
	adjacents := []int{}
	for i := x - 1; i <= x+1; i++ {
		prevAdj := 0
		for j := y - 1; j <= y+1; j++ {
			if c.isNumber(i, j) {
				num := c.NumberMap[[2]int{i, j}]
				// avoid getting the same number on the same line
				// 345
				// .*.
				// would only take 345 once
				if num != prevAdj {
					adjacents = append(adjacents, num)
					prevAdj = num
				}
			}
		}
	}
	return adjacents
}

func (c *CodeMap) processNumbers() {
	for i := 0; i < len(c.codeMap); i++ {
		var number int = 0
		var startPos int = -1
		for j := 0; j < len(c.codeMap[i]); j++ {
			if c.isNumber(i, j) {
				if startPos == -1 {
					startPos = j
					number = c.parseNumber(i, j)
				} else {
					number = number*10 + c.parseNumber(i, j)
				}
			} else if startPos != -1 {
				// save the number and reset
				for k := startPos; k < j; k++ {
					c.NumberMap[[2]int{i, k}] = number
				}
				startPos = -1
				number = 0
			}
		}
		if startPos != -1 {
			for k := startPos; k < len(c.codeMap[i]); k++ {
				c.NumberMap[[2]int{i, k}] = number
			}
		}
	}
}

func (c *CodeMap) isGear(x, y int) bool {
	return c.codeMap[x][y] == '*'
}

func (c *CodeMap) isDot(x, y int) bool {
	return c.codeMap[x][y] == '.'
}

func (c *CodeMap) parseNumber(x, y int) int {
	return int(c.codeMap[x][y] - '0')
}

func (c *CodeMap) isNumber(x, y int) bool {
	return c.codeMap[x][y] >= '0' && c.codeMap[x][y] <= '9'
}

func Part1(filename string) (int, error) {
	s, err := utils.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	codeMap := make([][]rune, 0)
	for s.Scan() {
		codeMap = append(codeMap, []rune(s.Text()))
	}

	c := NewCodeMap(codeMap)
	result, err := c.Part1()
	if err != nil {
		return 0, err
	}
	return result, nil
}

func Part2(filename string) (int, error) {
	s, err := utils.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	codeMap := make([][]rune, 0)
	for s.Scan() {
		codeMap = append(codeMap, []rune(s.Text()))
	}

	c := NewCodeMap(codeMap)
	result, err := c.Part2()
	if err != nil {
		return 0, err
	}
	return result, nil
}
