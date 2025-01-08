package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
    "slices"
)

func main() {
	input, err := readInput("input.txt")
    if err != nil {
        panic(err)
    }
    rules := parseRules(input)
    pages := parsePages(input)

    total := 0
    total2 := 0

    for _, p := range pages {
        fmt.Println("Pages: ", p)
        fmt.Println("In order: ", p.inOrder(rules))
        if p.inOrder(rules) {
            middle := p[(len(p) - 1) / 2]
            total += middle
        } else {
            ordered := p.orderPages(rules)
            fmt.Println("Ordered: ", ordered)
            middle := ordered[(len(ordered) - 1) / 2]
            total2 += middle
        }
    }
    fmt.Println("Part 1 result: ", total)
    fmt.Println("Part 2 result: ", total2)
}

type rules map[int][]int
type pages []int

func (r rules) addRule(key int, value int) {
    v, ok := r[key]
    if ok {
        r[key] = append(v, value)
    } else {
        r[key] = []int{value}
    }
}

func parseRules(input []string) rules {
    out := rules{}
    re := regexp.MustCompile(`\d+\|\d+`)
    for _,line := range input {
        match := re.FindString(line)
        if (len(match) > 0) {
            rule := regexp.MustCompile(`\|`).Split(match, 2)
            key, _ := strconv.Atoi(rule[0])
            value, _ := strconv.Atoi(rule[1])
            out.addRule(key, value)
        }
    }
    return out
}

func (p pages) inOrder(r rules) bool {
    if len(p) == 1 {
        return true
    }
    // Recursively check whether the last page should come before the other pages
    lastPageRules := r[p[len(p) - 1]]
    for i := 0; i < len(p)-1; i++ {
        for _, rule := range lastPageRules {
            if p[i] == rule {
                return false
            }
        }
    }
    // Check remaining pages
    return p[0:len(p) - 1].inOrder(r)
}

func (p pages) orderPages(r rules) pages {
    ordered := make([]int, len(p))
    for i := 0; i < len(p); i++ {
        currentRules := r[p[i]]
        for j := i - 1; j >= 0; j-- {
            if slices.Contains(currentRules, ordered[j]) {
                ordered[j], ordered[i] = ordered[i], ordered[j]
            }
        }
    }
    return ordered
}

func parsePages(input []string) []pages {
    out := []pages{}

    for _, line := range input {
        pages := regexp.MustCompile(",").Split(line, -1)
        if len(pages) > 1 {
            current := make([]int, len(pages))
            for i := range pages {
                current[i], _ = strconv.Atoi(pages[i])
            }
            out = append(out, current)
        }
    }
    return out
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
