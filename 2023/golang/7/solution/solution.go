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
	d := newDeck(
		'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A',
	)
	hands, err := scan(input, d)
	if err != nil {
		return 0, err
	}
	slices.SortFunc(hands, handSortFunc(p1ScoreFunc(d), d.rank))
	sum := 0
	for i, h := range hands {
		sum += (i + 1) * h.bid
	}
	return Answer(sum), nil
}

func SolveP2(input Input) (Answer, error) {
	d := newDeck(
		'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A',
	)
	hands, err := scan(input, d)
	if err != nil {
		return 0, err
	}
	slices.SortFunc(hands, handSortFunc(p2ScoreFunc(d), d.rank))
	sum := 0
	for i, h := range hands {
		sum += (i + 1) * h.bid
	}
	return Answer(sum), nil
}

func p1ScoreFunc(d *deck) func(h hand) (handType int, ranks [5]int) {
	return func(h hand) (handType int, ranks [5]int) {
		counts := make([][2]int, len(d.cards))
		for _, c := range h.cards {
			i := d.index[c]
			counts[i][0] += 1
			counts[i][1] = i + 1
		}
		slices.SortFunc(counts[:], func(i, j [2]int) int {
			if d := j[0] - i[0]; d != 0 {
				return d
			}
			if d := j[1] - i[1]; d != 0 {
				return d
			}
			return 0
		})
		c1, c2 := counts[0][0], counts[1][0]
		handType = handTypeFor(c1, c2)
		for i := range ranks {
			ranks[i] = counts[i][1]
		}
		return handType, ranks
	}
}

func p2ScoreFunc(d *deck) func(h hand) (handType int, ranks [5]int) {
	return func(h hand) (handType int, ranks [5]int) {
		counts := make([][2]int, len(d.cards))
		jokers := 0
		for _, c := range h.cards {
			i := d.index[c]
			if i == 0 {
				jokers += 1
				continue
			}
			counts[i][0] += 1
			counts[i][1] = i + 1
		}
		slices.SortFunc(counts[:], func(i, j [2]int) int {
			if d := j[0] - i[0]; d != 0 {
				return d
			}
			if d := j[1] - i[1]; d != 0 {
				return d
			}
			return 0
		})
		if jokers > 0 && counts[0][1] != 1 {
			counts[0][0] += jokers
		}
		c1, c2 := counts[0][0], counts[1][0]
		handType = handTypeFor(c1, c2)
		for i := range ranks {
			ranks[i] = counts[i][1]
		}
		return handType, ranks
	}
}

func scan(input Input, d *deck) ([]hand, error) {
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
		h, err := d.newHand([]rune(fields[0]), int(bid))
		if err != nil {
			return nil, err
		}
		hands = append(hands, h)
	}
	return hands, nil
}
