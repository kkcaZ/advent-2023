package util

import (
	"github.com/mitchellh/go-homedir"
	"io"
	"os"
	"path"
)

func StoreToken(token string) error {
	return addFileToConfig("token", token)
}

func GetToken() (string, error) {
	return getFileFromConfig("token")
}

func addFileToConfig(fileName string, content string) error {
	configDirectory, err := getOrCreateConfigDirectory()
	if err != nil {
		return err
	}

	file, err := os.Create(path.Join(configDirectory, fileName))
	if err != nil {
		return err
	}

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

func getFileFromConfig(fileName string) (string, error) {
	configDirectory, err := getOrCreateConfigDirectory()
	if err != nil {
		return "", err
	}

	file, err := os.Open(path.Join(configDirectory, fileName))
	if err != nil {
		return "", err
	}

	b, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func configFileExists(fileName string) (bool, error) {
	configDirectory, err := getOrCreateConfigDirectory()
	if err != nil {
		return false, err
	}

	_, err = os.Stat(path.Join(configDirectory, fileName))
	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func getOrCreateConfigDirectory() (string, error) {
	homeDir, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	configDirectory := path.Join(homeDir, ".advent-2024")
	_, err = os.Stat(configDirectory)
	if os.IsNotExist(err) {
		err := os.Mkdir(configDirectory, os.FileMode(0700))
		if err != nil {
			return "", err
		}
	}
	return configDirectory, nil
}
