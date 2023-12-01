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

	acc := 0
	input := strings.Split(string(f), "\n")

	for _, str := range input {
		left := 0
		right := len(str) - 1

		firstNumber := 0
		lastNumber := 0

		for firstNumber == 0 {
			char := rune(str[left])
			if unicode.IsNumber(char) {
				firstNumber = int(char) - '0'
			} else {
				left++
			}
		}

		for lastNumber == 0 {
			char := rune(str[right])
			if unicode.IsNumber(char) {
				lastNumber = int(char) - '0'
			} else {
				right--
			}
		}

		acc += (firstNumber*10 + lastNumber)
	}

	fmt.Println(acc)
}
