package utils

import (
	"bufio"
	"os"
	"unicode"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func LineToNumberSlice(line string) []int {
	numbers := make([]int, 0)
	for _, char := range line {
		if unicode.IsDigit(char) {
			numbers = append(numbers, int(char-'0'))
		}
	}
	return numbers
}
