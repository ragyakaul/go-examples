package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
)

// Pass an argument to the command line.
// Check if this argument matches the appropriate url format
// if not, manually create the url
// Perform a get request
func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		log.Fatal("Enter a valid url")
	}

	url := args[0]
	containsHttps, err := regexp.MatchString(`^http[s]://.*`, url)
	if err != nil {
		log.Fatal(err)
	}

	if !containsHttps {
		url = "https://" + url
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()

	fmt.Printf("%v", string(data))

	fmt.Printf("status code: %v\n", resp.StatusCode)

}
