package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	floor := 0
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	stringified_input := string(input)
	for i, r := range stringified_input {
		if r == '(' {
			floor += 1
		}
		if r == ')' {
			floor -= 1
		}
		if floor == -1 {
			// Need to add 1 to i because we want the index position to be 1-based
			fmt.Printf("Entered basement at index position %v on floor %v\n", i+1, floor)
			return
		}
	}
	fmt.Printf("Never reached basement, on floor %d\n", floor)
}
