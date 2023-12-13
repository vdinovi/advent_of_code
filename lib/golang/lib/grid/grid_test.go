package grid_test

import (
	_ "embed"
	"strings"
	"testing"
	"unicode"

	"github.com/vdinovi/advent_of_code/lib/golang/lib/grid"
)

func TestGrid(t *testing.T) {
	g := ticTacToeGrid(t)
	pos := grid.Position{Row: 0, Col: 0}
	for _, r := range ticTacToeString {
		if r == '\n' {
			pos.Row += 1
			pos.Col = 0
		} else if !unicode.IsSpace(r) {
			if entry, err := g.At(pos); err != nil {
				t.Errorf("unexpected error %s", err)
			} else if entry == nil {
				t.Errorf("expected entry at %v to be present but got nil", pos)
			} else if entry.Item != square(r) {
				t.Errorf("expected entry at %v to be %c but got %c", pos, r, entry.Item)
			}
			pos.Col += 1
		}
	}
	// fuzz around the edges
	for r := -1; r <= 4; r += 1 {
		for c := -1; c <= 4; c += 1 {
			pos := grid.Position{Row: r, Col: c}
			_, err := g.At(pos)
			if r < 0 || r > 2 || c < 0 || c > 2 {
				if err == nil {
					t.Errorf("expected error at %v but got nil", pos)
				} else if _, ok := err.(*grid.InvalidPositionError); !ok {
					t.Errorf("expected error at %v to be InvalidPositionError but got %s", pos, err)
				}
			} else if err != nil {
				t.Errorf("unexpected error %s", err)
			}
		}
	}
	if str := g.String(); str != ticTacToeString {
		t.Errorf("expected grid to have string %q but got %q", ticTacToeString, str)
	}
}

func TestGridIterator(t *testing.T) {
	g := ticTacToeGrid(t)
	var sb strings.Builder
	row := 0
	for it := g.Iterator(); it.Next(); {
		entry := it.Entry()
		if entry.Position.Row != row {
			row = entry.Position.Row
			sb.WriteRune('\n')
		}
		sb.WriteRune(rune(entry.Item))
	}
	if str := g.String(); str != ticTacToeString {
		t.Errorf("expected iterator to have produced %q but got %q", ticTacToeString, str)
	}
}

func TestGridInsert(t *testing.T) {
	g := grid.NewGrid[square]()
	for i := 0; i < 3; i += 1 {
		for j := 0; j < 3; j += 1 {
			g.Add('.', false)
		}
		g.NextRow()
	}
	if h := g.Height(); h != 3 {
		t.Fatalf("expected grid to have height=%d but got %d", 3, h)
	}
	if w := g.Width(); w != 3 {
		t.Fatalf("expected grid to have width=%d but got %d", 3, w)
	}

	g.InsertRow(1, square('R'), false)
	if h := g.Height(); h != 4 {
		t.Fatalf("expected grid to have height=%d but got %d", 3, h)
	}
	if w := g.Width(); w != 3 {
		t.Fatalf("expected grid to have width=%d but got %d", 3, w)
	}

	g.InsertColumn(2, square('C'), false)
	if h := g.Height(); h != 4 {
		t.Fatalf("expected grid to have height=%d but got %d", 3, h)
	}
	if w := g.Width(); w != 4 {
		t.Fatalf("expected grid to have width=%d but got %d", 3, w)
	}

	for i := 0; i < 4; i += 1 {
		for j := 0; j < 4; j += 1 {
			entry, err := g.At(grid.Position{i, j})
			if err != nil {
				t.Fatalf("unexpected error %s", err)
			}
			want := square('.')
			if j == 2 {
				want = square('C')
			} else if i == 1 {
				want = square('R')
			}
			if entry.Item != want {
				t.Log(g.String())
				t.Fatalf("expected entry at %s to be %c but got %c", entry.Position, want, entry.Item)
			}
		}
	}
}

