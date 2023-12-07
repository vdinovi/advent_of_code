package solution

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
)

type Input io.Reader
type Answer int

func SolveP1(input Input) (Answer, error) {
	hands, err := scan(input)
	if err != nil {
		return 0, err
	}
	slices.SortFunc(hands, handSortFunc)
	sum := 0
	for i, h := range hands {
		sum += (i + 1) * h.bid
	}
	return Answer(sum), nil
}

func SolveP2(input Input) (Answer, error) {
	return 0, nil
}

func scan(input Input) ([]hand, error) {
	scanner := bufio.NewScanner(input)
	hands := make([]hand, 0)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		line = strings.ToUpper(line)
		fields := strings.Fields(line)
		if len(fields) != 2 {
			return nil, fmt.Errorf("invalid line %s", line)
		}
		bid, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, err
		}
		h, err := newHand([]rune(fields[0]), int(bid))
		if err != nil {
			return nil, err
		}
		hands = append(hands, h)
	}
	return hands, nil
}
