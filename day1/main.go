package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type table struct {
	leftCol  []int
	rightCol []int
}

func main() {
	filePath := "input.txt"
	input, err := readInput(filePath)
	if err != nil {
		panic(err)
	}

	// Part 1

	// Sort both columns in ascending order
	sort.Ints(input.leftCol)
	sort.Ints(input.rightCol)

	diff := 0
	total := 0
	for i, left := range input.leftCol {
		diff = left - input.rightCol[i]
		if diff < 0 {
			diff = -diff
		}
		total += diff
	}
	fmt.Println("Part 1 result: ", total)

	// Part 2
	similarity := 0
	for _, left := range input.leftCol {
		occurrence := 0
		for _, right := range input.rightCol {
			// Assuming sorted input
			if right > left {
				break
			}
			if right == left {
				occurrence++
			}
		}
		similarity += left * occurrence
	}

	fmt.Println("Part 2 result: ", similarity)
}

func readInput(filePath string) (*table, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	out := &table{}

	for scanner.Scan() {
		line := scanner.Text()          // Read the line as a string
		columns := strings.Fields(line) // For space/tab-separated data

		firstNumber, err := strconv.Atoi(columns[0])
		if err != nil {
			fmt.Printf("Error converting first number to integer: %v\n", err)
			return nil, err
		}

		secondNumber, err := strconv.Atoi(columns[1])
		if err != nil {
			fmt.Printf("Error converting second number to integer: %v\n", err)
			return nil, err
		}

		out.leftCol = append(out.leftCol, firstNumber)
		out.rightCol = append(out.rightCol, secondNumber)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return nil, err
	}
	return out, nil
}
