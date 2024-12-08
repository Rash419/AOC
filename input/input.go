package input

import (
	"bufio"
	"os"
)

func ReadFileByLine(filePath string) ([]string, error) {
	lines := make([]string, 0)
	inputFile, err := os.Open(filePath)
	if err != nil {
		return lines, err
	}

	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}
