package main

import (
	"fmt"
	"os"

	soln "github.com/vdinovi/advent_of_code_2023/challenges/golang/1/solution"
	"github.com/vdinovi/advent_of_code_2023/lib/golang/lib"
)

func main() {
	str, err := lib.ReadAllAsString(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	input := soln.Input(str)
	answer, err := soln.Solve(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(answer)
}
