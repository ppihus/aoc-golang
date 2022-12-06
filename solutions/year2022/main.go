package year2022

import (
	"fmt"
	"path"
	"runtime"

	aochelpers "github.com/ppihus/aoc-golang/pkg/aocHelpers"
)

const year = "2022"

func Solve(day string) {
	filePath := constructDataPath(day, false)
	aochelpers.DownloadPuzzleInput(year, day, filePath)

	switch day {
	case "01":
		fmt.Println("🎅🏻🎅🏼🎅🏽 Day 1: Calorie Counting")
		printSolution(year, day, "1", Day_01_Part1(filePath))
		printSolution(year, day, "2", Day_01_Part2(filePath))
	case "02":
		fmt.Println("🎅🏻🎅🏼🎅🏽 Day 2: Rock Paper Scissors")
		printSolution(year, day, "1", Day_02_Part1(filePath))
		printSolution(year, day, "2", Day_02_Part1(filePath))
	case "03":
		fmt.Println("🎅🏻🎅🏼🎅🏽 Day 3: Rucksack Reorganization")
		printSolution(year, day, "1", Day_03_Part1(filePath))
		printSolution(year, day, "2", Day_03_Part2(filePath))
	case "04":
		fmt.Println("🎅🏻🎅🏼🎅🏽 Day 3: Rucksack Reorganization")
		printSolution(year, day, "1", Day_04_Part1(filePath))
		printSolution(year, day, "2", Day_04_Part2(filePath))
	default:
		fmt.Println("No day specificed") // Automate to solve latest available
	}
}

func constructDataPath(day string, test bool) string {
	testPath := ""
	if test {
		testPath = "/tests"
	}
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename) + "/data" + testPath + "/" + day + ".txt"
}

func printSolution(year string, day string, part string, solution int) {
	fmt.Printf("🎄🎄🎄 %s/%s/Part %s: %d ⬅️ \n", year, day, part, solution)
}
