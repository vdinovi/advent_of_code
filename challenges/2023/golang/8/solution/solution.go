package solution

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

type Input io.Reader
type Answer int

func SolveP1(input Input) (Answer, error) {
	chart, dirs, err := scan(input)
	if err != nil {
		return 0, err
	}
	current := node("AAA")
	target := node("ZZZ")
	steps := 0
	for i := 0; current != target; i += 1 {
		if i >= len(dirs) {
			i = 0
		}
		current = chart[current][dirs[i]]
		steps += 1
	}
	return Answer(steps), nil
}

func SolveP2(input Input) (Answer, error) {
	return 0, nil
}

type node string

type direction int

const (
	left direction = iota
	right
)

func (d direction) String() string {
	switch d {
	case left:
		return "L"
	case right:
		return "R"
	default:
		panic("unknown direction")
	}
}

var re = regexp.MustCompile(`^([A-Z]{3})\s*=\s*\(([A-Z]{3}),\s*([A-Z]{3})\)$`)

func scan(input Input) (chart map[node][2]node, dirs []direction, err error) {
	chart = make(map[node][2]node)
	dirs = []direction{}
	scanner := bufio.NewScanner(input)
	if !scanner.Scan() {
		return nil, nil, scanner.Err()
	}
	line := strings.TrimSpace(scanner.Text())
	line = strings.ToUpper(line)
	for _, r := range line {
		switch r {
		case 'L':
			dirs = append(dirs, left)
		case 'R':
			dirs = append(dirs, right)
		default:
			return nil, nil, fmt.Errorf("unknown direction '%c'", r)
		}
	}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		line = strings.ToUpper(line)
		match := re.FindAllStringSubmatch(line, -1)
		if match == nil || len(match) < 1 || len(match[0]) < 4 {
			return nil, nil, fmt.Errorf("invalid line %s", line)
		}
		chart[node(match[0][1])] = [2]node{node(match[0][2]), node(match[0][3])}
	}
	return chart, dirs, nil
}
