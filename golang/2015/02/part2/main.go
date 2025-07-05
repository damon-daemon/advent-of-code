package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	ribbonLength := 0 // in feet

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "x")

		dimensions := make([]int, len(parts))
		for i, part := range parts {
			dimensions[i], err = strconv.Atoi(part)
			if err != nil {
				log.Fatalf("Error converting dimension: %v", err)
			}
		}

		sort.Ints(dimensions)
		ribbonLength += dimensions[0] * 2
		ribbonLength += dimensions[1] * 2
		ribbonLength += dimensions[0] * dimensions[1] * dimensions[2]
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	fmt.Printf("Total required ribbon length: %v feet\n", ribbonLength)
}
