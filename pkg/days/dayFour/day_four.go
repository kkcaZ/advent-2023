package dayFour

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
	cards, err := getCards()
	if err != nil {
		return err
	}

	points := 0

	for _, card := range cards {
		if card == "" {
			continue
		}

		definition := strings.Split(card, ":")
		cardSplit := strings.Split(definition[1], "|")
		cardSplit[0] = strings.TrimLeft(cardSplit[0], " ")
		cardSplit[0] = strings.TrimRight(cardSplit[0], " ")
		cardSplit[1] = strings.TrimLeft(cardSplit[1], " ")
		cardSplit[1] = strings.TrimRight(cardSplit[1], " ")

		strWinningNumbers := strings.Split(cardSplit[0], " ")
		winningNumbers := make([]int, 0)
		for _, num := range strWinningNumbers {
			if num == "" {
				continue
			}

			n, err := strconv.Atoi(num)
			if err != nil {
				return err
			}

			winningNumbers = append(winningNumbers, n)
		}

		strGivenNumbers := strings.Split(cardSplit[1], " ")
		foundNumbers := 0
		for _, num := range strGivenNumbers {
			if num == "" {
				continue
			}

			n, err := strconv.Atoi(num)
			if err != nil {
				return err
			}

			for _, winningNum := range winningNumbers {
				if n == winningNum {
					if foundNumbers == 0 {
						foundNumbers++
					} else {
						foundNumbers *= 2
					}
				}
			}
		}
		points += foundNumbers
	}

	fmt.Println(points)

	return nil
}

func (d service) PartTwo() error {

	return nil
}

func getCards() ([]string, error) {
	input, err := util.GetInput(4)
	if err != nil {
		return nil, err
	}

	cards := strings.Split(input, "\n")
	return cards, err
}
