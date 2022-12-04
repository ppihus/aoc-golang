package year2022

import (
	"bufio"
	"log"
	"regexp"
	"strconv"
	"strings"

	aochelpers "github.com/ppihus/aoc-golang/pkg/aocHelpers"
)

func Day_02_Part1(filePath string) int {
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
		me, _ := strconv.Atoi(numbers[1])

		score += me
		if (me == 1 && opponent == 3) || (me == 2 && opponent == 1) || (me == 3 && opponent == 2) {
			score += 6
		} else if me == opponent {
			score += 3
		} else {
			score += 0
		}
	}

	if err := input.Err(); err != nil {
		log.Fatal(err)
	}

	return score
}
