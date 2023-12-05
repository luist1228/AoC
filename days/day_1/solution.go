package day1

import (
	"fmt"
	"log"
	"strings"
	"unicode"

	"github.com/luist1228/AoC/utils"
)

func part1() {
	day := "day_1"
	fileName := "/input_1"
	path := "./days/" + day + fileName + ".txt"

	lines, err := utils.ReadLines(path)

	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	sum := 0

	for _, line := range lines {
		numbers := utils.LineToNumberSlice(line)
		first := numbers[0]
		last := numbers[len(numbers)-1]
		result := (first * 10) + last
		sum = sum + result
	}
	fmt.Println(len(lines))
	fmt.Println(sum)
}

func checkStringContainsDigit(s string) (int, bool) {
	digits := []string{"cero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for i, digit := range digits {
		if strings.Contains(s, digit) {
			return i, true
		}
	}
	return 0, false

}

func part2() {
	day := "day_1"
	fileName := "/input_1"
	path := "./days/" + day + fileName + ".txt"

	lines, err := utils.ReadLines(path)

	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	sum := 0
	for _, line := range lines {
		numbers := make([]int, 0)
		auxString := ""
		for _, char := range line {
			if unicode.IsDigit(char) {
				numbers = append(numbers, int(char-'0'))
				auxString = ""
			}
			auxString = auxString + string(char)
			digit, found := checkStringContainsDigit(auxString)
			if found {
				numbers = append(numbers, digit)
				auxString = auxString[len(auxString)-2:]
			}
		}
		first := numbers[0]
		last := numbers[len(numbers)-1]
		result := (first * 10) + last
		sum = sum + result
	}
	fmt.Println(len(lines))
	fmt.Println(sum)
}
