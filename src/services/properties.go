package services

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

type Properties map[string]string

func ReadPropertiesFromPath(path string) (Properties, error) {

	properties := Properties{}

	if len(path) == 0 {
		return properties, errors.New("Path not found")
	}
	file, err := os.Open(path)
	if err != nil {
		return properties, err
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		fileLine := fileScanner.Text()
		if len(fileLine) > 0 {
			propertySplit := strings.Split(fileLine, "=")
			if len(propertySplit) == 2 && len(propertySplit[0]) > 0 && len(propertySplit[1]) > 0 {
				properties[propertySplit[0]] = propertySplit[1]
			}
		}
	}

	if err := fileScanner.Err(); err != nil {
		return properties, err
	}
	return properties, nil
}
