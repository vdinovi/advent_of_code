package solution_test

import (
	_ "embed"
	"strings"
	"testing"

	soln "github.com/vdinovi/advent_of_code/challenges/2023/golang/9/solution"
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
			input: strings.NewReader(`0 3 6 9 12 15
			1 3 6 10 15 21
			10 13 16 21 30 45`),
			expected: 114,
		},
		{
			id:       "Puzzle 1",
			input:    strings.NewReader(input_txt),
			expected: 1974913025,
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
			input: strings.NewReader(`0 3 6 9 12 15
			1 3 6 10 15 21
			10 13 16 21 30 45`),
			expected: 2,
		},
		{
			id:       "Puzzle 1",
			input:    strings.NewReader(input_txt),
			expected: 884,
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
