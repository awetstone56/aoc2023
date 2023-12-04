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
			}
		}

		if isPossible {
			possibleGameIdSum += gameId
		}
	}
	fmt.Println(possibleGameIdSum)
}
