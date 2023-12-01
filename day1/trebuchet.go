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
		fmt.Println(line)
		var firstNum string
		var lastNum string
		var firstNumFound bool

		for index, x := range line {
			// fmt.Println(index, string(x))
			char := string(x)
			_, err := strconv.Atoi(char)
			if err != nil {
				// not an int; skip
				continue
			}

			if !firstNumFound {
				firstNumFound = true
				firstNum = char
				lastNum = char
			} else {
				lastNum = char
			}

			if index == (len(line) - 1) {
				combinedNum := firstNum + lastNum
				fmt.Println(combinedNum)
				num, err := strconv.Atoi(combinedNum)
				if err != nil {
					log.Fatal(err)
				}
				calSum = calSum + num
			}
		}
	}

	fmt.Printf("The summation of all calibration values is %x", calSum)
}
