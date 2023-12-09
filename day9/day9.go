package day9

import (
	"strconv"
	"strings"

	"github.com/jecepeda/advent-of-code-2023/utils"
)

func difference(numbers []int) []int {
	var result = make([]int, len(numbers)-1)
	for i := 1; i < len(numbers); i++ {
		result[i-1] = numbers[i] - numbers[i-1]
	}
	return result
}

func allZeroes(numbers []int) bool {
	for _, n := range numbers {
		if n != 0 {
			return false
		}
	}
	return true
}

func nextSequence(numbers []int) int {
	if allZeroes(numbers) {
		return 0
	}
	next := difference(numbers)
	value := nextSequence(next)
	return value + numbers[len(numbers)-1]
}

func prevSequence(numbers []int) int {
	if allZeroes(numbers) {
		return 0
	}
	next := difference(numbers)
	value := prevSequence(next)
	return numbers[0] - value
}

func Part1(filename string) (int, error) {
	sequences, err := ReadFile(filename)
	if err != nil {
		return 0, err
	}
	var result int
	for _, seq := range sequences {
		result += nextSequence(seq)
	}
	return result, nil
}

func Part2(filename string) (int, error) {
	sequences, err := ReadFile(filename)
	if err != nil {
		return 0, err
	}
	var result int
	for _, seq := range sequences {
		result += prevSequence(seq)
	}
	return result, nil
}

func ReadFile(filename string) ([][]int, error) {
	s, err := utils.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var result [][]int
	for s.Scan() {
		fields := strings.Fields(s.Text())
		var row []int
		for _, field := range fields {
			value, err := strconv.Atoi(field)
			if err != nil {
				return nil, err
			}
			row = append(row, value)
		}
		result = append(result, row)
	}
	return result, nil
}
