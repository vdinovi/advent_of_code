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
	sequences, mappings, err := scan(input)
	if err != nil {
		return 0, err
	}
	seeds, ok := sequences["seeds"]
	if !ok {
		return 0, fmt.Errorf("sequence 'seeds' not found")
	}
	mappings = invert(mappings)
	minLoc := 0
	// TODO: implement a non-brute force
	maxLoc := 10834441
	log := int(maxLoc / 100)
	lowest := math.MaxInt64
	for loc := minLoc; loc < maxLoc; loc += 1 {
		if loc%log == 0 {
			fmt.Printf("> %d%%\n", loc/log)
		}
		seed, err := traverse(loc, "location", "seed", mappings)
		// fmt.Printf("location(%d) -> seed(%d)\n", loc, seed)
		if err != nil {
			return 0, err
		}
		for i := 0; i < len(seeds.elems); i += 2 {
			if seed >= seeds.elems[i] && seed < seeds.elems[i]+seeds.elems[i+1] {
				// fmt.Printf("lowest = min(%d, %d)\n", lowest, loc)
				lowest = min(lowest, loc)
			}
		}
	}
	return Answer(lowest), nil
}

func traverse(value int, source, destination string, mappings map[string]*mapping) (int, error) {
	for source != destination {
		m, ok := mappings[source]
		if !ok {
			return 0, fmt.Errorf("map not found for source %q", source)
		}
		//fmt.Printf("%s(%d) -> %s(%d)\n", source, value, m.destination, m.get(value))
		source, value = m.destination, m.get(value)
	}
	return value, nil
}

func invert(mappings map[string]*mapping) map[string]*mapping {
	inverted := make(map[string]*mapping, len(mappings))
	for _, m := range mappings {
		inv := &mapping{
			source:      m.destination,
			destination: m.source,
			trans:       make([][3]int, len(m.trans)),
		}
		for i := range m.trans {
			inv.trans[i] = [3]int{
				m.trans[i][1],
				m.trans[i][0],
				m.trans[i][2],
			}
		}
		inverted[m.destination] = inv
	}
	return inverted
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
