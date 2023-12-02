package day_2

import (
	"strconv"
	"strings"
)

type Set struct {
	RedCount   int
	GreenCount int
	BlueCount  int
}

type Game struct {
	Id   int
	Sets *[]Set
}

func extractId(logEntry string) (int, error) {
	startOfId := strings.Index(logEntry, " ") + 1
	endOfId := strings.Index(logEntry, ":")

	return strconv.Atoi(logEntry[startOfId:endOfId])
}

func ParseGame(logEntry string) Game {
	// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green

	id, err := extractId(logEntry)

	if err != nil {
		panic(err)
	}

	sets := extractSets(logEntry)
	return Game{
		Id:   id,
		Sets: sets,
	}
}

func extractSets(logEntry string) *[]Set {
	startOfSets := strings.Index(logEntry, ":") + 1
	setLog := logEntry[startOfSets:]

	unparsedSets := strings.Split(setLog, ";")

	sets := make([]Set, len(unparsedSets))

	for i, x := range unparsedSets {
		segments := strings.Split(x, ",")

		for _, segment := range segments {
			parts := strings.Split(strings.TrimLeft(segment, " "), " ")
			switch parts[1] {
			case "blue":
				parsed, err := strconv.Atoi(parts[0])

				if err != nil {
					panic(err)
				}

				sets[i].BlueCount = parsed
			case "green":
				parsed, err := strconv.Atoi(parts[0])

				if err != nil {
					panic(err)
				}

				sets[i].GreenCount = parsed
			case "red":
				parsed, err := strconv.Atoi(parts[0])

				if err != nil {
					panic(err)
				}

				sets[i].RedCount = parsed
			}
		}
	}

	return &sets

}
