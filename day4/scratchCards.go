package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("puzzleInput.txt")
	// file, _ := os.Open("sampleInput.txt") // this answer should be 13
	scanner := bufio.NewScanner(file)

	totalPoints := 0
	for scanner.Scan() {
		line := scanner.Text()

		var matchingNumsMap = map[int]int{} // might want this for part 2 even though idk what it is
		numOfMatches := 0

		gameNums := strings.Split(line, ":")[1]
		gameNumsTokens := strings.Split(gameNums, "|")
		re := regexp.MustCompile(`\d+`)
		winningNums := re.FindAllString(gameNumsTokens[0], -1)
		myNums := re.FindAllString(gameNumsTokens[1], -1)

		for _, numStr := range winningNums {
			num, _ := strconv.Atoi(numStr)
			matchingNumsMap[num] = 0
		}

		for _, numStr := range myNums {
			num, _ := strconv.Atoi(numStr)
			if _, exists := matchingNumsMap[num]; exists {
				matchingNumsMap[num] += 1
				numOfMatches++
			}
		}

		if numOfMatches != 0 {
			cardPoints := int(math.Pow(2, float64(numOfMatches-1)))
			totalPoints += cardPoints
		}
	}

	fmt.Printf("The cards you have are worth %d points.", totalPoints)
	fmt.Println()
}
