package lib_test

import (
	"strings"
	"testing"
	"unicode"

	"github.com/vdinovi/advent_of_code/lib/golang/lib"
)

type gridItem rune

func (g gridItem) String() string {
	return string(g)
}

func TestGrid(t *testing.T) {
	game := `X.O
	.XO
	O.X
	`
	g := lib.NewGrid[gridItem]()
	for _, r := range game {
		if r == '\n' {
			g.NextRow()
		} else if r == 'X' || r == 'O' {
			g.Add(gridItem(r), true)
		} else if !unicode.IsSpace(r) {
			g.Add(gridItem(r), false)
		}
	}
	if h := g.Height(); h != 3 {
		t.Fatalf("expected height of 3, got %d", h)
	}
	if w := g.Width(); w != 3 {
		t.Fatalf("expected width of 3, got %d", w)
	}
	pos := lib.Position{Row: 0, Col: 0}
	for _, r := range game {
		if r == '\n' {
			pos.Row += 1
			pos.Col = 0
		} else if !unicode.IsSpace(r) {
			if entry, err := g.At(pos); err != nil {
				t.Errorf("unexpected error %s", err)
			} else if entry == nil {
				t.Errorf("expected entry at %v to be present but got nil", pos)
			} else if entry.Item != gridItem(r) {
				t.Errorf("expected entry at %v to be %c but got %c", pos, r, entry.Item)
			}
			pos.Col += 1
		}
	}
	for r := -1; r <= 4; r += 1 {
		for c := -1; c <= 4; c += 1 {
			pos := lib.Position{Row: r, Col: c}
			_, err := g.At(pos)
			if r < 0 || r > 2 || c < 0 || c > 2 {
				if err == nil {
					t.Errorf("expected error at %v but got nil", pos)
				} else if _, ok := err.(*lib.InvalidPositionError); !ok {
					t.Errorf("expected error at %v to be InvalidPositionError but got %s", pos, err)
				}
			} else if err != nil {
				t.Errorf("unexpected error %s", err)
			}
		}
	}
	game = strings.Replace(game, "\t", "", -1)
	if str := g.String(); str != game {
		t.Errorf("expected grid to have string %q but got %q", game, str)
	}
}

func TestGridIterator(t *testing.T) {
	game := `X.O
	.XO
	O.X
	`
	g := lib.NewGrid[gridItem]()
	for _, r := range game {
		if r == '\n' {
			g.NextRow()
		} else if r == 'X' || r == 'O' {
			g.Add(gridItem(r), true)
		} else if !unicode.IsSpace(r) {
			g.Add(gridItem(r), false)
		}
	}
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
	game = strings.Replace(game, "\t", "", -1)
	if str := g.String(); str != game {
		t.Errorf("expected iterator to have produced %q but got %q", game, str)
	}
}

func TestGridDistances(t *testing.T) {
	game := `X.O
	.XO
	O.X
	`
	g := lib.NewGrid[gridItem]()
	for _, r := range game {
		if r == '\n' {
			g.NextRow()
		} else if r == 'X' || r == 'O' {
			g.Add(gridItem(r), true)
		} else if !unicode.IsSpace(r) {
			g.Add(gridItem(r), false)
		}
	}
	var include = func(g *lib.GridEntry[gridItem]) bool {
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
		from *lib.GridEntry[gridItem]
		to   map[*lib.GridEntry[gridItem]]int
	}{
		{
			from: g.Get(lib.Position{Row: 0, Col: 0}),
			to: map[*lib.GridEntry[gridItem]]int{
				g.Get(lib.Position{Row: 0, Col: 2}): 2,
				g.Get(lib.Position{Row: 1, Col: 1}): 2,
				g.Get(lib.Position{Row: 1, Col: 2}): 3,
				g.Get(lib.Position{Row: 2, Col: 0}): 2,
				g.Get(lib.Position{Row: 2, Col: 2}): 4,
			},
		},
		{
			from: g.Get(lib.Position{Row: 1, Col: 2}),
			to: map[*lib.GridEntry[gridItem]]int{
				g.Get(lib.Position{Row: 0, Col: 0}): 3,
				g.Get(lib.Position{Row: 0, Col: 2}): 1,
				g.Get(lib.Position{Row: 1, Col: 1}): 1,
				g.Get(lib.Position{Row: 2, Col: 0}): 3,
				g.Get(lib.Position{Row: 2, Col: 2}): 1,
			},
		},
		{
			from: g.Get(lib.Position{Row: 2, Col: 0}),
			to: map[*lib.GridEntry[gridItem]]int{
				g.Get(lib.Position{Row: 0, Col: 0}): 2,
				g.Get(lib.Position{Row: 0, Col: 2}): 4,
				g.Get(lib.Position{Row: 1, Col: 1}): 2,
				g.Get(lib.Position{Row: 1, Col: 2}): 3,
				g.Get(lib.Position{Row: 2, Col: 2}): 2,
			},
		},
		{
			from: g.Get(lib.Position{Row: 2, Col: 2}),
			to: map[*lib.GridEntry[gridItem]]int{
				g.Get(lib.Position{Row: 0, Col: 0}): 4,
				g.Get(lib.Position{Row: 0, Col: 2}): 2,
				g.Get(lib.Position{Row: 1, Col: 1}): 2,
				g.Get(lib.Position{Row: 1, Col: 2}): 1,
				g.Get(lib.Position{Row: 2, Col: 0}): 2,
			},
		},
		{
			from: g.Get(lib.Position{Row: 0, Col: 0}),
			to: map[*lib.GridEntry[gridItem]]int{
				g.Get(lib.Position{Row: 0, Col: 2}): 2,
				g.Get(lib.Position{Row: 1, Col: 1}): 2,
				g.Get(lib.Position{Row: 1, Col: 2}): 3,
				g.Get(lib.Position{Row: 2, Col: 0}): 2,
				g.Get(lib.Position{Row: 2, Col: 2}): 4,
			},
		},
		{
			from: g.Get(lib.Position{Row: 0, Col: 2}),
			to: map[*lib.GridEntry[gridItem]]int{
				g.Get(lib.Position{Row: 0, Col: 0}): 2,
				g.Get(lib.Position{Row: 1, Col: 1}): 2,
				g.Get(lib.Position{Row: 1, Col: 2}): 1,
				g.Get(lib.Position{Row: 2, Col: 0}): 4,
				g.Get(lib.Position{Row: 2, Col: 2}): 2,
			},
		},
		{
			from: g.Get(lib.Position{Row: 1, Col: 1}),
			to: map[*lib.GridEntry[gridItem]]int{
				g.Get(lib.Position{Row: 0, Col: 0}): 2,
				g.Get(lib.Position{Row: 0, Col: 2}): 2,
				g.Get(lib.Position{Row: 1, Col: 2}): 1,
				g.Get(lib.Position{Row: 2, Col: 0}): 2,
				g.Get(lib.Position{Row: 2, Col: 2}): 2,
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
