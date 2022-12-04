package aochelpers

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// Only download when the file does not exist
// Otherwise we would hammer adventofcode servers
func checkIfFileExists(year string, day string, filePath string) (bool, error) {
	if _, err := os.Stat(filePath); err == nil {
		return true, nil
	} else if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else {
		return false, err //Schrodinger's file. Might exist and also might not.
	}
}

func DownloadPuzzleInput(year string, day string, filePath string) {
	fileExists, _ := checkIfFileExists(year, day, filePath)
	if fileExists {
		return
	}
	fmt.Printf("ℹ️ %s/%s/data.txt does not exist. Going to download\n", year, day)

	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// Get the data
	day = strings.TrimLeft(day, "0") // Remove leading zeros
	var url string = fmt.Sprintf("https://adventofcode.com/%s/day/%s/input", year, day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	var sessionCookie string = os.Getenv("aoc_session_cookie")
	if sessionCookie == "" {
		panic("aoc_session_cookie not set")
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionCookie,
	})

	if os.Getenv("aoc_user_agent") != "" {
		req.Header.Set("User-Agent", os.Getenv("aoc_user_agent"))
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		panic("bad status: " + resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
}
