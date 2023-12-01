package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// open input file
	file, err := os.Open("puzzleInput.txt")
	if err != nil {
		log.Fatal(err)
	}

	// attempt to close file and log if errors
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// b, err := io.ReadAll(file)
	// fmt.Println(b)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() { // internally, it advances token based on sperator
		fmt.Println(scanner.Text()) // token in unicode-char
		// fmt.Println(scanner.Bytes()) // token in bytes
	}

}
