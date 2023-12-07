package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadFile(filename string) (*bufio.Scanner, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	r := bufio.NewScanner(f)
	r.Split(bufio.ScanLines)
	return r, nil
}

func ReadNumbers(line string) ([]int, error) {
	var result []int
	for _, num := range strings.Split(line, " ") {
		if num == "" {
			continue
		}
		i, err := strconv.Atoi(num)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}
	return result, nil
}
