package year2022

import (
	"bufio"
	"io"
	"log"
	"regexp"

	aochelpers "github.com/ppihus/aoc-golang/pkg/aocHelpers"
)

func Day_04_Part2(filePath string) int {

	file := aochelpers.GetInputFile(filePath)
	defer file.Close()
	r := bufio.NewReader(file)

	var answer int

	for {
		if line, err := r.ReadString('\n'); err == nil {
			r := regexp.MustCompile(`(?P<first_start>[0-9]+)-(?P<first_end>[0-9]+),(?P<second_start>[0-9]+)-(?P<second_end>[0-9]+)`)
			pairs := r.FindStringSubmatch(line)[1:]
			if (pInt(pairs[2]) <= pInt(pairs[0]) && pInt(pairs[0]) <= pInt(pairs[3]) || pInt(pairs[2]) <= pInt(pairs[1]) && pInt(pairs[1]) <= pInt(pairs[3])) ||
				(pInt(pairs[0]) <= pInt(pairs[2]) && pInt(pairs[2]) <= pInt(pairs[1]) || pInt(pairs[0]) <= pInt(pairs[3]) && pInt(pairs[3]) <= pInt(pairs[1])) {
				answer += 1
			}
		} else {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}
	}
	return answer
}
