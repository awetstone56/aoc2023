package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// open input file
	// file, err := os.Open("puzzleInput.txt")
	file, err := os.Open("sampleInput.txt") // this answer should be 13
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

	totalPoints := 0
	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(line)

		var matchingNumsMap = map[int]int{}

		gameNums := strings.Split(line, ":")[1]
		gameNumsTokens := strings.Split(gameNums, "|")
		winningNums := strings.Split(strings.TrimSpace(gameNumsTokens[0]), " ")
		myNums := strings.Split(strings.TrimSpace(gameNumsTokens[1]), " ")

		for _, numStr := range winningNums {
			num, _ := strconv.Atoi(numStr)
			matchingNumsMap[num] = 0
		}

		for _, numStr := range myNums {
			num, _ := strconv.Atoi(numStr)
			if _, exists := matchingNumsMap[num]; exists {
				matchingNumsMap[num] = 1
				fmt.Printf("num: %d, matchingNumsMap[num]: %d, ", num, matchingNumsMap[num])
			}
		}
		fmt.Println()

		numOfMatches := 0
		for key, value := range matchingNumsMap {
			fmt.Print(key)
			fmt.Printf(": %d, ", value)
			if value > 0 {
				numOfMatches++
			}
		}

		fmt.Println()
		fmt.Printf("numOfMatches: %d", numOfMatches)
		fmt.Println()
		if numOfMatches != 0 {
			cardPoints := int(math.Pow(2, float64(numOfMatches-1)))
			totalPoints += cardPoints
			fmt.Printf("Card Points: %d, New Total Points: %d", cardPoints, totalPoints)
			fmt.Println()
		}
	}

	fmt.Println("The cards you have are worth:", totalPoints)
}
