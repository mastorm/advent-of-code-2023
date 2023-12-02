package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	f, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	numberLookup := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	input := strings.Split(string(f), "\r\n")

	acc := 0

	for _, line := range input {
		numbers := make([]int, 0)

		for i := 0; i < len(line); i++ {
			if unicode.IsNumber(rune(line[i])) {
				numbers = append(numbers, int(line[i])-'0')
			} else {
				
				word := ""
				
				for numberLookup. {
					word += string(line[i])
					i++
				}
				numbers = append(numbers, numberLookup[word])
			}
		}

		numbers = append(numbers, 1)

		acc += (numbers[0] + numbers[len(numbers)-1])
	}
	fmt.Println(acc)
}
