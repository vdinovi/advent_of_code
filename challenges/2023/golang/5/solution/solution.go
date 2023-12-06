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
		loc, err := traverseValue(seed, "seed", "location", mappings)
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
	ranges := make([][2]int, len(seeds.elems)/2)
	for i := 0; i < len(seeds.elems); i += 2 {
		ranges[i/2][0] = seeds.elems[i]
		ranges[i/2][1] = seeds.elems[i] + seeds.elems[i+1]
	}
	result, err := traverseRanges(ranges, "seed", "location", mappings)
	if err != nil {
		return 0, err
	}
	lowest := math.MaxInt
	for _, r := range result {
		lowest = min(lowest, r[0])
	}
	return Answer(lowest), nil
}

func traverseValue(value int, source, destination string, mappings map[string]*mapping) (int, error) {
	for source != destination {
		m, ok := mappings[source]
		if !ok {
			return 0, fmt.Errorf("map not found for source %q", source)
		}
		source, value = m.destination, m.get(value)
	}
	return value, nil
}

func traverseRanges(ranges [][2]int, source, destination string, mappings map[string]*mapping) ([][2]int, error) {
	for source != destination {
		m, ok := mappings[source]
		if !ok {
			return nil, fmt.Errorf("map not found for source %q", source)
		}
		ranges = m.getRanges(ranges)
		source = m.destination
	}
	return ranges, nil
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

func (m *mapping) getRanges(rs [][2]int) [][2]int {
	ranges := make(rangeSet)
	ranges.add(rs...)
	mapped := make(rangeSet)
	for _, t := range m.trans {
		filter := make(rangeSet)
		for r := range ranges {
			if s, e := r[0], min(r[1], t[1]); s < e {
				// left portion not empty, preserve
				filter.add([2]int{s, e})
			}
			if s, e := max(r[0], t[1]), min(t[1]+t[2], r[1]); s < e {
				// overlapping portion not empty, add mapped range
				mapped.add([2]int{s - t[1] + t[0], e - t[1] + t[0]})
			}
			if s, e := max(t[1]+t[2], r[0]), r[1]; s < e {
				// right portion not empty, preserve
				filter.add([2]int{s, e})
			}
		}
		ranges = filter
	}
	ranges.add(mapped.asSlice()...)
	return ranges.asSlice()
}

type rangeSet map[[2]int]bool

func (rs rangeSet) add(ranges ...[2]int) {
	for _, r := range ranges {
		rs[r] = true
	}
}

func (rs rangeSet) asSlice() [][2]int {
	s := make([][2]int, len(rs))
	i := 0
	for r := range rs {
		s[i] = r
		i += 1
	}
	return s
}
