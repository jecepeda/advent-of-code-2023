package day1

import (
	"strconv"
	"strings"

	"github.com/jecepeda/advent-of-code-2023/utils"
)

var numbers = [][2]string{
	{"one", "o1e"},
	{"two", "t2o"},
	{"three", "t3e"},
	{"four", "f4r"},
	{"five", "f5e"},
	{"six", "s6x"},
	{"seven", "s7n"},
	{"eight", "e8t"},
	{"nine", "n9e"},
}

func getNumber(s string) (int, error) {
	var runes [2]rune = [2]rune{-1, -1}
	for _, c := range s {
		if c >= '0' && c <= '9' {
			if runes[0] == -1 {
				runes[0] = c
			}
			runes[1] = c
		}
	}
	result, err := strconv.Atoi(string(runes[:]))
	if err != nil {
		return 0, err
	}
	return result, nil
}

func cleanLine(s string) string {
	i := 0
	for i < len(s) {
		for _, item := range numbers {
			strNum := item[0]
			num := item[1]
			if i+len(strNum) > len(s) {
				continue
			}
			if s[i:i+len(strNum)] == strNum {
				s = strings.Replace(s, strNum, num, 1)
				break
			}
		}
		i++
	}
	return s

}

func Part1(filename string) (int, error) {
	r, err := utils.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	result := 0
	for r.Scan() {
		number, err := getNumber(r.Text())
		if err != nil {
			return 0, err
		}
		result += number
	}
	if r.Err() != nil {
		return 0, r.Err()
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
		number, err := getNumber(cleanLine(r.Text()))
		if err != nil {
			return 0, err
		}
		result += number
	}
	if r.Err() != nil {
		return 0, r.Err()
	}

	return result, nil
}
