package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

//	func GetData() string {
//		return `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
//
// Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
// Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
// Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
// Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
// Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`
// }
func GetData() string {
	f, err := os.ReadFile("day-4/input.txt")

	if err != nil {
		panic(err)
	}

	return string(f)
}

func cardContents(line string) ([]int, []int) {
	cleanedLine := line[strings.Index(line, ":")+1:]
	cardSegments := strings.Split(cleanedLine, "|")

	winners := parseNumbers(cardSegments[0])
	choices := parseNumbers(cardSegments[1])

	return winners, choices
}

func parseNumbers(s string) []int {
	numbers := strings.Split(s, " ")
	result := make([]int, 0)

	for _, numericLike := range numbers {
		if numericLike == "" {
			continue
		}

		number, err := strconv.Atoi(numericLike)
		if err != nil {
			panic(err)
		}

		result = append(result, number)
	}

	return result
}

func main() {
	data := GetData()
	lines := strings.Split(data, "\r\n")

	points := 0

	for _, line := range lines {
		winners, choices := cardContents(line)
		points += sumCard(winners, choices)
	}

	fmt.Println(points)

}

func sumCard(winners []int, choices []int) int {
	sum := 0
	for _, choice := range choices {
		for _, winner := range winners {
			if choice == winner {

				if sum == 0 {
					sum = 1
				} else {
					sum *= 2
				}
				break
			}
		}
	}

	return sum
}
