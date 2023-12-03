package main

import (
	"fmt"
	"os"

	soln "github.com/vdinovi/advent_of_code_2023/challenges/golang/1/solution"
)

func main() {
	answer, err := soln.Solve(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(answer)
}
