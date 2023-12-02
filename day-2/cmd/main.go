package main

import (
	"fmt"
	part1 "github.com/mastorm/day2-part1"
	"os"
	"strings"
)

func main() {
	f, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	contents := string(f)
	gameLogs := strings.Split(contents, "\r\n")

	validGames := 0
	setPowers := 0

	for _, logEntry := range gameLogs {
		game := part1.ParseGame(logEntry)

		isValid := true

		largestRed := 0
		largestGreen := 0
		largestBlue := 0

		for _, set := range *game.Sets {
			if set.RedCount > 12 || set.GreenCount > 13 || set.BlueCount > 14 {
				isValid = false
			}

			if set.RedCount > largestRed {
				largestRed = set.RedCount
			}

			if set.GreenCount > largestGreen {
				largestGreen = set.GreenCount
			}

			if set.BlueCount > largestBlue {
				largestBlue = set.BlueCount
			}
		}

		setPowers += largestRed * largestBlue * largestGreen
		if isValid {
			validGames += game.Id
		}

	}

	fmt.Printf("answer for part 1: %d; answer for part 2: %d", validGames, setPowers)
}
