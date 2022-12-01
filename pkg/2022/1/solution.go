package main

import (
	"fmt"
	"sort"

	aochelpers "github.com/ppihus/golangtest/pkg/aocHelpers"
)

func main() {
	err := aochelpers.DownloadPuzzleInput("2022", "1")
	if err != nil {
		panic(err) // Schrodinger's data.txt. Might exist and also might not.
	}
	data := aochelpers.ReadIntsFromFile("data.txt")
	// data := []int{1000, 2000, 3000, 0, 4000, 0, 5000, 6000, 0, 7000, 8000, 9000, 0, 10000}

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

	fmt.Println("1st part: ", calorieSums[len(calorieSums)-1])
	fmt.Println("2nd part: ",
		calorieSums[len(calorieSums)-1]+
			calorieSums[len(calorieSums)-2]+
			calorieSums[len(calorieSums)-3])

}
