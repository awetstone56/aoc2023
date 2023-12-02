package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Game struct {
	Id             int
	NumOfReds      int
	NumOfGreens    int
	NumOfBlues     int
	IsPossibleGame bool
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

	// Begin Problem Logic Here
	compNumOfRed := 12
	compNumOfGreen := 13
	compNumOfBlue := 14

	for scanner.Scan() {
		// each line is one game with multiple attempts with each attempt being separated
		// by a ; and the game starts after :
		line := scanner.Text()
		fmt.Print(line)

		// grab the id integer for the game

		// grab the attempt tokens

		// parse the attempts record the highest of each color

		// create a struct with the highest of each color and the respective id
		// and isPossible to true on the struct for initial value

		// append to the struct list with append
	}

	// loop through the struct list
	// 		compare num of reds to comp value; if > than comp value set possible to false and exit
	//		if gets to end of loop, then add the id int value to the sum
	// print the sum for the answer

}
