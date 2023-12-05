package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	soln "github.com/vdinovi/advent_of_code_2023/challenges/golang/3/solution"
)

func main() {
	buf := bytes.Buffer{}
	r := io.TeeReader(os.Stdin, &buf)
	answer, err := soln.SolveP1(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Answer 1 = %d\n", answer)

	answer, err = soln.SolveP2(&buf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Answer 2 = %d\n", answer)
}
