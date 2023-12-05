package solution

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"unicode"
	"unicode/utf8"
)

type Input io.Reader
type Answer int

func SolveP1(input Input) (Answer, error) {
	parts, symbols, width, height, err := parse(input)
	if err != nil {
		return 0, err
	}
	sm := symbolMap(symbols, width, height)
	sum := 0
	for _, p := range parts {
		if _, ok := isAdjacent(p, sm, width, height); ok {
			sum += p.number
		}
	}
	return Answer(sum), nil
}

func SolveP2(input Input) (Answer, error) {
	parts, symbols, width, height, err := parse(input)
	if err != nil {
		return 0, err
	}
	sm := symbolMap(symbols, width, height)
	for _, p := range parts {
		if s, ok := isAdjacent(p, sm, width, height); ok {
			if s.adjacencies == nil {
				s.adjacencies = []int{p.number}
			} else {
				s.adjacencies = append(s.adjacencies, p.number)
			}
		}
	}
	sum := 0
	for _, row := range sm {
		for _, s := range row {
			if s != nil && len(s.adjacencies) == 2 {
				sum += s.adjacencies[0] * s.adjacencies[1]
			}
		}
	}
	return Answer(sum), nil
}

type position struct {
	row   int
	start int
	end   int
}

type part struct {
	number int
	pos    position
}

type symbol struct {
	rune
	pos         position
	adjacencies []int
}

func isAdjacent(part part, sm [][]*symbol, width, height int) (*symbol, bool) {
	for i := max(part.pos.row-1, 0); i <= min(part.pos.row+1, height-1); i += 1 {
		for j := max(part.pos.start-1, 0); j <= min(part.pos.end, width-1); j += 1 {
			if sm[i][j] != nil {
				return sm[i][j], true
			}
		}
	}
	return nil, false
}

func parse(input Input) ([]part, []symbol, int, int, error) {
	var (
		width  int
		height int
	)
	parts := []part{}
	symbols := []symbol{}
	scanner := bufio.NewScanner(input)
	pos := &position{}
	scanner.Split(pos.split)
	for scanner.Scan() {
		token := scanner.Text()
		height = max(height, pos.row+1)
		width = max(width, pos.end)
		r := []rune(token)[0]
		if isSymbol(r) {
			symbols = append(symbols, symbol{rune: r, pos: *pos})
		} else {
			num, err := strconv.ParseInt(token, 10, 64)
			if err != nil {
				return nil, nil, 0, 0, err
			}
			parts = append(parts, part{number: int(num), pos: *pos})
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, 0, 0, err
	}
	return parts, symbols, width, height, nil
}

func symbolMap(syms []symbol, width, height int) [][]*symbol {
	m := make([][]*symbol, height)
	for i := 0; i < height; i += 1 {
		m[i] = make([]*symbol, width)
	}
	for i, s := range syms {
		m[s.pos.row][s.pos.start] = &syms[i]
	}
	return m
}

func (p *position) split(data []byte, atEOF bool) (advance int, token []byte, err error) {
	p.start = p.end
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if r == '\n' {
			p.start = 0
			p.end = 0
			p.row += 1
		} else if r == '.' {
			p.start += 1
			p.end += 1
		} else if unicode.IsSpace(r) {
			continue
		} else if isSymbol(r) {
			p.end += 1
			return start + width, data[start : start+width], nil
		} else if unicode.IsDigit(r) {
			break
		} else {
			err = fmt.Errorf("unknown character %c", r)
			return start + width, data[start : start+width], err
		}
	}
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if !unicode.IsDigit(r) {
			return i, data[start:i], nil
		}
		p.end += 1
	}
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	return start, nil, nil
}

var symbols = map[rune]bool{
	'+': true,
	'-': true,
	'*': true,
	'/': true,
	'$': true,
	'#': true,
	'&': true,
	'@': true,
	'%': true,
	'=': true,
}

func isSymbol(r rune) bool {
	return symbols[r]
}
