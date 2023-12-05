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

	colorConfig := map[string]int{
		"blue":  14,
		"green": 13,
		"red":   12,
	}

	sum := 0

	for i, line := range lines {
		id := i + 1
		game := strings.Split(line, fmt.Sprintf("Game %d:", id))
		hands := strings.Split(game[1], ";")
		pass := true
		for _, hand := range hands {
			cubes := strings.Split(hand, ",")
			for _, cube := range cubes {
				for color, configuration := range colorConfig {
					if strings.Contains(cube, color) && pass {
						pass = checkIfValidCubes(cube, configuration)
					}
				}
			}
		}
		if pass {
			sum = sum + id
		}
		pass = true
	}
	println(sum)
}

func checkIfValidCubes(cubes string, max int) bool {
	numberString := re.FindAllString(cubes, -1)[0]
	number, _ := strconv.Atoi(numberString)
	return number <= max
}