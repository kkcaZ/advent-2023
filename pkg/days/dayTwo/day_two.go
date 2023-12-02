package dayTwo

import (
	"fmt"
	"github.com/kkcaZ/advent-2024/pkg/domain"
	"github.com/kkcaZ/advent-2024/pkg/util"
	"strconv"
	"strings"
)

type service struct {
}

func New() domain.DayService {
	return &service{}
}

func (d service) PartOne() error {
	games, err := getGames()
	if err != nil {
		return err
	}

	gameNumber := 1
	possibleGames := make([]int, 0)

	colouredCubes := make(map[string]int)
	colouredCubes["red"] = 12
	colouredCubes["green"] = 13
	colouredCubes["blue"] = 14

	for _, game := range games {
		if game == "" {
			continue
		}

		possible := true

		game = game[strings.Index(game, ":")+2:]
		rounds := strings.Split(game, ";")

		for _, round := range rounds {
			balls := strings.Split(round, ",")
			for _, ball := range balls {
				ball = strings.TrimLeft(ball, " ")
				vals := strings.Split(ball, " ")
				number, err := strconv.Atoi(vals[0])
				if err != nil {
					panic(err)
				}
				colour := strings.ReplaceAll(vals[1], ",", "")

				if number > colouredCubes[colour] {
					possible = false
					break
				}
			}

			if !possible {
				break
			}
		}

		if possible {
			possibleGames = append(possibleGames, gameNumber)
		}

		gameNumber++
	}

	total := 0
	for _, game := range possibleGames {
		total += game
	}

	fmt.Println("Possible games: ", possibleGames)
	fmt.Println("Total: ", total)
	return nil
}

func (d service) PartTwo() error {
	games, err := getGames()
	if err != nil {
		return err
	}

	powers := make([]int, 0)

	for _, game := range games {
		if game == "" {
			continue
		}

		game = game[strings.Index(game, ":")+2:]
		rounds := strings.Split(game, ";")

		required := make(map[string]int)
		required["red"] = 0
		required["green"] = 0
		required["blue"] = 0

		for _, round := range rounds {
			balls := strings.Split(round, ",")
			for _, ball := range balls {
				ball = strings.TrimLeft(ball, " ")
				vals := strings.Split(ball, " ")
				number, err := strconv.Atoi(vals[0])
				if err != nil {
					panic(err)
				}
				colour := strings.ReplaceAll(vals[1], ",", "")

				if number > required[colour] {
					required[colour] = number
				}
			}
		}

		power := required["red"] * required["green"] * required["blue"]
		powers = append(powers, power)
	}

	total := 0
	for _, power := range powers {
		total += power
	}

	fmt.Println("Total power:", total)
	return nil
}

func getGames() ([]string, error) {
	input, err := util.GetInput(2)
	if err != nil {
		return nil, err
	}

	days := strings.Split(input, "\n")
	return days, err
}
