package solution_test

import (
	_ "embed"
	"testing"

	soln "github.com/vdinovi/advent_of_code/challenges/2023/golang/6/solution"
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
			input: soln.Input{
				Times:     []int{7, 15, 30},
				Distances: []int{9, 40, 200},
			},
			expected: 288,
		},
		{
			id: "Puzzle 1",
			input: soln.Input{
				Times:     []int{42, 89, 91, 89},
				Distances: []int{308, 1170, 1291, 1467},
			},
			expected: 3317888,
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
			input: soln.Input{
				Times:     []int{71530},
				Distances: []int{940200},
			},
			expected: 71503,
		},
		{
			id: "Puzzle 1",
			input: soln.Input{
				Times:     []int{42899189},
				Distances: []int{308117012911467},
			},
			expected: 24655068,
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
