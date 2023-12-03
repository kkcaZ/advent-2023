package days

import (
	"github.com/kkcaZ/advent-2024/pkg/days/dayThree"
	"github.com/kkcaZ/advent-2024/pkg/days/dayTwo"
	"github.com/kkcaZ/advent-2024/pkg/domain"
	"github.com/pkg/errors"
)

func NewDayService(day int) (domain.DayService, error) {
	dayServices := make(map[int]domain.DayService)
	dayServices[2] = dayTwo.New()
	dayServices[3] = dayThree.New()

	if dayServices[day] == nil {
		return nil, errors.New("day not found")
	}

	return dayServices[day], nil
}
