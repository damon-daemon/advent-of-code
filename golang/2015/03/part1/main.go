package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	stringifiedInput := string(input)

	x, y := 0, 0
	visited := make(map[string]int)
	visited[fmt.Sprintf("%d,%d", x, y)]++

	for _, r := range stringifiedInput {
		x, y = getLocation(string(r), x, y) // Update x and y using returned values
		visited[fmt.Sprintf("%d,%d", x, y)]++
	}

	uniqueHouses := len(visited)
	fmt.Printf("Number of unique houses visited: %d\n", uniqueHouses)
	fmt.Printf("Santa's final location is at (%d, %d)\n", x, y)
}

func getLocation(input string, x, y int) (int, int) {
	switch input {
	case ">":
		x++
	case "<":
		x--
	case "^":
		y++
	case "v":
		y--
	}
	return x, y // Return updated x and y
}
