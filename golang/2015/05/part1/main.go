package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	nice_string_count := 0
	input, _ := os.Open("input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if !containsForbiddenStrings(line) && containsVowels(line) && containsDoubleLetters(line) {
			nice_string_count++
		}
	}
	fmt.Printf("The number of nice strings is: %d\n", nice_string_count)
}

func containsForbiddenStrings(line string) bool {
	forbiddenStrings := []string{"ab", "cd", "pq", "xy"}
	for _, forbiddforbiddenStrings := range forbiddenStrings {
		if strings.Contains(line, forbiddforbiddenStrings) {
			return true
		}
	}
	return false
}

func containsVowels(line string) bool {
	vowelCount := 0
	for _, char := range line {
		if strings.ContainsRune("aeiou", char) {
			vowelCount++
		}
	}
	return vowelCount >= 3
}

func containsDoubleLetters(line string) bool {
	for i := 0; i < len(line)-1; i++ {
		if line[i] == line[i+1] {
			return true
		}
	}
	return false
}
