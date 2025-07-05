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
	visited := make(map[string]int)
	santaX, santaY := 0, 0
	visited[fmt.Sprintf("%d,%d", santaX, santaY)]++
	roboX, roboY := 0, 0
	visited[fmt.Sprintf("%d,%d", roboX, roboY)]++

	for i, r := range stringifiedInput {
		x, y := 0, 0
		if i%2 == 0 {
			// Santa's turn
			x, y = getLocation(string(r), santaX, santaY)
			santaX, santaY = x, y
		} else {
			// Robo-Santa's turn
			x, y = getLocation(string(r), roboX, roboY)
			roboX, roboY = x, y
		}
		visited[fmt.Sprintf("%d,%d", x, y)]++
	}

	uniqueHouses := len(visited)
	fmt.Printf("Number of unique houses visited: %d\n", uniqueHouses)
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
