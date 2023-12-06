package dayFive

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
	seeds, err := getSeeds()
	if err != nil {
		return err
	}

	fmt.Println(seeds)

	mappings, err := getMappings()
	if err != nil {
		return err
	}

	locations := calculateMapping(seeds, "seed", mappings)
	fmt.Println(locations)

	lowest := 1000000000000000000
	for _, location := range locations {
		if location < lowest {
			lowest = location
		}
	}

	fmt.Println(lowest)

	return nil
}

func (d service) PartTwo() error {

	return nil
}

func calculateMapping(values []int, currentState string, mappings map[mappingKey][]mappingValue) []int {
	fmt.Println("Current State", currentState)
	currentKey := mappingKey{}
	currentMappings := make([]mappingValue, 0)

	newValues := make([]int, 0)

	for key, values := range mappings {
		if key.source == currentState {
			currentKey = key
			currentMappings = values
		}
	}

	for _, val := range values {
		mappingFound := false
		for _, mapping := range currentMappings {
			if val >= mapping.sourceStart && val <= mapping.sourceStart+mapping.bothRange {
				difference := mapping.destinationStart - mapping.sourceStart
				newValue := val + difference
				newValues = append(newValues, newValue)
				mappingFound = true
				break
			}
		}

		if !mappingFound {
			newValues = append(newValues, val)
		}
	}

	if currentState == "location" {
		return newValues
	} else {
		return calculateMapping(newValues, currentKey.destination, mappings)
	}
}

func getSeeds() ([]int, error) {
	input, err := util.GetInput(5)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(input, "\n")
	seeds := make([]int, 0)
	seedDefinition := lines[0]
	seedSplit := strings.Split(seedDefinition, " ")
	for i := 1; i < len(seedSplit); i++ {
		if seedSplit[i] == "" {
			continue
		}

		n, err := strconv.Atoi(seedSplit[i])
		if err != nil {
			return nil, err
		}

		seeds = append(seeds, n)
	}

	return seeds, err
}

func getMappings() (map[mappingKey][]mappingValue, error) {
	input, err := util.GetInput(5)
	if err != nil {
		return nil, err
	}

	mappings := make(map[mappingKey][]mappingValue)
	lines := strings.Split(input, "\n")
	var currentKey mappingKey
	for i, line := range lines {
		if line == "" || i < 2 {
			continue
		}

		if strings.Contains(line, ":") {
			words := strings.Split(line, " ")
			mapping := strings.Split(words[0], "-to-")

			key := mappingKey{
				source:      mapping[0],
				destination: mapping[1],
			}
			currentKey = key
			mappings[key] = make([]mappingValue, 0)
			continue
		}

		values := strings.Split(line, " ")
		intValues := make([]int, 0)
		for _, value := range values {
			if value == "" {
				continue
			}

			n, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}

			intValues = append(intValues, n)
		}

		mapping := mappingValue{
			destinationStart: intValues[0],
			sourceStart:      intValues[1],
			bothRange:        intValues[2],
		}

		mappings[currentKey] = append(mappings[currentKey], mapping)
	}

	return mappings, err
}

type mappingKey struct {
	source      string
	destination string
}

type mappingValue struct {
	destinationStart int
	sourceStart      int
	bothRange        int
}

type seedCalcValue struct {
	sourceStart  int
	sourceEnd    int
	differential int
}
