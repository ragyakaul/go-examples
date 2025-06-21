package main

import (
	"bufio"
	"fmt"
	"os"
)

// Create a function to count the number of unique lines in a file
// Increment a counter for each unique line
// Accept file name as an command line argument

func main() {
	counter := make(map[string]int)
	args := os.Args[1:]
	if len(args) < 1 {
		countLines(os.Stdin, &counter)
	} else {
		for _, fileName := range args {
			fileHandler, err := os.Open(fileName)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error opening file %s: %v\n", fileName, err)
			}
			countLines(fileHandler, &counter)
		}
	}

	for line, count := range counter {
		fmt.Printf("%s: %d\n", line, count)
	}
}

func countLines(f *os.File, counter *map[string]int) error {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		(*counter)[line]++
	}
	return nil
}
