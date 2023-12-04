package solution

import (
	"fmt"
	"io"
	"math"
)

type Input io.Reader
type Answer int

func SolveP1(input Input) (Answer, error) {
	sum := 0
	sc := NewDigitScanner(input)
	sc.AddWords(DigitLiterals)
	for sc.Scan() {
		digits := sc.Digits()
		if len(digits) < 1 {
			return Answer(0), fmt.Errorf("line %q produced %d digits", sc.Text(), len(digits))
		}
		digits = []int{digits[0], digits[len(digits)-1]}
		for i, digit := range digits {
			place := len(digits) - i - 1
			sum += digit * int(math.Pow10(place))
		}
	}
	return Answer(sum), sc.Err()
}

func SolveP2(input Input) (Answer, error) {
	sum := 0
	sc := NewDigitScanner(input)
	sc.AddWords(DigitLiterals)
	sc.AddWords(DigitSpellings)
	for sc.Scan() {
		digits := sc.Digits()
		if len(digits) < 1 {
			return Answer(0), fmt.Errorf("line %q produced %d digits", sc.Text(), len(digits))
		}
		digits = []int{digits[0], digits[len(digits)-1]}
		for i, digit := range digits {
			place := len(digits) - i - 1
			sum += digit * int(math.Pow10(place))
		}
	}
	return Answer(sum), sc.Err()
}
