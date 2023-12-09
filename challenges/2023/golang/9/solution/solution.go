package solution

import (
	"io"

	"github.com/vdinovi/advent_of_code/lib/golang/lib"
)

type Input io.Reader
type Answer int

func SolveP1(input Input) (Answer, error) {
	histories, err := scan(input)
	if err != nil {
		return 0, err
	}
	sum := 0
	for _, ns := range histories {
		sum += predict(ns, forward)
	}
	return Answer(sum), nil
}

func SolveP2(input Input) (Answer, error) {
	histories, err := scan(input)
	if err != nil {
		return 0, err
	}
	sum := 0
	for _, ns := range histories {
		sum += predict(ns, backward)
	}
	return Answer(sum), nil
}

type direction int

const (
	forward direction = iota
	backward
)

func predict(ns []int, dir direction) (prediction int) {
	if !lib.AllEqual(ns) {
		deltas := make([]int, len(ns)-1)
		for i := 1; i < len(ns); i++ {
			deltas[i-1] = ns[i] - ns[i-1]
		}
		prediction = predict(deltas, dir)
	}
	switch dir {
	case forward:
		return ns[len(ns)-1] + prediction
	case backward:
		return ns[0] - prediction
	default:
		panic("invalid direction")
	}
}

func scan(r io.Reader) (histories [][]int, err error) {
	return lib.ScanLines(r, func(line string) ([]int, bool, error) {
		numbers, err := lib.ParseInts(line)
		return numbers, len(numbers) > 0, err
	})
}
