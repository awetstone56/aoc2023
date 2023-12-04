package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fileContent, err := os.ReadFile("puzzleInput.txt")
	if err != nil {
		log.Fatal(err)
	}

	// identify symbols in each line then map their points into symbolGrid
	symbolGrid := map[image.Point]rune{}
	for y, line := range strings.Fields(string(fileContent)) {
		for x, char := range line {
			if char != '.' && !unicode.IsDigit(char) {
				symbolGrid[image.Point{x, y}] = char // is a symbol
			}
		}
	}

	// map the byte content into grid again then identify number strings
	// AND subsequently identify the points that border around the number strings
	// THEN we should have all the info we need to identify the actual parts and then sum them together
	enginePartsSum := 0
	// parts := map[image.Point][]int{}
	for y, line := range strings.Fields(string(fileContent)) {
		// match on strings of numbers, then grab their start/end indices to check around each point
		// we won't care about looking at anything inbetween and it doesn't appear that inputs have more/less than 2-3 digits
		for _, potentialPartCoords := range regexp.MustCompile(`\d+`).FindAllStringIndex(line, -1) {
			numberBorder := map[image.Point]struct{}{} // create the border around the number
			for x := potentialPartCoords[0]; x < potentialPartCoords[1]; x++ {
				// loop through, looking all around any given point around the number
				for _, potentialSymbol := range []image.Point{
					{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
				} {
					numberBorder[image.Point{x, y}.Add(potentialSymbol)] = struct{}{} // map the unique point into the boundary
				}
			}

			num, err := strconv.Atoi(line[potentialPartCoords[0]:potentialPartCoords[1]])
			if err != nil {
				fmt.Println("Error while parsing potential part string to an int: ", err)
			}
			for borderPoint := range numberBorder {
				if _, existsInMap := symbolGrid[borderPoint]; existsInMap {
					enginePartsSum += num
				}
			}
		}
	}

	fmt.Println(enginePartsSum)
}
