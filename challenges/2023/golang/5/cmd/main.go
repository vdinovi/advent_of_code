package main

import (
	"fmt"
	"os"

	soln "github.com/vdinovi/advent_of_code/challenges/2023/golang/5/solution"
)

func main() {
	// answer, err := soln.SolveP1(os.Stdin)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "error: %s\n", err)
	// 	os.Exit(1)
	// }
	// fmt.Printf("Answer 1 = %d\n", answer)

	answer, err := soln.SolveP2(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Answer 2 = %d\n", answer)
}
