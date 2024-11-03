package solution

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strings"

	"github.com/vdinovi/advent_of_code/lib/golang/lib/grid"
)

type Input io.Reader
type Answer int

func SolveP1(input Input) (Answer, error) {
	g, err := scan(input)
	if err != nil {
		return 0, err
	}
	if err := scaleEmptySectors(g, 2); err != nil {
		return 0, err
	}
	distances, err := minimumDistances(g)
	if err != nil {
		return 0, err
	}
	sum := 0
	for _, dist := range distances {
		sum += dist
	}
	return Answer(sum), nil
}

func SolveP2(input Input) (Answer, error) {
	g, err := scan(input)
	if err != nil {
		return 0, err
	}
	if err := scaleEmptySectors(g, 1000000); err != nil {
		return 0, err
	}
	distances, err := minimumDistances(g)
	if err != nil {
		return 0, err
	}
	sum := 0
	for _, dist := range distances {
		sum += dist
	}
	return Answer(sum), nil
}

type sector struct {
	label rune
	id    int
	dist  int
}

func (s sector) String() string {
	if s.label == '#' {
		return fmt.Sprintf("%d", s.id)
	}
	return string(s.label)
}

func minimumDistances(g *grid.Grid[sector]) (map[[2]int]int, error) {
	distances := map[[2]int]int{}
	var (
		iter1 *grid.GridIterator[sector]
		iter2 *grid.GridIterator[sector]
	)
	for iter1 = g.Iterator(); iter1.Next(); {
		from := iter1.Entry()
		if !from.Present {
			continue
		}
		for iter2 = g.Iterator(); iter2.Next(); {
			to := iter2.Entry()
			if !to.Present || from == to {
				continue
			}
			index := [2]int{
				min(from.Item.id, to.Item.id),
				max(from.Item.id, to.Item.id),
			}
			dist, err := distance(g, from, to)
			if err != nil {
				return nil, err
			}
			if _, ok := distances[index]; !ok {
				distances[index] = math.MaxInt
			}
			distances[index] = min(distances[index], dist)
		}
		if err := iter2.Err(); err != nil {
			return nil, err
		}
	}
	return distances, iter1.Err()
}

func distance(g *grid.Grid[sector], from, to *grid.GridEntry[sector]) (dist int, err error) {
	if to.Position.Row >= from.Position.Row {
		for r := from.Position.Row + 1; r <= to.Position.Row; r += 1 {
			entry, err := g.At(grid.Position{Row: r, Col: from.Position.Col})
			if err != nil {
				return 0, err
			}
			dist += entry.Item.dist
		}
	} else {
		for r := from.Position.Row - 1; r >= to.Position.Row; r -= 1 {
			entry, err := g.At(grid.Position{Row: r, Col: from.Position.Col})
			if err != nil {
				return 0, err
			}
			dist += entry.Item.dist
		}
	}
	if to.Position.Col >= from.Position.Col {
		for c := from.Position.Col + 1; c <= to.Position.Col; c += 1 {
			entry, err := g.At(grid.Position{Row: from.Position.Row, Col: c})
			if err != nil {
				return 0, err
			}
			dist += entry.Item.dist
		}
	} else {
		for c := from.Position.Col - 1; c >= to.Position.Col; c -= 1 {
			entry, err := g.At(grid.Position{Row: from.Position.Row, Col: c})
			if err != nil {
				return 0, err
			}
			dist += entry.Item.dist
		}
	}
	return dist, nil
}

func scaleEmptySectors(g *grid.Grid[sector], factor int) error {
	presentRows := make([]bool, g.Height())
	presentCols := make([]bool, g.Width())
	var it *grid.GridIterator[sector]
	for it = g.Iterator(); it.Next(); {
		entry := it.Entry()
		if entry.Present {
			presentRows[entry.Row] = true
			presentCols[entry.Col] = true
		}
	}
	if err := it.Err(); err != nil {
		return err
	}
	for it := g.Iterator(); it.Next(); {
		entry := it.Entry()
		if !presentRows[entry.Row] || !presentCols[entry.Col] {
			entry.Item.dist = factor
		}
	}
	if err := it.Err(); err != nil {
		return err
	}
	return nil
}

func scan(input Input) (g *grid.Grid[sector], err error) {
	id := 1
	g = grid.NewGrid[sector]()
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		for _, r := range line {
			switch r {
			case '#':
				g.Add(sector{label: r, id: id, dist: 1}, true)
				id += 1
			case '.':
				g.Add(sector{label: r, id: 0, dist: 1}, false)
			}
		}
		g.NextRow()
	}
	return g, scanner.Err()
}
