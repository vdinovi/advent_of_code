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
	start, g, err := scan(input)
	if err != nil {
		return 0, err
	}
	var loop *walker
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
		if loop == nil || w.steps > loop.steps {
			loop = &w
		}
	}
	total := interior(loop.vertices, len(loop.visited))
	return Answer(total), nil
}

// https://en.wikipedia.org/wiki/Pick's_theorem
// interior_points = area - (boundary_points /2) + 1
func interior(vertices []grid.Position, boundaryPoints int) int {
	return area(vertices) - int(boundaryPoints/2) + 1
}

// https://en.wikipedia.org/wiki/Shoelace_formula
// a = 1/2 * SUM(i=0..n-1) (y_i + y_i=1) * (x_i - x_i+1)
func area(vertices []grid.Position) (area int) {
	for i := range vertices {
		j := (i + 1) % len(vertices)
		area += (vertices[i].Col) * (vertices[j].Row)
		area -= (vertices[i].Row) * (vertices[j].Col)
	}
	if area < 0 {
		area *= -1
	}
	return area >> 1
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
	steps    int
	grid     *grid.Grid[tile]
	cur      *grid.GridEntry[tile]
	visited  map[grid.Position]bool
	vertices []grid.Position
}

func (w *walker) step() bool {
	w.visited[w.cur.Position] = true
	switch w.cur.Item {
	case 'F', '7', 'J', 'L', 'S':
		w.vertices = append(w.vertices, w.cur.Position)
	}
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