func TestGridNeighbors(t *testing.T) {
	g := ticTacToeGrid(t)
	expected := []struct {
		from      *grid.GridEntry[square]
		neighbors [4]*grid.GridEntry[square]
	}{
		{
			from: get(t, g, 0, 0),
			neighbors: [4]*grid.GridEntry[square]{
				nil,
				get(t, g, 0, 1),
				get(t, g, 1, 0),
				nil,
			},
		},
		{
			from: get(t, g, 0, 1),
			neighbors: [4]*grid.GridEntry[square]{
				nil,
				get(t, g, 0, 2),
				get(t, g, 1, 1),
				get(t, g, 0, 0),
			},
		},
		{
			from: get(t, g, 0, 2),
			neighbors: [4]*grid.GridEntry[square]{
				nil,
				nil,
				get(t, g, 1, 2),
				get(t, g, 0, 1),
			},
		},
		{
			from: get(t, g, 1, 0),
			neighbors: [4]*grid.GridEntry[square]{
				get(t, g, 0, 0),
				get(t, g, 1, 1),
				get(t, g, 2, 0),
				nil,
			},
		},
		{
			from: get(t, g, 1, 1),
			neighbors: [4]*grid.GridEntry[square]{
				get(t, g, 0, 1),
				get(t, g, 1, 2),
				get(t, g, 2, 1),
				get(t, g, 1, 0),
			},
		},
		{
			from: get(t, g, 1, 2),
			neighbors: [4]*grid.GridEntry[square]{
				get(t, g, 0, 2),
				nil,
				get(t, g, 2, 2),
				get(t, g, 1, 1),
			},
		},
		{
			from: get(t, g, 2, 0),
			neighbors: [4]*grid.GridEntry[square]{
				get(t, g, 1, 0),
				get(t, g, 2, 1),
				nil,
				nil,
			},
		},
		{
			from: get(t, g, 2, 1),
			neighbors: [4]*grid.GridEntry[square]{
				get(t, g, 1, 1),
				get(t, g, 2, 2),
				nil,
				get(t, g, 2, 0),
			},
		},
		{
			from: get(t, g, 2, 2),
			neighbors: [4]*grid.GridEntry[square]{
				get(t, g, 1, 2),
				nil,
				nil,
				get(t, g, 2, 1),
			},
		},
	}
	for _, exp := range expected {
		e := exp.from
		for i, n := range g.Neighbors(e) {
			if n != exp.neighbors[i] {
				if m := exp.neighbors[i]; m == nil {
					t.Errorf("expected neighbor[%d] of %s to be nil but got %s", i, e.Position, n.Position)
				} else if n == nil {
					t.Errorf("expected neighbor[%d] of %s to be %s but got nil", i, e.Position, m.Position)
				} else {
					t.Errorf("expected neighbor[%d] of %s to be %s but got %s",
						i, e.Position, m.Position, n.Position)
				}
			}
		}
	}
}

