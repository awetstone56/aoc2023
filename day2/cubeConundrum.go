package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var colorMap = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

var colorMapTwo = map[string]int{
	"red":   0,
	"green": 0,
	"blue":  0,
}

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
	possibleGameIdSum := 0
	gamePowerSum := 0
	for scanner.Scan() {
		// each line is one game with multiple attempts with each attempt being separated
		// by a ; and the game starts after :
		line := scanner.Text()

		isPossible := true
		gameIdStr := strings.Split(strings.Split(line, " ")[1], ":")[0]
		gameId, err := strconv.Atoi(gameIdStr) // grab the attempt tokens
		if err != nil {
			log.Fatal("Error parsing gameId:", err)
			return
		}
		attemptList := strings.Split(strings.Split(line, ":")[1], ";")

		for _, attempt := range attemptList {
			colorList := strings.Split(attempt, ",")
			for _, color := range colorList {
				colorAryStr := strings.Split(color, " ")
				colorCount, err := strconv.Atoi(strings.TrimSpace(colorAryStr[1]))
				if err != nil {
					log.Fatal("Error parsing color count:", err)
					return
				}
				colorName := colorAryStr[2]

				if colorCount > colorMap[colorName] {
					isPossible = false
				}

				if colorCount > colorMapTwo[colorName] {
					colorMapTwo[colorName] = colorCount
				}
			}
		}

		gamePower := 1
		// Multiply all values in the map
		for _, value := range colorMapTwo {
			gamePower *= value
		}
		fmt.Println(gamePower)
		gamePowerSum += gamePower

		// reset the map for next game
		for key := range colorMapTwo {
			colorMapTwo[key] = 0
		}

		if isPossible {
			possibleGameIdSum += gameId
		}
	}
	fmt.Println("The sum for part 1 is:", possibleGameIdSum)
	fmt.Println("The sum for part 2 is:", gamePowerSum)
}
