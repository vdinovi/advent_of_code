package solution

import "math"

type Input struct {
	Times     []int
	Distances []int
}
type Answer int

// distance(t, T) := t * (T - t) => Tt * t^2
func distance(t, T int) int {
	return t * (T - t)
}

// velocity(t, T) := d(distance(t, T))/dt = T * 2t
func velocity(t, T int) int {
	return T - 2*t
}

// optimal(T) := 0 = velocity(t, T) => (1/2)t
func optimal(T float64) float64 {
	return 0.5 * T
}

// solve set D = dist(t, T) and solve for t
func roots(D, T float64) (float64, float64) {
	discr := math.Sqrt(math.Pow(-T, 2) - 4*D)
	r1 := (T - discr) / 2
	r2 := (T + discr) / 2
	return min(r1, r2), max(r1, r2)
}

func SolveP1(input Input) (Answer, error) {
	product := 1
	for i, T := range input.Times {
		D := input.Distances[i]
		r1, r2 := roots(float64(D), float64(T))
		times := int(math.Floor(r2-0.01)) - int(math.Ceil(r1+0.01)) + 1
		product *= times
	}
	return Answer(product), nil
}

func SolveP2(input Input) (Answer, error) {
	T := input.Times[0]
	D := input.Distances[0]
	r1, r2 := roots(float64(D), float64(T))
	times := int(math.Floor(r2-0.01)) - int(math.Ceil(r1+0.01)) + 1
	return Answer(times), nil
}
