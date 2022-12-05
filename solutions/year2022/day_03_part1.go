package year2022

import (
	"bufio"
	"io"
	"log"

	"github.com/juliangruber/go-intersect/v2"
	aochelpers "github.com/ppihus/aoc-golang/pkg/aocHelpers"
)

func Day_03_Part1(filePath string) int {
	file := aochelpers.GetInputFile(filePath)
	defer file.Close()
	r := bufio.NewReader(file)

	itemsInBothHalves := make(map[int]bool)
	var rucksack []int
	var answer int

	for {
		if c, _, err := r.ReadRune(); err == nil {
			if c == '\n' { // end of rucksack

				half := len(rucksack) / 2
				duplicatesInOnerucksack := intersect.SimpleGeneric(rucksack[:half], rucksack[half:])

				for _, duplicate := range duplicatesInOnerucksack {
					itemsInBothHalves[duplicate] = false
				}

				for d := range itemsInBothHalves {
					answer += d
				}

				itemsInBothHalves = make(map[int]bool)
				rucksack = []int{}
			} else {
				rucksack = append(rucksack, getCharacterPosition(c))
			}
		} else {
			if err == io.EOF {
				break
			} else { // unknown error
				log.Fatal(err)
			}
		}
	}
	return answer
}

func createEmptySet() map[string]bool {
	return make(map[string]bool)
}

func getCharacterPosition(c rune) int {
	subtract := 38
	if int(c) > 92 {
		subtract = 96
	}
	return int(c) - subtract
}
