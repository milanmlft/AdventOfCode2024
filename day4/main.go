package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	// Part 1
	nHorizontal := countHorizontalInstances(input)
	nVertical := countVerticalInstances(input)
	nDiagonal := countDiagonalInstances(input)
	total := nHorizontal + nVertical + nDiagonal
	fmt.Println("Result: ", total)

	// Part 2
	fmt.Println("Result part 2:", countCrossMasses(input))
}

func countCrossMasses(lines []string) int {
	total := 0
	lineLengh := len(lines[0])
	// Find all 'A' positions as centre of the cross
	re := regexp.MustCompile("A")
	for l := 1; l < len(lines)-1; l++ {
		locA := re.FindAllStringIndex(lines[l], -1)
		for _, loc := range locA {
			if loc[0] < lineLengh-1 && loc[0] > 0 {
				if detectCrossMass(lines, l, loc[0]) {
					total++
				}
			}
		}
	}
	return total
}

func detectCrossMass(lines []string, line int, loc int) bool {
	tl := lines[line-1][loc-1]
	tr := lines[line-1][loc+1]
	bl := lines[line+1][loc-1]
	br := lines[line+1][loc+1]

	forwardDiagonal := []byte{tl, lines[line][loc], br}
	backwardDiagonal := []byte{bl, lines[line][loc], tr}

	forwardMatchedMAS, _ := regexp.Match("MAS", forwardDiagonal)
	forwardMatchedSAM, _ := regexp.Match("SAM", forwardDiagonal)
	forwardMatched := forwardMatchedMAS || forwardMatchedSAM

	backwardMatchedMAS, _ := regexp.Match("MAS", backwardDiagonal)
	backwardMatchedSAM, _ := regexp.Match("SAM", backwardDiagonal)
	backwardMatched := backwardMatchedMAS || backwardMatchedSAM

	if forwardMatched && backwardMatched {
		return true
	}

	return false
}

func countHorizontalInstances(lines []string) int {
	total := 0
	for _, line := range lines {
		total += countOccurrences(line)
	}
	return total
}

func countVerticalInstances(lines []string) int {
	total := 0

	// Get all vertical lines from input
	lineLenth := len(lines[0])
	for i := range lineLenth {
		column := ""
		for _, line := range lines {
			column += string(line[i])
		}
		total += countOccurrences(column)
	}

	return total
}

func countDiagonalInstances(lines []string) int {
	total := 0

	// Treat the lines as a matrix and extract forward and backward diagonals
	// https://stackoverflow.com/a/43311126/11801854
	nRows := len(lines)
	nCols := len(lines[0])
	nForwardDiagonals := nRows + nCols - 1
	nBackwardDiagonals := nForwardDiagonals
	bDiagonalsOffset := -nRows + 1

	forwardDiagonals := make([]string, nForwardDiagonals)
	backwardDiagonals := make([]string, nBackwardDiagonals)

	for i := range nRows {
		for j := range nCols {
			forwardDiagonals[i+j] += string(lines[i][j])
			backwardDiagonals[i-j-bDiagonalsOffset] += string(lines[i][j])
		}
	}

	for i := range forwardDiagonals {
		total += countOccurrences(forwardDiagonals[i])
		total += countOccurrences(backwardDiagonals[i])

	}

	return total
}

func countOccurrences(line string) int {
	out := 0
	patterns := regexPatterns()
	for _, p := range patterns {
		out += len(p.FindAllString(line, -1))
	}
	return out
}

func regexPatterns() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile("XMAS"),
		regexp.MustCompile("SAMX"),
	}
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
