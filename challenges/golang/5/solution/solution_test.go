package solution_test

import (
	_ "embed"
	"strings"
	"testing"

	soln "github.com/vdinovi/advent_of_code_2023/challenges/golang/5/solution"
)

//go:embed solution_test/1.txt
var one_txt string

func TestSolutionP1(t *testing.T) {
	tests := []struct {
		id       string
		input    soln.Input
		expected soln.Answer
	}{
		{
			id: "Sample 1",
			input: strings.NewReader(`seeds: 79 14 55 13

			seed-to-soil map:
			50 98 2
			52 50 48

			soil-to-fertilizer map:
			0 15 37
			37 52 2
			39 0 15

			fertilizer-to-water map:
			49 53 8
			0 11 42
			42 0 7
			57 7 4

			water-to-light map:
			88 18 7
			18 25 70

			light-to-temperature map:
			45 77 23
			81 45 19
			68 64 13

			temperature-to-humidity map:
			0 69 1
			1 0 69

			humidity-to-location map:
			60 56 37
			56 93 4`),
			expected: 35,
		},
		{
			id:       "Puzzle 1",
			input:    strings.NewReader(one_txt),
			expected: 836040384,
		},
	}

	for _, test := range tests {
		answer, err := soln.SolveP1(test.input)
		if err != nil {
			t.Errorf("unexpected error: %s", err)
			continue
		}
		if answer != test.expected {
			t.Errorf("expected test %q to yield answer %v but got %v ", test.id, test.expected, answer)
		}
	}
}

func TestSolutionP2(t *testing.T) {
	tests := []struct {
		id       string
		input    soln.Input
		expected soln.Answer
	}{
		{
			id: "Sample 1",
			input: strings.NewReader(`seeds: 79 14 55 13

			seed-to-soil map:
			50 98 2
			52 50 48
			
			soil-to-fertilizer map:
			0 15 37
			37 52 2
			39 0 15
			
			fertilizer-to-water map:
			49 53 8
			0 11 42
			42 0 7
			57 7 4
			
			water-to-light map:
			88 18 7
			18 25 70
			
			light-to-temperature map:
			45 77 23
			81 45 19
			68 64 13
			
			temperature-to-humidity map:
			0 69 1
			1 0 69
			
			humidity-to-location map:
			60 56 37
			56 93 4`),
			expected: 0,
		},
		// {
		// 	id:       "Puzzle 1",
		// 	input:    strings.NewReader(one_txt),
		// 	expected: 0,
		// },
	}

	for _, test := range tests {
		answer, err := soln.SolveP2(test.input)
		if err != nil {
			t.Errorf("unexpected error: %s", err)
			continue
		}
		if answer != test.expected {
			t.Errorf("expected test %q to yield answer %v but got %v ", test.id, test.expected, answer)
		}
	}
}
