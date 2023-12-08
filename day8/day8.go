package day8

import (
	"strings"

	"github.com/jecepeda/advent-of-code-2023/utils"
)

const (
	Left  byte = 'L'
	Right byte = 'R'
)

type Node struct {
	Left  string
	Right string
	Value string
}
type NodeMap map[string]Node

type Problem struct {
	Map          NodeMap
	Instructions string
}

func NewProblem() Problem {
	return Problem{
		Map: make(NodeMap),
	}
}

func greatestCommonDivisor(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func leastCommonMultiple(a, b int, integers ...int) int {
	result := a * b / greatestCommonDivisor(a, b)

	for i := 0; i < len(integers); i++ {
		result = leastCommonMultiple(result, integers[i])
	}

	return result
}

func (p Problem) Solve2(endingLetter, goal string) int {
	startingNodes := map[string]int{}
	for k := range p.Map {
		if strings.HasSuffix(k, endingLetter) {
			startingNodes[k] = 0
		}
	}
	for k := range startingNodes {
		startingNodes[k] = p.Solve(k, goal)
	}
	minimumSteps := []int{}
	for _, v := range startingNodes {
		minimumSteps = append(minimumSteps, v)
	}
	return leastCommonMultiple(minimumSteps[0], minimumSteps[1], minimumSteps[2:]...)
}

func (p Problem) Solve(startingPoint, endingLetter string) int {
	var (
		result       int    = 0
		i            int    = 0
		currentPoint string = startingPoint
		currentNode  Node
	)
	for {
		result++
		whereToGo := p.Instructions[i]
		if whereToGo == Left {
			currentNode = p.Map[currentPoint]
			currentPoint = currentNode.Left
		} else {
			currentNode = p.Map[currentPoint]
			currentPoint = currentNode.Right
		}
		if strings.HasSuffix(currentPoint, endingLetter) {
			return result
		}
		i = (i + 1) % len(p.Instructions)
	}
}

func Part1(filename string) (int, error) {
	problem, err := ReadFile(filename)
	if err != nil {
		return 0, err
	}
	return problem.Solve("AAA", "ZZZ"), nil
}

func Part2(filename string) (int, error) {
	problem, err := ReadFile(filename)
	if err != nil {
		return 0, err
	}
	return problem.Solve2("A", "Z"), nil
}

func ReadFile(filename string) (Problem, error) {
	s, err := utils.ReadFile(filename)
	if err != nil {
		return Problem{}, err
	}
	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	problem := NewProblem()
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if i == 0 {
			problem.Instructions = line
			continue
		} else if i == 1 {
			continue
		}
		node := ParseInstructions(line)
		problem.Map[node.Value] = node
	}
	return problem, nil
}

func ParseInstructions(line string) Node {
	values := strings.Split(line, " = ")

	leftAndRight := strings.Split(values[1], ", ")
	leftAndRight[0] = strings.ReplaceAll(leftAndRight[0], "(", "")
	leftAndRight[1] = strings.ReplaceAll(leftAndRight[1], ")", "")

	return Node{
		Left:  leftAndRight[0],
		Right: leftAndRight[1],
		Value: values[0],
	}
}
