package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	scanner := bufio.NewScanner(file)

	var calSum int
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Print(line)
		firstNum := ""
		lastNum := ""
		firstNumFound := false
		for index, x := range line {
			// fmt.Println(index, string(x))
			char := string(x)
			_, err := strconv.Atoi(char)
			if err == nil {
				// is an int
				if !firstNumFound {
					firstNumFound = true
					firstNum = char
				}
				lastNum = char
			}

			if index == (len(line) - 1) {
				combinedNum := fmt.Sprintf("%s%s", firstNum, lastNum)
				// fmt.Println(combinedNum)
				num, err := strconv.Atoi(combinedNum)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("Adding %d to %d", num, calSum)
				calSum += num
				firstNumFound = false
			}
		}
		fmt.Printf(" = %d", calSum)
		fmt.Println()
	}
}
