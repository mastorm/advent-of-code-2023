package main

import (
	"fmt"
	day_3 "github.com/mastorm/advent-of-code-2023/day-3"
	"os"
)

func main() {
	f, err := os.ReadFile("C:\\Projects\\advent-of-code-2023\\day-3\\input.txt")
	if err != nil {
		panic(err)
	}

	contents := string(f)

	result := day_3.Solve(contents)
	fmt.Println(result)
}
