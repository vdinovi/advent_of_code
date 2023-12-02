package solution_test

import (
	"testing"

	soln "github.com/vdinovi/advent_of_code_2023/challenges/golang/1/solution"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		id       string
		input    soln.Input
		expected soln.Answer
	}{
		{
			id: "Sample 1",
			input: `1abc2
			pqr3stu8vwx
			a1b2c3d4e5f
			treb7uchet`,
			expected: 142,
		},
	}

	for _, test := range tests {
		answer, err := soln.Solve(test.input)
		if err != nil {
			t.Errorf("unexpected error: %s", err)

		}
		if answer != test.expected {
			t.Errorf("expected test %q to yield answer %v but got %v ", test.id, test.expected, answer)
		}
	}
}
