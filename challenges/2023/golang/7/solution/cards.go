package solution

import (
	"fmt"
	"slices"
	"unicode"
)

type card rune

func newCard(r rune) (card, error) {
	r = unicode.ToUpper(r)
	if c, ok := runes[r]; ok {
		return c, nil
	}
	return 0, &InvalidCardError{Card: r}
}

func (c card) String() string {
	return string(c)
}

type hand struct {
	cards [5]card
	bid   int
}

func newHand(cards []rune, bid int) (h hand, err error) {
	if len(cards) != len(h.cards) {
		return h, &InvalidHandError{Cards: cards}
	}
	h.bid = bid
	for i, r := range cards[:len(h.cards)] {
		h.cards[i], err = newCard(r)
		if err != nil {
			return h, err

		}
	}
	return h, nil
}

func (h hand) score() (handType int, ranks [5]int) {
	counts := [len(cards)][2]int{}
	for _, c := range h.cards {
		i := index[c]
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

func handTypeFor(c1, c2 int) int {
	switch c1 {
	case 5:
		return fiveOfAKind
	case 4:
		return fourOfAKind
	case 3:
		switch c2 {
		case 2:
			return fullHouse
		default:
			return threeOfAKind
		}
	case 2:
		switch c2 {
		case 2:
			return twoPair
		default:
			return onePair
		}
	default:
		return highCard
	}
}

func handSortFunc(i, j hand) int {
	iht, _ := i.score()
	jht, _ := j.score()
	if d := iht - jht; d != 0 {
		return d
	}
	for k := range i.cards {
		if d := index[i.cards[k]] - index[j.cards[k]]; d != 0 {
			return d
		}
	}
	return 0
}

func (h hand) String() string {
	ht, ranks := h.score()
	return fmt.Sprintf("%s %d (%d %v)", string(h.cards[:]), h.bid, ht, ranks)
}

var cards = [...]card{
	'1', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A',
}
var index map[card]int
var runes map[rune]card

const (
	highCard int = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

func init() {
	index = make(map[card]int, len(cards))
	runes = make(map[rune]card)
	for i, c := range cards {
		cards[i] = c
		runes[rune(c)] = c
		index[c] = i
	}
}

type InvalidCardError struct {
	Card rune
}

func (e *InvalidCardError) Error() string {
	return fmt.Sprintf("invalid card %c", rune(e.Card))
}

type InvalidHandError struct {
	Cards []rune
}

func (e *InvalidHandError) Error() string {
	return fmt.Sprintf("invalid hand %s", string(e.Cards))
}