func TestGridDiagonals(t *testing.T) {
	g := ticTacToeGrid(t)
	expected := []struct {
		from      *grid.GridEntry[square]
		diagonals [4]*grid.GridEntry[square]
	}{
		{
			from: get(t, g, 0, 0),
			diagonals: [4]*grid.GridEntry[square]{
				nil,
				nil,
				get(t, g, 1, 1),
				nil,
			},
		},
		{
			from: get(t, g, 0, 1),
			diagonals: [4]*grid.GridEntry[square]{
				nil,
				nil,
				get(t, g, 1, 2),
				get(t, g, 1, 0),
			},
		},
		{
			from: get(t, g, 0, 2),
			diagonals: [4]*grid.GridEntry[square]{
				nil,
				nil,
				nil,
				get(t, g, 1, 1),
			},
		},
		{
			from: get(t, g, 1, 0),
			diagonals: [4]*grid.GridEntry[square]{
				nil,
				get(t, g, 0, 1),
				get(t, g, 2, 1),
				nil,
			},
		},
		{
			from: get(t, g, 1, 1),
			diagonals: [4]*grid.GridEntry[square]{
				get(t, g, 0, 0),
				get(t, g, 0, 2),
				get(t, g, 2, 2),
				get(t, g, 2, 0),
			},
		},
		{
			from: get(t, g, 1, 2),
			diagonals: [4]*grid.GridEntry[square]{
				get(t, g, 0, 1),
				nil,
				nil,
				get(t, g, 2, 1),
			},
		},
		{
			from: get(t, g, 2, 0),
			diagonals: [4]*grid.GridEntry[square]{
				nil,
				get(t, g, 1, 1),
				nil,
				nil,
			},
		},
		{
			from: get(t, g, 2, 1),
			diagonals: [4]*grid.GridEntry[square]{
				get(t, g, 1, 0),
				get(t, g, 1, 2),
				nil,
				nil,
			},
		},
		{
			from: get(t, g, 2, 2),
			diagonals: [4]*grid.GridEntry[square]{
				get(t, g, 1, 1),
				nil,
				nil,
				nil,
			},
		},
	}
	for _, exp := range expected {
		e := exp.from
		for i, n := range g.Diagonals(e) {
			if n != exp.diagonals[i] {
				if m := exp.diagonals[i]; m == nil {
					t.Errorf("expected diagonal[%d] of %s to be nil but got %s", i, e.Position, n.Position)
				} else if n == nil {
					t.Errorf("expected diagonal[%d] of %s to be %s but got nil", i, e.Position, m.Position)
				} else {
					t.Errorf("expected diagonal[%d] of %s to be %s but got %s",
						i, e.Position, m.Position, n.Position)
				}
			}
		}
	}
}

