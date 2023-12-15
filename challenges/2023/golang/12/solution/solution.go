package solution

import (
	"fmt"
	"io"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/vdinovi/advent_of_code/lib/golang/lib"
)

type Input io.Reader
type Answer int

func SolveP1(input Input) (Answer, error) {
	records, err := scan(input)
	if err != nil {
		return 0, err
	}
	sum := 0
	for _, r := range records {
		sum += len(r.arrangements())
	}
	return Answer(sum), nil
}

func SolveP2(input Input) (Answer, error) {
	records, err := scan(input)
	if err != nil {
		return 0, err
	}
	repeat := 5
	sum := 0
	for _, r := range records {
		unfolded := r

		unfolded.springs = make([]spring, len(r.springs)*repeat+(repeat-1))
		for start, i := 0, 0; i < repeat; i += 1 {
			start += copy(unfolded.springs[start:], r.springs)
			if i < repeat-1 {
				unfolded.springs[start] = unknown
				start += 1
			}
		}

		unfolded.groups = make([]int, len(r.groups)*repeat)
		for start, i := 0, 0; i < repeat; i += 1 {
			start += copy(unfolded.groups[start:], r.groups)
		}

		sum += len(unfolded.arrangements())
	}

	return Answer(sum), nil
}

type spring rune

const (
	operational spring = '.'
	damaged     spring = '#'
	unknown     spring = '?'
)

type record struct {
	springs []spring
	groups  []int
}

func (rec record) arrangements() [][]spring {
	var unknowns float64
	for _, s := range rec.springs {
		if s == unknown {
			unknowns++
		}
	}
	result := make([][]spring, 0, int(math.Pow(2, unknowns)))
	rec.enumerate(&result, []spring{}, rec.springs, rec.groups)
	return result
}

func (rec *record) enumerate(results *[][]spring, pre, post []spring, groups []int) {
	if len(post) == 0 {
		if rec.valid(pre) {
			*results = append(*results, slices.Clone(pre))
		}
		return
	}
	if len(groups) == 0 {
		post = []spring(strings.ReplaceAll(string(post), "?", "."))
		rec.enumerate(results, append(pre, post...), []spring{}, groups)
		return
	}
	if groups[0] == 0 {
		rec.enumerate(results, pre, post, groups[1:])
		return
	}
	if rec.prune(pre, post, groups) {
		return
	}
	switch post[0] {
	case unknown:
		rec.enumerate(results, append(pre, operational), post[1:], slices.Clone(groups))
		rec.enumerate(results, append(pre, damaged), post[1:], append([]int{groups[0] - 1}, groups[1:]...))
	case operational:
		rec.enumerate(results, append(pre, operational), post[1:], slices.Clone(groups))
	case damaged:
		rec.enumerate(results, append(pre, damaged), post[1:], append([]int{groups[0] - 1}, groups[1:]...))
	}
}

func (r *record) prune(pre, post []spring, groups []int) bool {
	// TODO: implement pruning criteria. We should be able to tell based on the
	//       remaining springs and groups whether or not we can get a possible answer
	return false
}

func (r *record) valid(springs []spring) bool {
	groups := slices.Clone(r.groups)
	i := 0
	prev := operational
	for _, s := range springs {
		switch s {
		case unknown:
			return false
		case damaged:
			if i >= len(groups) {
				return false
			}
			groups[i] -= 1
			if groups[i] < 0 {
				return false
			}
		case operational:
			if prev == damaged {
				i += 1
			}
		}
		prev = s
	}
	for _, g := range groups {
		if g != 0 {
			return false
		}
	}
	return true
}

func scan(in Input) ([]record, error) {
	return lib.ScanLines[record](in, func(line string) (rec record, ok bool, err error) {
		parts := strings.Fields(line)
		if len(parts) != 2 {
			return rec, false, fmt.Errorf("invalid line %q", line)
		}
		rec.springs = make([]spring, len(parts[0]))
		for i, r := range parts[0] {
			switch s := spring(r); s {
			case operational, damaged, unknown:
				rec.springs[i] = s
			default:
				return rec, false, fmt.Errorf("invalid character %c", r)
			}
		}
		groups := strings.Split(parts[1], ",")
		rec.groups = make([]int, len(groups))
		for i, g := range groups {
			n, err := strconv.Atoi(g)
			if err != nil {
				return rec, false, err
			}
			rec.groups[i] = n
		}
		return rec, true, nil
	})
}
