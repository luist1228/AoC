package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/luist1228/AoC/utils"
)

var re = regexp.MustCompile("[0-9]+")

func main() {
	day := "day_2"
	fileName := "/input_1"
	path := "./days/" + day + fileName + ".txt"

	lines, err := utils.ReadLines(path)

	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	sum := 0

	for i, line := range lines {
		id := i + 1
		game := strings.Split(line, fmt.Sprintf("Game %d:", id))
		hands := strings.Split(game[1], ";")

		colorsMin := map[string]int{
			"blue":  0,
			"green": 0,
			"red":   0,
		}
		for _, hand := range hands {
			cubes := strings.Split(hand, ",")
			for _, cube := range cubes {
				for color, min := range colorsMin {
					if strings.Contains(cube, color) {
						colorsMin[color] = checkMinCubes(cube, min)
					}
				}
			}
		}
		mult := 1
		for _, min := range colorsMin {
			mult = mult * min
		}
		if mult > 0 {
			sum = sum + mult
		}

	}
	println(sum)

}

func checkMinCubes(cubes string, max int) int {
	numberString := re.FindAllString(cubes, -1)[0]
	number, _ := strconv.Atoi(numberString)

	if number > max {
		return number
	}

	return max
}
