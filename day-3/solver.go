package day_3

import (
	"strconv"
	"strings"
	"unicode"
)

const LineBreak = "\n"

func Solve(input string) int {
	rows := strings.Split(input, LineBreak)
	sum := 0
	for i, row := range rows {
		captureStart := -1
		captureEnd := -1

		for x := range row {
			if unicode.IsDigit(rune(rows[i][x])) {
				if captureStart == -1 {
					captureStart = x
				}
			} else {
				if captureStart != -1 {
					captureEnd = x
				}
			}

			if captureStart != -1 && captureEnd != -1 {
				if hasAdjacentSymbols(rows, i, captureStart, captureEnd) {
					val, err := strconv.Atoi(rows[i][captureStart:captureEnd])

					if err != nil {
						panic(err)
					}

					sum += val
				}

				captureStart = -1
				captureEnd = -1

			}
		}
	}

	return sum
}

func hasAdjacentSymbols(rows []string, row int, start int, end int) bool {
	// define scan perimeter
	searchStart := start - 1

	if searchStart < 0 {
		searchStart = 0
	}

	searchEnd := end + 1

	rowLength := len(rows[row]) - 1
	if searchEnd > rowLength {
		searchEnd = rowLength
	}

	if ContainsSymbol(rows, row, searchStart, searchEnd) {
		return true
	}

	previousRow := row - 1
	if previousRow > 0 {

		if ContainsSymbol(rows, previousRow, searchStart, searchEnd) {
			return true
		}
	}

	nextRow := row + 1
	if nextRow < len(rows)-1 {

		if ContainsSymbol(rows, nextRow, searchStart, searchEnd) {
			return true
		}
	}

	return false
}

func ContainsSymbol(rows []string, previousRow int, searchStart int, searchEnd int) bool {
	relevantChars := rows[previousRow][searchStart:searchEnd]

	for _, c := range relevantChars {
		if IsSymbol(c) {
			return true
		}
	}
	return false
}

func IsSymbol(r rune) bool {
	return !unicode.IsDigit(r) && r != '.'
}
