package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	requiredWrappingPaper := 0 // in square feet
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "x")
		length, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("Error converting length: %v", err)
		}
		width, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("Error converting width: %v", err)
		}
		height, err := strconv.Atoi(parts[2])
		if err != nil {
			log.Fatalf("Error converting height: %v", err)
		}
		areaLength := 2 * length * width
		areaWidth := 2 * width * height
		areaHeight := 2 * height * length
		areaSmallestSide := min(areaLength/2, areaWidth/2, areaHeight/2)
		requiredWrappingPaper += areaLength + areaWidth + areaHeight + areaSmallestSide
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	fmt.Printf("Total required wrapping paper: %d square feet\n", requiredWrappingPaper)
}
