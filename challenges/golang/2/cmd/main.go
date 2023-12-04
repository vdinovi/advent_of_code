package main

import (
	"fmt"
	"os"

	soln "github.com/vdinovi/advent_of_code_2023/challenges/golang/2/solution"
)

func main() {
	records, err := soln.Parse(os.Stdin)
	if err != nil {
		panic(err)
	}
	for _, r := range records {
		fmt.Printf("Game %d: ", r.Game)
		for _, note := range r.Notes {
			fmt.Printf("%d red, %d blue, %d green; ", note[soln.Red], note[soln.Blue], note[soln.Green])
		}
		fmt.Println()
	}
	cubes := [3]int{12, 14, 13}
	sum := 0
	for _, record := range records {
		if record.Possible(cubes) {
			sum += record.Game
		}
	}
	fmt.Printf("answer = %d\n", sum)

	// answer, err := soln.SolveP1(os.Stdin)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "error: %s\n", err)
	// 	os.Exit(1)
	// }
	// fmt.Printf("Answer 1 = %d\n", answer)

	// answer, err := soln.SolveP2(os.Stdin)
	// if err != nil {
	//	fmt.Fprintf(os.Stderr, "error: %s\n", err)
	//	os.Exit(1)
	// }
	// fmt.Printf("Answer 2 = %d\n", answer)
}
