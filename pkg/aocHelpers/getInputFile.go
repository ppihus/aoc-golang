package aochelpers

import (
	"log"
	"os"
)

func GetInputFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
