package aochelpers

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
)

func ReadFromFile(filename string) string {
	_, currentFile, _, _ := runtime.Caller(1)
	filePath := path.Dir(currentFile) + "/" + filename

	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func ReadIntsFromFile(filename string) []int {
	_, currentFile, _, _ := runtime.Caller(1)
	filePath := path.Dir(currentFile) + "/" + filename

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)

	var ints []int
	for scanner.Scan() {
		lineStr := scanner.Text()
		num, _ := strconv.Atoi(lineStr)
		ints = append(ints, num)
	}
	return ints

}
