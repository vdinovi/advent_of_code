package solution

import (
	"bufio"
	"io"
	"strings"
)

const empty = -1

var DigitLiterals = map[string]int{
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

var DigitSpellings = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

type DigitScanner struct {
	scanner  *bufio.Scanner
	prefixes *prefix
	line     []rune
	digits   []int
	err      error
}

func NewDigitScanner(input Input) *DigitScanner {
	s := &DigitScanner{
		scanner:  bufio.NewScanner(input),
		prefixes: newPrefix(),
	}
	s.scanner.Split(bufio.ScanLines)
	return s
}

func (s *DigitScanner) AddWords(words map[string]int) {
	s.prefixes.fill(words)
}

func (s *DigitScanner) Scan() bool {
	s.line = nil
	s.digits = []int{}
	if ok := s.scanner.Scan(); !ok {
		s.err = s.scanner.Err()
		return false
	}
	s.line = []rune(strings.TrimSpace(s.scanner.Text()))
	for i := 0; i < len(s.line); i += 1 {
		if n, ok := match(s.line[i:], s.prefixes); ok {
			s.digits = append(s.digits, n)
		}
	}
	return true
}

func (s *DigitScanner) Digits() []int {
	return s.digits
}

func (s *DigitScanner) Text() string {
	return string(s.line)
}

func (s *DigitScanner) Err() error {
	if s.err == io.EOF {
		return nil
	}
	return s.err
}

type prefix struct {
	val      int
	children map[rune]*prefix
}

func newPrefix() *prefix {
	return &prefix{val: empty, children: make(map[rune]*prefix)}
}

func (p *prefix) fill(words map[string]int) {
	for word, val := range words {
		iter := p
		for _, r := range word {
			child, ok := iter.children[r]
			if !ok {
				iter.children[r] = newPrefix()
				child = iter.children[r]
			}
			iter = child
		}
		iter.val = val
	}
}

func (p *prefix) next(r rune) *prefix {
	if child, ok := p.children[r]; ok {
		return child
	}
	return nil
}

func (p *prefix) isTerminal() bool {
	return len(p.children) == 0
}

func (p *prefix) value() int {
	return p.val
}

func match(line []rune, prefixes *prefix) (int, bool) {
	r := line[0]
	p := prefixes.next(r)
	if p == nil {
		return empty, false
	}
	if p.isTerminal() {
		return p.value(), true
	}
	for j := 1; j < len(line); j += 1 {
		r = line[j]
		p = p.next(r)
		if p == nil {
			return empty, false
		}
		if p.isTerminal() {
			return p.value(), true
		}
	}
	return empty, false
}
