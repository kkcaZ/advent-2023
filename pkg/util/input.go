package util

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

func GetInput(day int) (string, error) {
	inputExists, err := configFileExists(fmt.Sprintf("day_%v_input", day))
	if err != nil {
		return "", errors.Wrap(err, "failed to check if input exists")
	}

	if inputExists {
		return getFileFromConfig(fmt.Sprintf("day_%v_input", day))
	}

	input, err := getInputFromAdvent(day)
	if err != nil {
		return "", errors.Wrap(err, "failed to get input from advent")
	}

	err = addFileToConfig(fmt.Sprintf("day_%v_input", day), input)
	if err != nil {
		return "", errors.Wrap(err, "failed to store input")
	}

	return input, nil
}

func getInputFromAdvent(day int) (string, error) {
	client := http.Client{}
	session, err := GetToken()
	if err != nil {
		return "", errors.Wrap(err, "failed to get session token")
	}

	url := fmt.Sprintf("https://adventofcode.com/2023/day/%v/input", day)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Cookie", fmt.Sprintf("session=%s", session))

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("non 200 status code returned: %v", resp.StatusCode)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
