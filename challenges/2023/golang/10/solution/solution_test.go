package solution_test

import (
	_ "embed"
	"strings"
	"testing"

	soln "github.com/vdinovi/advent_of_code/challenges/2023/golang/10/solution"
)

//go:embed solution_test/input.txt
var input_txt string

func TestSolutionP1(t *testing.T) {
	tests := []struct {
		id       string
		input    soln.Input
		expected soln.Answer
	}{
		{
			id: "Sample 1",
			input: strings.NewReader(`.....
			.S-7.
			.|.|.
			.L-J.
			.....`),
			expected: 4,
		},
		{
			id: "Sample 1",
			input: strings.NewReader(`..F7.
			.FJ|.
			SJ.L7
			|F--J
			LJ...`),
			expected: 8,
		},
		{
			id:       "Puzzle 1",
			input:    strings.NewReader(input_txt),
			expected: 6690,
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
			input: strings.NewReader(`.....
			.S-7.
			.|.|.
			.L-J.
			.....`),
			expected: -1,
		},
		{
			id: "Sample 1",
			input: strings.NewReader(`..F7.
			.FJ|.
			SJ.L7
			|F--J
			LJ...`),
			expected: -1,
		},
		{
			id:       "Puzzle 1",
			input:    strings.NewReader(input_txt),
			expected: -1,
		},
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
