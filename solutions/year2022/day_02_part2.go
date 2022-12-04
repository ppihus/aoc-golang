package year2022

import (
	"bufio"
	"log"
	"regexp"
	"strconv"
	"strings"

	aochelpers "github.com/ppihus/aoc-golang/pkg/aocHelpers"
)

func Day_02_Part2(filePath string) int {
	file := aochelpers.GetInputFile(filePath)
	defer file.Close()
	input := bufio.NewScanner(file)

	score := 0
	for input.Scan() {
		line := input.Text()
		re := regexp.MustCompile(`A|X`)
		line = re.ReplaceAllLiteralString(line, "1")
		re = regexp.MustCompile(`B|Y`)
		line = re.ReplaceAllLiteralString(line, "2")
		re = regexp.MustCompile(`C|Z`)
		line = re.ReplaceAllLiteralString(line, "3")

		numbers := strings.Split(line, " ")
		opponent, _ := strconv.Atoi(numbers[0])
		result, _ := strconv.Atoi(numbers[1])

		if result == 2 {
			score += 3
		} else if result == 3 {
			score += 6
		}

		me := 0
		if (result == 1 && opponent == 3) || (result == 2 && opponent == 2) || (result == 3 && opponent == 1) {
			me = 2 // I'm paper
		} else if (result == 1 && opponent == 2) || (result == 2 && opponent == 1) || (result == 3 && opponent == 3) {
			me = 1 // I'm rock
		} else if (result == 1 && opponent == 1) || (result == 2 && opponent == 3) || (result == 3 && opponent == 2) {
			me = 3 // I'm scissors
		}
		score += me
	}

	if err := input.Err(); err != nil {
		log.Fatal(err)
	}

	return score
}
