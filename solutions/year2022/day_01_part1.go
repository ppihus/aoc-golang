package year2022

import (
	"sort"

	aochelpers "github.com/ppihus/aoc-golang/pkg/aocHelpers"
)

func Day_01_Part1(filePath string) int {
	data := aochelpers.ReadIntsFromFile(filePath)

	var calorieSums []int
	var calorieSum int

	for i, line := range data {
		calorieSum += line
		if line == 0 || i == len(data)-1 {
			calorieSums = append(calorieSums, calorieSum)
			calorieSum = 0
			continue
		}
	}

	sort.Ints(calorieSums)

	return calorieSums[len(calorieSums)-1]
}
