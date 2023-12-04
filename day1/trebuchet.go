package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var numMap = map[string]string{
	"zero":  "z0o",
	"one":   "o1e",
	"two":   "t2e",
	"three": "t3e",
	"four":  "f4r",
	"five":  "5e",
	"six":   "s6x",
	"seven": "s7n",
	"eight": "e8t",
	"nine":  "n9e",
}

func main() {
	// open input file
	file, err := os.Open("puzzleInput.txt")
	// file, err := os.Open("smallPuzzleInput.txt") // this answer should be 585
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

		fmt.Print(line + " => ")
		// part two logic
		// can replace numbers spelt out with the actual number char,
		// then go for the for loop I built for first part
		line = replaceWordsWithNums(line)
		fmt.Println(line)

		firstNum := ""
		lastNum := ""
		firstNumFound := false
		for index, x := range line {
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

// replaceSpelledOutNumbers replaces spelled-out numbers in the input string with numeric values
func replaceWordsWithNums(line string) string {

	keys := make([]string, 0, len(numMap))
	for k := range numMap {
		keys = append(keys, k)
		regexStr := "(" + strings.Join(keys, "|") + ")"
		re := regexp.MustCompile("(" + regexStr + "|\\d+)")
		matches := re.FindAllString(line, -1) // -1 indicates get ALL matches not a specific number of them
		for _, match := range matches {
			if replacementValue, exists := numMap[match]; exists {
				// getLastChar is to deal with edge case of a letter being needed for the next match..
				// this might break it though in other edge cases..
				line = strings.ReplaceAll(line, match, replacementValue)
			}
		}
	}

	return line
}
