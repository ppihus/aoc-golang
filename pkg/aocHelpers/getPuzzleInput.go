package aochelpers

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"runtime"
)

func DownloadPuzzleInput(year string, day string) error {

	// Find the solution file path
	_, filename, _, _ := runtime.Caller(1)
	filePath := path.Dir(filename) + "/data.txt"

	// Only download when the file does not exist
	// Otherwise we would hammer adventofcode servers
	if _, err := os.Stat(filePath); err == nil {
		fmt.Println("🎅 Input data.txt already exists, skipping downloading")
		return nil
	} else if errors.Is(err, os.ErrNotExist) {
		fmt.Println("🎅 Input data.txt does not exist, downloading...")
		downloadInputData(year, day, filePath)
		return nil
	} else {
		return err
	}
}

func downloadInputData(year string, day string, filePath string) {

	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// Get the data
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
