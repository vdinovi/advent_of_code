package solution_test

import (
	_ "embed"
	"strings"
	"testing"

	soln "github.com/vdinovi/advent_of_code/challenges/2023/golang/8/solution"
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
			input: strings.NewReader(`RL

			AAA = (BBB, CCC)
			BBB = (DDD, EEE)
			CCC = (ZZZ, GGG)
			DDD = (DDD, DDD)
			EEE = (EEE, EEE)
			GGG = (GGG, GGG)
			ZZZ = (ZZZ, ZZZ)`),
			expected: 2,
		},
		{
			id: "Sample 2",
			input: strings.NewReader(`LLR

			AAA = (BBB, BBB)
			BBB = (AAA, ZZZ)
			ZZZ = (ZZZ, ZZZ)`),
			expected: 6,
		},
		{
			id:       "Puzzle 1",
			input:    strings.NewReader(one_txt),
			expected: 22199,
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
			id:       "Sample 1",
			input:    strings.NewReader(""),
			expected: 0,
		},
		{
			id:       "Puzzle 1",
			input:    strings.NewReader(one_txt),
			expected: 0,
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
