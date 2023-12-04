package solution

import (
	"io"
)

type Input io.Reader
type Answer int

func SolveP1(input Input) (Answer, error) {
	records, err := Parse(input)
	if err != nil {
		return 0, err
	}
	cubes := [3]int{12, 14, 13}
	sum := 0
	for _, record := range records {
		if record.Possible(cubes) {
			sum += record.Game
		}
	}
	return Answer(sum), nil
}

func SolveP2(input Input) (Answer, error) {
	records, err := Parse(input)
	if err != nil {
		return 0, err
	}
	sum := 0
	for _, record := range records {
		product := 1
		for _, cube := range record.Fewest() {
			product *= cube
		}
		sum += product
	}
	return Answer(sum), nil
}
