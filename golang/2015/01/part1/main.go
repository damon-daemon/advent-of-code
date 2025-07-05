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
	for _, r := range stringified_input {
		if r == '(' {
			floor += 1
		}
		if r == ')' {
			floor -= 1
		}
	}
	fmt.Printf("Floor: %d\n", floor)
}
