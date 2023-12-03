package day_3

import "testing"

func TestSolve(t *testing.T) {
	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

	sum := Solve(input)

	if sum != 4361 {
		t.Errorf("Expected sum: %d, got %d", 4361, sum)
	}
}