func TestGridDistances(t *testing.T) {
	g := ticTacToeGrid(t)
	var include = func(g *grid.GridEntry[square]) bool {
		return g.Present
	}
	distances, err := g.Distances(include)
	if err != nil {
		t.Fatalf("unexpected error %s", err)
	}
	if len(distances) != 6 {
		t.Fatalf("expected 6 distances, got %d", len(distances))
	}
	expected := []struct {
		from *grid.GridEntry[square]
		to   map[*grid.GridEntry[square]]int
	}{
		{
			from: g.Get(grid.Position{Row: 0, Col: 0}),
			to: map[*grid.GridEntry[square]]int{
				g.Get(grid.Position{Row: 0, Col: 2}): 2,
				g.Get(grid.Position{Row: 1, Col: 1}): 2,
				g.Get(grid.Position{Row: 1, Col: 2}): 3,
				g.Get(grid.Position{Row: 2, Col: 0}): 2,
				g.Get(grid.Position{Row: 2, Col: 2}): 4,
			},
		},
		{
			from: g.Get(grid.Position{Row: 1, Col: 2}),
			to: map[*grid.GridEntry[square]]int{
				g.Get(grid.Position{Row: 0, Col: 0}): 3,
				g.Get(grid.Position{Row: 0, Col: 2}): 1,
				g.Get(grid.Position{Row: 1, Col: 1}): 1,
				g.Get(grid.Position{Row: 2, Col: 0}): 3,
				g.Get(grid.Position{Row: 2, Col: 2}): 1,
			},
		},
		{
			from: g.Get(grid.Position{Row: 2, Col: 0}),
			to: map[*grid.GridEntry[square]]int{
				g.Get(grid.Position{Row: 0, Col: 0}): 2,
				g.Get(grid.Position{Row: 0, Col: 2}): 4,
				g.Get(grid.Position{Row: 1, Col: 1}): 2,
				g.Get(grid.Position{Row: 1, Col: 2}): 3,
				g.Get(grid.Position{Row: 2, Col: 2}): 2,
			},
		},
		{
			from: g.Get(grid.Position{Row: 2, Col: 2}),
			to: map[*grid.GridEntry[square]]int{
				g.Get(grid.Position{Row: 0, Col: 0}): 4,
				g.Get(grid.Position{Row: 0, Col: 2}): 2,
				g.Get(grid.Position{Row: 1, Col: 1}): 2,
				g.Get(grid.Position{Row: 1, Col: 2}): 1,
				g.Get(grid.Position{Row: 2, Col: 0}): 2,
			},
		},
		{
			from: g.Get(grid.Position{Row: 0, Col: 0}),
			to: map[*grid.GridEntry[square]]int{
				g.Get(grid.Position{Row: 0, Col: 2}): 2,
				g.Get(grid.Position{Row: 1, Col: 1}): 2,
				g.Get(grid.Position{Row: 1, Col: 2}): 3,
				g.Get(grid.Position{Row: 2, Col: 0}): 2,
				g.Get(grid.Position{Row: 2, Col: 2}): 4,
			},
		},
		{
			from: g.Get(grid.Position{Row: 0, Col: 2}),
			to: map[*grid.GridEntry[square]]int{
				g.Get(grid.Position{Row: 0, Col: 0}): 2,
				g.Get(grid.Position{Row: 1, Col: 1}): 2,
				g.Get(grid.Position{Row: 1, Col: 2}): 1,
				g.Get(grid.Position{Row: 2, Col: 0}): 4,
				g.Get(grid.Position{Row: 2, Col: 2}): 2,
			},
		},
		{
			from: g.Get(grid.Position{Row: 1, Col: 1}),
			to: map[*grid.GridEntry[square]]int{
				g.Get(grid.Position{Row: 0, Col: 0}): 2,
				g.Get(grid.Position{Row: 0, Col: 2}): 2,
				g.Get(grid.Position{Row: 1, Col: 2}): 1,
				g.Get(grid.Position{Row: 2, Col: 0}): 2,
				g.Get(grid.Position{Row: 2, Col: 2}): 2,
			},
		},
	}
	for _, exp := range expected {
		for _, edge := range distances[exp.from] {
			dist := exp.to[edge.Entry]
			if dist != edge.Weight {
				t.Errorf("expected dist(%v, %v) to be %d but got %d",
					exp.from.Position, edge.Entry.Position, dist, edge.Weight)
			}
		}
	}
}

//go:embed grid_test/tictactoe.txt
var ticTacToeString string
var (
	ticTacToeHeight int
	ticTacToeWidth  int
	ticTacToeLines  []string
)

func init() {
	ticTacToeString = strings.ReplaceAll(ticTacToeString, "\t", "")
	ticTacToeLines = strings.Split(ticTacToeString, "\n")
	ticTacToeHeight = len(ticTacToeLines)
	ticTacToeWidth = len(ticTacToeLines[0])
}

type square rune

func (s square) String() string {
	return string(s)
}

func ticTacToeGrid(t *testing.T) *grid.Grid[square] {
	t.Helper()
	g := grid.NewGrid[square]()
	for _, r := range ticTacToeString {
		if r == '\n' {
			g.NextRow()
		} else if r == 'X' || r == 'O' {
			g.Add(square(r), true)
		} else if !unicode.IsSpace(r) {
			g.Add(square(r), false)
		}
	}
	if h := g.Height(); h != ticTacToeHeight {
		t.Fatalf("expected grid to have height=%d but got %d", ticTacToeHeight, h)
	}
	if w := g.Width(); w != ticTacToeWidth {
		t.Fatalf("expected grid to have width=%d but got %d", ticTacToeWidth, w)
	}
	return g
}

func get(t *testing.T, g *grid.Grid[square], row, col int) *grid.GridEntry[square] {
	t.Helper()
	pos := grid.Position{Row: row, Col: col}
	entry, err := g.At(pos)
	if err != nil {
		t.Fatalf("expected grid to have entry at %s but did not ", pos)
	}
	return entry
}
