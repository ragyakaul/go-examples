package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Create a function to check all the files that contain duplicates.
// If any duplicates exist, add that file to a duplicateFiles slice
// Print out all the duplicate files in the end

func main() {
	filePaths := os.Args[1:]

	if len(filePaths) <= 0 {
		log.Fatal("Please add files to check for duplicates")
	}

	duplicateFiles := make([]string, 0)

	for _, file := range filePaths {
		fileData, err := os.ReadFile(file) // returns byte slice of file contents
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file %s: %v\n", file, err)
		}
		lineCounter := make(map[string]int)

		lines := strings.Split(string(fileData), "\n")
		for _, line := range lines {
			if lineCounter[line] > 0 {
				duplicateFiles = append(duplicateFiles, file)
				break
			}
			lineCounter[line]++
		}
	}

	if duplicateFiles == nil {
		fmt.Println("No duplicate files found")
	} else {
		fmt.Println("Duplicate files:")
		for _, file := range duplicateFiles {
			fmt.Println(file)
		}
	}
}

/*

file1.txt
apple
apple
banana
orange

counter:
apple: 1

duplicateFiles
["file1.txt"]
*/
