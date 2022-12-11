package utils

import (
	"bufio"
	"os"
)

func LoadNamesFile(path string) []string {
	names := []string{}

	file, err := os.Open(path)
	if err != nil {
		panic(err) // TODO: Handle error
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		names = append(names, scanner.Text())
	}

	return names
}
