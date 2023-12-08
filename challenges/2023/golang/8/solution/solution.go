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
	stop := func(n node) bool {
		for _, r := range n {
			if r != 'Z' {
				return false
			}
		}
		return true
	}
	steps := traverse(node("AAA"), stop, dirs, chart)
	return Answer(steps), nil
}

func SolveP2(input Input) (Answer, error) {
	chart, dirs, err := scan(input)
	if err != nil {
		return 0, err
	}
	starts := []node{}
	for n := range chart {
		if n[len(n)-1] == 'A' {
			starts = append(starts, n)
		}
	}
	stop := func(n node) bool {
		return n[len(n)-1] == 'Z'
	}
	periods := make([]int, len(starts))
	for i, cur := range starts {
		periods[i] = traverse(cur, stop, dirs, chart)
	}
	return Answer(lcm(periods...)), nil

}

type stopFunc = func(node) bool

func traverse(start node, stop stopFunc, dirs []direction, chart map[node][2]node) (steps int) {
	current := start
	for i := 0; !stop(current); i += 1 {
		if i >= len(dirs) {
			i = 0
		}
		current = chart[current][dirs[i]]
		steps += 1
	}
	return steps
}

func lcm(n ...int) int {
	result := n[0]
	for _, n := range n[1:] {
		result = n * result / gcd(n, result)
	}
	return result
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
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

var re = regexp.MustCompile(`^([A-Z\d]{3})\s*=\s*\(([A-Z\d]{3}),\s*([A-Z\d]{3})\)$`)

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
