package dayThree

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
	partNumbers, err := getPartNumberLocations()
	if err != nil {
		return err
	}

	rows, err := getSplitInput()
	if err != nil {
		return err
	}

	symbols := []string{"*", "@", "#", "$", "+", "%", "/", "&", "=", "-"}

	foundPartNumbers := make([]partNumber, 0)

	for rowIndex, row := range rows {
		//fmt.Println(row)
		for col, character := range row {
			char := string(character)
			//fmt.Println("current character", char)

			for _, symbol := range symbols {
				if char == symbol {
					fmt.Println("found symbol", symbol)

					coordChecks := getCoordChecks(col, rowIndex)

					for _, check := range coordChecks {
						for _, partNum := range partNumbers {
							for _, coordinate := range partNum.coordinates {
								if check.x == coordinate.x && check.y == coordinate.y {
									fmt.Println("found part number", partNum)
									partExists := false

									for _, foundPartNumber := range foundPartNumbers {
										if foundPartNumber.val == partNum.val && foundPartNumber.instance == partNum.instance {
											partExists = true
											break
										}
									}

									if !partExists {
										foundPartNumbers = append(foundPartNumbers, partNum)
									}
								}
							}
						}
					}
				}
			}
		}
	}

	fmt.Println(foundPartNumbers)
	total := 0
	for _, partNum := range foundPartNumbers {
		total += partNum.val
	}
	fmt.Println("Total:", total)

	return nil
}

func (d service) PartTwo() error {
	partNumbers, err := getPartNumberLocations()
	if err != nil {
		return err
	}

	rows, err := getSplitInput()
	if err != nil {
		return err
	}

	symbols := []string{"*", "@", "#", "$", "+", "%", "/", "&", "=", "-"}

	gearRatios := make([]int, 0)

	for rowIndex, row := range rows {
		//fmt.Println(row)
		for col, character := range row {
			char := string(character)
			//fmt.Println("current character", char)

			for _, symbol := range symbols {
				if char == symbol {
					fmt.Println("found symbol", symbol)

					coordChecks := getCoordChecks(col, rowIndex)
					foundParts := make([]partNumber, 0)

					for _, check := range coordChecks {
						for _, partNum := range partNumbers {
							for _, coordinate := range partNum.coordinates {
								if check.x == coordinate.x && check.y == coordinate.y {
									fmt.Println("found part number", partNum)
									partExists := false

									for _, foundPartNumber := range foundParts {
										if foundPartNumber.val == partNum.val && foundPartNumber.instance == partNum.instance {
											partExists = true
											break
										}
									}

									if !partExists {
										foundParts = append(foundParts, partNum)
									}
								}
							}
						}
					}

					if len(foundParts) == 2 {
						ratio := foundParts[0].val * foundParts[1].val
						gearRatios = append(gearRatios, ratio)
					}
				}
			}
		}
	}

	fmt.Println(gearRatios)
	total := 0
	for _, ratio := range gearRatios {
		total += ratio
	}
	fmt.Println("Total:", total)

	return nil
}

func getPartNumberLocations() ([]partNumber, error) {
	rows, err := getSplitInput()
	if err != nil {
		return nil, err
	}

	partNumbers := make([]partNumber, 0)

	for rowIndex, row := range rows {
		fmt.Println(row)
		row = strings.TrimSpace(row)
		latestCol := -1
		for col, _ := range row {
			if col <= latestCol {
				continue
			}

			numString := ""
			coordinates := make([]coordinate, 0)
			// fmt.Println("col", col)
			for {
				if col <= len(row)-1 {
					character := string(row[col])
					if _, err := strconv.Atoi(character); err == nil {
						// fmt.Printf("%v is a number\n", character)
						numString += character
						coordinates = append(coordinates, coordinate{
							x: col,
							y: rowIndex,
						})
						col++
						continue
					}
				}

				// fmt.Printf("%v is not a number\n", character)
				if numString == "" {
					break
				}

				number, err := strconv.Atoi(numString)
				if err != nil {
					return nil, err
				}

				instance := 0

				for _, pn := range partNumbers {
					if pn.val == number {
						instance++
					}
				}

				partNumber := partNumber{
					val:         number,
					instance:    instance,
					coordinates: coordinates,
				}
				partNumbers = append(partNumbers, partNumber)

				latestCol = col
				break
			}
		}
	}

	fmt.Println(partNumbers)
	return partNumbers, nil
}

func getSplitInput() ([]string, error) {
	input, err := util.GetInput(3)
	if err != nil {
		return nil, err
	}

	rows := strings.Split(input, "\n")
	return rows, nil
}

func getCoordChecks(col int, rowIndex int) []coordinate {
	coordChecks := []coordinate{
		{
			x: col - 1,
			y: rowIndex - 1,
		},
		{
			x: col,
			y: rowIndex - 1,
		},
		{
			x: col + 1,
			y: rowIndex - 1,
		},
		{
			x: col - 1,
			y: rowIndex,
		},
		{
			x: col + 1,
			y: rowIndex,
		},
		{
			x: col - 1,
			y: rowIndex + 1,
		},
		{
			x: col,
			y: rowIndex + 1,
		},
		{
			x: col + 1,
			y: rowIndex + 1,
		},
	}

	filteredCoordChecks := make([]coordinate, 0)
	for _, coordCheck := range coordChecks {
		if coordCheck.x >= 0 || coordCheck.y >= 0 {
			filteredCoordChecks = append(filteredCoordChecks, coordCheck)
		}
	}

	return filteredCoordChecks
}

type partNumber struct {
	val         int
	instance    int
	coordinates []coordinate
}

type coordinate struct {
	x int
	y int
}
