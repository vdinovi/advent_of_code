package solution

import (
	"bufio"
	"io"
	"math"
	"slices"
	"unicode"
)

type Input io.Reader
type Answer int

func Solve(input Input) (Answer, error) {
	sum := 0
	sc := bufio.NewScanner(input)
	sc.Split(bufio.ScanLines)
	for sc.Scan() {
		digits := []int{-1, -1}
		line := sc.Text()
		for _, ch := range line {
			if unicode.IsDigit(ch) {
				n := int(ch - '0')
				if digits[0] < 0 {
					digits[0] = n
				}
				digits[1] = n
			}
		}
		slices.Reverse(digits)
		for i, digit := range digits {
			sum += digit * int(math.Pow10(i))
		}
	}
	return Answer(sum), sc.Err()
}
