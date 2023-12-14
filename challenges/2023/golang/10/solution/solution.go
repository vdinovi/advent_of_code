package solution

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"unicode"

	"github.com/vdinovi/advent_of_code/lib/golang/lib/grid"
)

type Input io.Reader
type Answer int

func SolveP1(input Input) (Answer, error) {
	start, g, err := scan(input)
	if err != nil {
		return 0, err
	}
	distances := make([][]int, g.Height())
	for i := range distances {
		distances[i] = make([]int, g.Width())
	}
	var steps int
	for _, move := range possibleMoves(g, start) {
		if move == nil {
			continue
		}
		w := walker{
			grid:    g,
			cur:     move,
			visited: make(map[grid.Position]bool, g.Width()*g.Height()),
		}
		for w.step() {
		}
		steps = max(steps, w.steps)
	}
	if steps%2 == 0 {
		steps /= 2
	} else {
		steps = (steps / 2) + 1
	}
	return Answer(steps), nil
}

func SolveP2(input Input) (Answer, error) {
	return 0, nil
}

func scan(input Input) (start *grid.GridEntry[tile], g *grid.Grid[tile], err error) {
	g = grid.NewGrid[tile]()
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, "\t", "")
		for _, r := range line {
			if !unicode.IsSpace(r) {
				present := true
				if r != '.' {
					present = false
				}
				g.Add(tile(r), present)
				if r == 'S' {
					start = g.Last()
				}
			}
		}
		g.NextRow()
	}
	if start == nil {
		return nil, nil, fmt.Errorf("no start found")
	}
	return start, g, scanner.Err()
}

type direction int

const (
	up direction = iota
	right
	down
	left
)

type tile rune

func (t tile) String() string {
	return string(t)
}

type walker struct {
	steps   int
	grid    *grid.Grid[tile]
	cur     *grid.GridEntry[tile]
	visited map[grid.Position]bool
}

func (w *walker) step() bool {
	w.visited[w.cur.Position] = true
	moves := w.moves()
	for _, m := range moves {
		if m != nil {
			w.cur = m
			w.steps += 1
			return true
		}
	}
	return false
}

func (w *walker) moves() [4]*grid.GridEntry[tile] {
	moves := possibleMoves(w.grid, w.cur)
	for i, n := range moves {
		if n != nil && w.visited[n.Position] {
			moves[i] = nil
		}
	}
	return moves
}

func possibleMoves(g *grid.Grid[tile], from *grid.GridEntry[tile]) [4]*grid.GridEntry[tile] {
	neighbors := g.Neighbors(from)
	for i, n := range neighbors {
		if n != nil && n.Item == '.' {
			neighbors[i] = nil
		}
	}
	switch from.Item {
	case 'S':
		// all neighbors are valid
	case '|':
		neighbors[left] = nil
		neighbors[right] = nil
	case '-':
		neighbors[up] = nil
		neighbors[down] = nil
	case 'L':
		neighbors[left] = nil
		neighbors[down] = nil
	case 'J':
		neighbors[right] = nil
		neighbors[down] = nil
	case '7':
		neighbors[right] = nil
		neighbors[up] = nil
	case 'F':
		neighbors[left] = nil
		neighbors[up] = nil
	}
	return neighbors
}
