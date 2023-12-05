package solution

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Input io.Reader
type Answer int

func SolveP1(input Input) (Answer, error) {
	pairs, err := parse(input)
	if err != nil {
		return 0, err
	}
	sum := 0
	for _, p := range pairs {
		winners := make(map[int]bool, len(p[1].numbers))
		for _, n := range p[1].numbers {
			winners[n] = true
		}
		matches := 0
		for _, n := range p[0].numbers {
			if winners[n] {
				matches += 1
			}
		}
		sum += int(math.Exp2(float64(matches - 1)))
	}
	return Answer(sum), nil
}

func SolveP2(input Input) (Answer, error) {
	return 0, nil
}

type card struct {
	id      int
	numbers []int
}

var re = regexp.MustCompile(`^Card\s+(\d+): (?:([\s\d]+)) \| (?:([\s\d]+))`)

func parse(input Input) (pairs [][2]card, err error) {
	pairs = [][2]card{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		match := re.FindAllStringSubmatch(line, -1)
		if match == nil || len(match) < 1 || len(match[0]) < 3 {
			return nil, fmt.Errorf("invalid line %q", line)
		}
		id, err := strconv.ParseInt(match[0][1], 10, 64)
		if err != nil {
			return nil, err
		}
		numbers := [][]string{strings.Fields(match[0][2]), strings.Fields(match[0][3])}
		pair := [2]card{{id: int(id)}, {id: int(id)}}
		for c := 0; c < 2; c += 1 {
			pair[c].numbers = make([]int, len(numbers[c]))
			for i, s := range numbers[c] {
				n, err := strconv.ParseInt(s, 10, 64)
				if err != nil {
					return nil, err
				}
				pair[c].numbers[i] = int(n)
			}
		}
		pairs = append(pairs, pair)
	}
	err = scanner.Err()
	return pairs, err
}
