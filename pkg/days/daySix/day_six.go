package daySix

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
	times, distances, err := getPartOneInput()
	if err != nil {
		return err
	}

	total := 0

	for i, time := range times {
		beatenDistances := 0

		for j := 0; j < time; j++ {
			newDistance := j * (time - j)
			fmt.Printf("Race %v: %v milliseconds - %v distance\n", i, j, newDistance)
			if newDistance > distances[i] {
				beatenDistances++
			}
		}

		fmt.Println(beatenDistances)

		if total == 0 {
			total = beatenDistances
		} else {
			total *= beatenDistances
		}
	}

	fmt.Println(times, distances)
	fmt.Println(total)
	return nil
}

func (d service) PartTwo() error {
	time, distance, err := getPartTwoInput()
	if err != nil {
		return err
	}

	earliest := 10000000000000
	latest := 0
	for i := 0; i < time; i++ {
		newDistance := i * (time - i)
		if newDistance > distance {
			if i < earliest {
				earliest = i
				break
			}
		}
	}

	for i := time; i > 0; i-- {
		newDistance := i * (time - i)
		if newDistance > distance {
			if i > latest {
				latest = i
				break
			}
		}
	}

	fmt.Println(earliest, latest)
	fmt.Println(latest - earliest + 1)

	return nil
}

func getPartOneInput() ([]int, []int, error) {
	input, err := util.GetInput(6)
	if err != nil {
		return nil, nil, err
	}

	lines := strings.Split(input, "\n")

	times := make([]int, 0)
	distances := make([]int, 0)

	for i, line := range lines {
		if line == "" {
			continue
		}

		split := strings.Split(line, ":")
		vals := strings.Split(split[1], " ")

		for _, val := range vals {
			if val == "" {
				continue
			}

			val := strings.ReplaceAll(val, " ", "")
			intVal, err := strconv.Atoi(val)
			if err != nil {
				return nil, nil, err
			}

			if i == 0 {
				times = append(times, intVal)
			} else {
				distances = append(distances, intVal)
			}
		}
	}

	return times, distances, nil
}

func getPartTwoInput() (int, int, error) {
	input, err := util.GetInput(6)
	if err != nil {
		return 0, 0, err
	}

	lines := strings.Split(input, "\n")

	timeSplit := strings.Split(lines[0], ":")
	timeStr := strings.ReplaceAll(timeSplit[1], " ", "")
	time, err := strconv.Atoi(timeStr)
	if err != nil {
		return 0, 0, err
	}

	distanceSplit := strings.Split(lines[1], ":")
	distanceStr := strings.ReplaceAll(distanceSplit[1], " ", "")
	distance, err := strconv.Atoi(distanceStr)
	if err != nil {
		return 0, 0, err
	}

	return time, distance, nil
}
