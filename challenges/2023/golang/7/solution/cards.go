package solution

import (
	"fmt"
	"unicode"
)

type deck struct {
	cards []card
	index map[card]int
	runes map[rune]card
}

func newDeck(cards ...card) *deck {
	deck := &deck{
		cards: make([]card, len(cards)),
		index: make(map[card]int, len(cards)),
		runes: make(map[rune]card),
	}
	for i, c := range cards {
		deck.cards[i] = c
		deck.runes[rune(c)] = c
		deck.index[c] = i
	}
	return deck
}

type card rune

func (d *deck) newCard(r rune) (card, error) {
	r = unicode.ToUpper(r)
	if c, ok := d.runes[r]; ok {
		return c, nil
	}
	return 0, &InvalidCardError{Card: r}
}

func (d *deck) rank(c card) (rank int) {
	return d.index[c]
}

func (c card) String() string {
	return string(c)
}

type hand struct {
	cards [5]card
	bid   int
}

func (d *deck) newHand(cards []rune, bid int) (h hand, err error) {
	if len(cards) != len(h.cards) {
		return h, &InvalidHandError{Cards: cards}
	}
	h.bid = bid
	for i, r := range cards[:len(h.cards)] {
		h.cards[i], err = d.newCard(r)
		if err != nil {
			return h, err

		}
	}
	return h, nil
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

type scoreFunc func(h hand) (handType int, ranks [5]int)
type rankFunc func(c card) (rank int)

func handSortFunc(score scoreFunc, rank rankFunc) func(i, j hand) int {
	return func(i, j hand) int {
		iht, _ := score(i)
		jht, _ := score(j)
		if d := iht - jht; d != 0 {
			return d
		}
		for k := range i.cards {
			if d := rank(i.cards[k]) - rank(j.cards[k]); d != 0 {
				return d
			}
		}
		return 0
	}
}

func (h hand) String(score scoreFunc) string {
	ht, ranks := score(h)
	return fmt.Sprintf("%s %d (%s %v)", string(h.cards[:]), h.bid, handTypeString(ht), ranks)
}

const (
	highCard int = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

func handTypeString(ht int) string {
	switch ht {
	case highCard:
		return "HighCard"
	case onePair:
		return "OnePair"
	case twoPair:
		return "TwoPair"
	case threeOfAKind:
		return "ThreeOfAKind"
	case fullHouse:
		return "FullHouse"
	case fourOfAKind:
		return "FourOfAKind"
	case fiveOfAKind:
		return "FiveOfAKind"
	default:
		return "Unknown"
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
