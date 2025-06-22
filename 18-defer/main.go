package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func fileDefer(f *os.File) {
	fmt.Printf("Closing file: %s\n", f.Name())
	f.Close()
}

func processFile(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fileDefer(f)
	r := bufio.NewReader(f)
	length := r.Size()
	fmt.Printf("File %s has length %d\n", fileName, length)
}

func main() {
	files := []string{"a.txt", "b.txt"}

	// dangerous because only calls close after all files have been open (resource leaks, too many open files)
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer fileDefer(f)
		r := bufio.NewReader(f) // runs after the loop ends
		length := r.Size()
		fmt.Printf("File %s has length %d\n", file, length)
	}

	// Instead do this: Each file is opened and closed individually in each loop iteration.
	for _, file := range files {
		processFile(file)
	}
}
