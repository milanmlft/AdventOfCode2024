package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	filePath := "input.txt"
	input, err := readInput(filePath)
	if err != nil {
		panic(err)
	}

	total := 0
	for _, line := range input {
		instructions := extractInstructions(line)
		for _, instruction := range instructions {
			total += executeInstruction(instruction)
		}
	}
	fmt.Printf("Part 1 solution: %d\n", total)
}

func extractInstructions(line string) []string {
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	return re.FindAllString(line, -1)
}

func executeInstruction(s string) int {
	re := regexp.MustCompile(`\d{1,3}`)
	matches := re.FindAllString(s, 2)
	lhs, err := strconv.Atoi(matches[0])
	if err != nil {
		panic(err)
	}
	rhs, err := strconv.Atoi(matches[1])
	if err != nil {
		panic(err)
	}
	return lhs * rhs
}

func readInput(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	out := []string{}

	for scanner.Scan() {
		line := scanner.Text() // Read the line as a string
		out = append(out, line)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return nil, err
	}
	return out, nil
}
