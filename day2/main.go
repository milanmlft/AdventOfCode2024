package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filePath := "input.txt"
	input, err := readInput(filePath)
	if err != nil {
		panic(err)
	}

	nSafe := 0
	nDampedSafe := 0
	for _, r := range input {
		r.printReport()
		if r.isSafe() {
			nSafe++
			nDampedSafe++
			fmt.Printf(" : Safe")
		} else {
			if r.isDampedSafe() {
				nDampedSafe++
				fmt.Printf(" : Damped Safe")
			}
		}
		fmt.Println()
	}
	fmt.Println("Part 1 result: ", nSafe)
	fmt.Println("Part 2 result: ", nDampedSafe)
}

type report struct {
	levels []int
}

func (r report) isSafe() bool {
	isIncreasing := r.levels[1] > r.levels[0]

	for i := 0; i < len(r.levels)-1; i++ {
		diff := r.levels[i+1] - r.levels[i]
		if diff == 0 {
			return false
		}

		if isIncreasing {
			if diff < 0 {
				return false
			}
		} else {
			if diff > 0 {
				return false
			}
			// Take absolute value
			diff = -diff
		}
		if diff > 3 {
			return false
		}

	}
	return true
}

func (r report) isDampedSafe() bool {
	for i := 0; i < len(r.levels); i++ {
		// Check report after removing each level by
		// explicitly creating a new slice and copying over the levels
		dampedLevels := make([]int, 0, len(r.levels)-1)
		dampedLevels = append(dampedLevels, r.levels[:i]...)
		dampedLevels = append(dampedLevels, r.levels[i+1:]...)
		dampedReport := report{dampedLevels}
		if dampedReport.isSafe() {
			return true
		}
	}
	return false
}

func (r report) printReport() {
	for _, v := range r.levels {
		fmt.Printf("%d ", v)
	}
}

func readInput(filePath string) ([]report, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	out := []report{}

	for scanner.Scan() {
		line := scanner.Text()          // Read the line as a string
		entries := strings.Fields(line) // For space/tab-separated data

		report := report{}
		for _, v := range entries {
			entryValue, err := strconv.Atoi(v)
			if err != nil {
				fmt.Printf("Error converting entry %s to integer: %v\n", v, err)
				return nil, err
			}
			report.levels = append(report.levels, entryValue)
		}
		out = append(out, report)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return nil, err
	}
	return out, nil
}
