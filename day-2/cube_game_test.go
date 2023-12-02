package day_2

import "testing"

func TestParseGameCorrectlyParsesGames(t *testing.T) {
	entries := map[string]Game{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green": {
			Id: 1,
			Sets: &[]Set{
				{RedCount: 4, BlueCount: 3, GreenCount: 0},
				{RedCount: 1, BlueCount: 6, GreenCount: 2},
				{RedCount: 0, BlueCount: 0, GreenCount: 2},
			},
		},
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue": {
			Id: 2,
			Sets: &[]Set{
				{RedCount: 0, BlueCount: 1, GreenCount: 2},
				{RedCount: 1, BlueCount: 4, GreenCount: 3},
				{RedCount: 0, BlueCount: 1, GreenCount: 1},
			},
		},
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red": {
			Id: 3,
			Sets: &[]Set{
				{RedCount: 20, BlueCount: 6, GreenCount: 8},
				{RedCount: 4, BlueCount: 5, GreenCount: 13},
				{RedCount: 1, BlueCount: 0, GreenCount: 5},
			},
		},
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red": {
			Id: 4,
			Sets: &[]Set{
				{RedCount: 3, BlueCount: 6, GreenCount: 1},
				{RedCount: 6, BlueCount: 0, GreenCount: 3},
				{RedCount: 14, BlueCount: 15, GreenCount: 3},
			},
		},
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green": {
			Id: 5,
			Sets: &[]Set{
				{RedCount: 6, BlueCount: 1, GreenCount: 3},
				{RedCount: 1, BlueCount: 2, GreenCount: 2},
			},
		},
	}

	for logEntry, expectedGame := range entries {
		game := ParseGame(logEntry)

		if game.Id != expectedGame.Id {
			t.Errorf("Expected game ID: %d; got %d", expectedGame.Id, game.Id)
		}

		gameCount := len(*game.Sets)
		expectedGameCount := len(*expectedGame.Sets)
		if gameCount != expectedGameCount {
			t.Errorf("Expected game to have set count of: %d; got %d", gameCount, expectedGameCount)
		}

		for i, x := range *expectedGame.Sets {
			if (*game.Sets)[i].BlueCount != x.BlueCount {
				t.Errorf("Expected BlueCount to be: %d; got %d", (*game.Sets)[i].BlueCount, x.BlueCount)
			}
			if (*game.Sets)[i].RedCount != x.RedCount {
				t.Errorf("Expected RedCount to be: %d; got %d", (*game.Sets)[i].RedCount, x.RedCount)
			}
			if (*game.Sets)[i].GreenCount != x.GreenCount {
				t.Errorf("Expected GreenCount to be: %d; got %d", (*game.Sets)[i].GreenCount, x.GreenCount)
			}
		}
	}
}
