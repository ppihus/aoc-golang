package year2022

import (
	"bufio"
	"io"
	"log"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/juliangruber/go-intersect/v2"
	aochelpers "github.com/ppihus/aoc-golang/pkg/aocHelpers"
)

func Day_03_Part2(filePath string) int {
	file := aochelpers.GetInputFile(filePath)
	defer file.Close()
	r := bufio.NewReader(file)

	var answer int
	var group []string
	var groupsOfThree [][]string
	duplicatesSet := mapset.NewSet[string]()

	i := 0
	for {
		if c, _, err := r.ReadRune(); err == nil {
			if c == '\n' { // end of rucksack
				groupsOfThree = append(groupsOfThree, group)
				if (i+1)%3 == 0 { // Third rucksack of the group
					duplicates := getIntersectionOfThreeSlices(groupsOfThree[0], groupsOfThree[1], groupsOfThree[2])

					for _, duplicate := range duplicates {
						duplicatesSet.Add(duplicate)
					}

					for duplicate := range duplicatesSet.Iterator().C {
						answer += getCharacterPosition([]rune(duplicate)[0])
					}

					duplicatesSet = mapset.NewSet[string]() // reset each time

					groupsOfThree = nil
				}
				group = nil
				i++
			} else { // rucksack character
				group = append(group, string(c))
			}
		} else {
			if err == io.EOF { //end of file
				break
			} else { // unknown error
				log.Fatal(err)
			}
		}
	}
	return answer
}

func getIntersectionOfThreeSlices(first []string, second []string, third []string) []string {
	duplicates := intersect.SimpleGeneric(
		intersect.SimpleGeneric(
			first, second),
		third)
	return duplicates
}
