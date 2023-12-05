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
		n := matches(p[0], p[1])
		sum += score(n)
	}
	return Answer(sum), nil
}

func SolveP2(input Input) (Answer, error) {
	pairs, err := parse(input)
	if err != nil {
		return 0, err
	}
	sum := 0
	copies := make([]int, len(pairs))
	for i, p := range pairs {
		n := matches(p[0], p[1])
		for j := i + 1; j < i+1+n; j += 1 {
			copies[j] += 1 + copies[i]
		}
	}
	for i := range pairs {
		sum += 1 + copies[i]
	}
	return Answer(sum), nil
}

func matches(a, b card) int {
	winners := make(map[int]bool, len(b.numbers))
	for _, n := range b.numbers {
		winners[n] = true
	}
	matches := 0
	for _, n := range a.numbers {
		if winners[n] {
			matches += 1
		}
	}
	return matches
}

func score(n int) int {
	return int(math.Exp2(float64(n - 1)))
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
