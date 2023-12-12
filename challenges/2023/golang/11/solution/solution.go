package solution

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strings"

	"github.com/vdinovi/advent_of_code/lib/golang/lib"
)

type Input io.Reader
type Answer int

func SolveP1(input Input) (Answer, error) {
	grid, err := scan(input)
	if err != nil {
		return 0, err
	}
	scaleEmptyGalaxies(grid)
	distances, err := calculateMinimumDistances(grid)
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
	return 0, nil
}

type sector struct {
	label rune
	id    int
}

func (s sector) String() string {
	if s.label == '#' {
		return fmt.Sprintf("%d", s.id)
	}
	return string(s.label)
}

func calculateMinimumDistances(g *lib.Grid[sector]) (map[[2]int]int, error) {
	distances, err := g.Distances(func(entry *lib.GridEntry[sector]) bool {
		return entry.Present
	})
	if err != nil {
		return nil, err
	}
	shortest := map[[2]int]int{}
	for from, edges := range distances {
		for _, edge := range edges {
			pair := [2]int{
				min(from.Item.id, edge.Entry.Item.id),
				max(from.Item.id, edge.Entry.Item.id),
			}
			if _, ok := shortest[pair]; !ok {
				shortest[pair] = math.MaxInt
			}
			shortest[pair] = min(shortest[pair], edge.Weight)
		}
	}
	return shortest, nil
}

func scaleEmptyGalaxies(g *lib.Grid[sector]) {
	presentRows := make([]bool, g.Height())
	presentCols := make([]bool, g.Width())
	for it := g.Iterator(); it.Next(); {
		entry := it.Entry()
		if entry.Present {
			presentRows[entry.Row] = true
			presentCols[entry.Col] = true
		}
	}
	shift := 0
	for r, presentRow := range presentRows {
		if !presentRow {
			g.InsertRow(r+shift, sector{label: '.', id: 0}, false)
			shift += 1
		}
	}
	shift = 0
	for c, presentCol := range presentCols {
		if !presentCol {
			g.InsertColumn(c+shift, sector{label: '.', id: 0}, false)
			shift += 1
		}
	}
}

func scan(input Input) (grid *lib.Grid[sector], err error) {
	id := 1
	grid = lib.NewGrid[sector]()
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		for _, r := range line {
			switch r {
			case '#':
				grid.Add(sector{label: r, id: id}, true)
				id += 1
			case '.':
				grid.Add(sector{label: r, id: 0}, false)
			}
		}
		grid.NextRow()
	}
	return grid, scanner.Err()
}
