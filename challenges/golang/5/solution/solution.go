package solution

import (
	"fmt"
	"io"
	"math"
)

type Input io.Reader
type Answer int

func SolveP1(input Input) (Answer, error) {
	sequences, mappings, err := scan(input)
	if err != nil {
		return 0, err
	}
	seeds, ok := sequences["seeds"]
	if !ok {
		return 0, fmt.Errorf("sequence 'seeds' not found")
	}
	lowest := math.MaxInt64
	for _, seed := range seeds.elems {
		loc, err := traverse(seed, "seed", "location", mappings)
		if err != nil {
			return 0, err
		}
		lowest = min(lowest, loc)
	}
	return Answer(lowest), nil
}

func SolveP2(input Input) (Answer, error) {
	return 0, nil
}

func traverse(value int, source, destination string, mappings map[string]*mapping) (int, error) {
	for source != destination {
		m, ok := mappings[source]
		if !ok {
			return 0, fmt.Errorf("map not found for source %q", source)
		}
		fmt.Printf("%s(%d) -> %s(%d)\n", source, value, m.destination, m.get(value))
		source, value = m.destination, m.get(value)
	}
	return value, nil
}

func scan(input Input) (map[string]*sequence, map[string]*mapping, error) {
	sequences := make(map[string]*sequence)
	mappings := make(map[string]*mapping)
	scanner := newScanner(input)
	for scanner.Scan() {
		if seq := scanner.Sequence(); seq != nil {
			sequences[seq.name] = seq
		} else if mapp := scanner.Mapping(); mapp != nil {
			mappings[mapp.source] = mapp
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}
	return sequences, mappings, nil
}

type sequence struct {
	name  string
	elems []int
}

type mapping struct {
	source      string
	destination string
	trans       [][3]int
}

func (m *mapping) get(value int) int {
	for _, t := range m.trans {
		if value >= t[1] && value < t[1]+t[2] {
			return t[0] + value - t[1]
		}
	}
	return value
}
