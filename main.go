package main

import (
	"fmt"
	"log"

	"github.com/luist1228/AoC/utils"
)

func main() {
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
