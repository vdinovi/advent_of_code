package lib

import (
	"fmt"
	"strings"
)

type Position struct {
	Row int
	Col int
}

type GridEntry[T fmt.Stringer] struct {
	Present bool
	Position
	Item T
}

func (e *GridEntry[T]) Distance(other *GridEntry[T]) (height int, width int) {
	if e.Position.Row > other.Position.Row {
		height = e.Position.Row - other.Position.Row
	} else {
		height = other.Position.Row - e.Position.Row
	}
	if e.Position.Col > other.Position.Col {
		width = e.Position.Col - other.Position.Col
	} else {
		width = other.Position.Col - e.Position.Col
	}
	return height, width
}

type Grid[T fmt.Stringer] struct {
	rows [][]GridEntry[T]
	pos  Position
}

func NewGrid[T fmt.Stringer]() *Grid[T] {
	return &Grid[T]{rows: make([][]GridEntry[T], 0)}
}

func (g *Grid[T]) Height() int {
	return len(g.rows)
}

func (g *Grid[T]) Width() (width int) {
	for _, row := range g.rows {
		width = max(width, len(row))
	}
	return width
}

func (g *Grid[T]) Add(item T, present bool) {
	if len(g.rows) <= g.pos.Row {
		g.rows = append(g.rows, make([]GridEntry[T], 0))
	}
	g.rows[g.pos.Row] = append(g.rows[g.pos.Row], GridEntry[T]{
		Item:     item,
		Present:  present,
		Position: g.pos,
	})
	g.pos.Col += 1
}

func (g *Grid[T]) NextRow() {
	g.pos.Row += 1
	g.pos.Col = 0
}

func (g *Grid[T]) At(pos Position) (*GridEntry[T], error) {
	if pos.Row < 0 || pos.Row >= len(g.rows) {
		return nil, &InvalidPositionError{pos: pos}
	}
	if pos.Col < 0 || pos.Col >= len(g.rows[pos.Row]) {
		return nil, &InvalidPositionError{pos: pos}
	}
	return &g.rows[pos.Row][pos.Col], nil
}

// Like `At“ but panics on error -- used for testing purposes
func (g *Grid[T]) Get(pos Position) *GridEntry[T] {
	e, err := g.At(pos)
	if err != nil {
		panic(err)
	}
	return e
}

type InvalidPositionError struct {
	pos Position
}

func (e *InvalidPositionError) Error() string {
	return fmt.Sprintf("invalid position (%d, %d)", e.pos.Row, e.pos.Col)
}

func (g *Grid[T]) String() string {
	var sb strings.Builder
	for _, row := range g.rows {
		for _, entry := range row {
			if _, err := sb.WriteString(entry.Item.String()); err != nil {
				panic(err)
			}
		}
		if _, err := sb.WriteRune('\n'); err != nil {
			panic(err)
		}
	}
	return sb.String()
}

func (g *Grid[T]) InsertRow(before int, item T, present bool) {
	g.rows = append(g.rows, make([]GridEntry[T], len(g.rows[before])))
	for r := len(g.rows) - 1; r >= before; r -= 1 {
		for c := 0; c < len(g.rows[r]); c += 1 {
			g.rows[r][c] = g.rows[r-1][c]
			g.rows[r][c].Position.Row += 1
		}
	}
	for c := 0; c < len(g.rows[before]); c += 1 {
		g.rows[before][c] = GridEntry[T]{
			Item:     item,
			Present:  present,
			Position: Position{before, c},
		}
	}
}

func (g *Grid[T]) InsertColumn(before int, item T, present bool) {
	for r := 0; r < len(g.rows); r += 1 {
		g.rows[r] = append(g.rows[r], GridEntry[T]{
			Item:     item,
			Present:  present,
			Position: Position{r, before},
		})
		for c := len(g.rows[r]) - 1; c > before; c -= 1 {
			g.rows[r][c] = g.rows[r][c-1]
			g.rows[r][c].Position.Col += 1
		}
		g.rows[r][before] = GridEntry[T]{
			Item:     item,
			Present:  present,
			Position: Position{r, before},
		}
	}
}

type GridIterator[T fmt.Stringer] struct {
	g   *Grid[T]
	cur Position
	err error
}

func (g *Grid[T]) Iterator() *GridIterator[T] {
	return &GridIterator[T]{
		g:   g,
		cur: Position{Row: 0, Col: -1},
	}
}

func (it *GridIterator[T]) Next() bool {
	if it.done() {
		return false
	}
	it.cur.Col += 1
	if it.cur.Col >= len(it.g.rows[it.cur.Row]) {
		it.cur.Row += 1
		it.cur.Col = 0
	}
	return !it.done()
}

func (it *GridIterator[T]) Entry() *GridEntry[T] {
	if it.done() {
		return nil
	}
	return &it.g.rows[it.cur.Row][it.cur.Col]
}

func (it *GridIterator[T]) Err() error {
	return it.err
}

func (it *GridIterator[T]) done() bool {
	return it.err != nil ||
		it.cur.Row >= len(it.g.rows)
}

type Edge[T fmt.Stringer, W any] struct {
	Weight int
	Entry  *GridEntry[T]
}

func (g *Grid[T]) Distances(include func(*GridEntry[T]) bool) (map[*GridEntry[T]][]Edge[T, int], error) {
	dists := make(map[*GridEntry[T]][]Edge[T, int])
	var (
		iter1 *GridIterator[T]
		iter2 *GridIterator[T]
	)
	for iter1 = g.Iterator(); iter1.Next(); {
		entry := iter1.Entry()
		if include(entry) {
			dists[entry] = make([]Edge[T, int], 0)
			for iter2 = g.Iterator(); iter2.Next(); {
				other := iter2.Entry()
				if include(other) && entry != other {
					width, height := entry.Distance(other)
					dists[entry] = append(dists[entry], Edge[T, int]{
						Weight: width + height,
						Entry:  other,
					})
				}
			}
			if iter2.Err() != nil {
				return nil, iter2.Err()
			}
		}
	}
	return dists, iter1.Err()
}

// func (g *Grid[T]) DistanceBetween(from, to *GridEntry[T]) int {
// 	return 0
// }