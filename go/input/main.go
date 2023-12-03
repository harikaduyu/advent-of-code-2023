package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func getWithAOCCookie(url string, cookie string) []byte {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Fatalf("making request: %s", err)
	}

	sessionCookie := http.Cookie{
		Name:  "session",
		Value: cookie,
	}
	req.AddCookie(&sessionCookie)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("making request: %s", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("reading response body: %s", err)
	}
	fmt.Println("response length is", len(body))

	return body
}

func writeToFile(filename string, contents []byte) {
	err := os.MkdirAll(filepath.Dir(filename), os.ModePerm)
	if err != nil {
		log.Fatalf("making directory: %s", err)
	}
	err = os.WriteFile(filename, contents, os.FileMode(0644))
	if err != nil {
		log.Fatalf("writing file: %s", err)
	}
}

func getInput(day int, cookie string) {
	fmt.Printf("fetching for day %d", day)

	// make the request
	url := fmt.Sprintf("https://adventofcode.com/2023/day/%d/input", day)
	body := getWithAOCCookie(url, cookie)
	fmt.Println("body", string(body))
	if strings.HasPrefix(string(body), "Puzzle inputs differ by user") {
		panic("'Puzzle inputs differ by user' response")
	}

	_, cur_filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("getting calling function")
	}
	dirname := filepath.Dir(cur_filename)
	// write to file
	filename := filepath.Join(dirname, "../..", fmt.Sprintf("/inputs/day%02d.txt", day))

	writeToFile(filename, body)

	fmt.Println("Wrote to file: ", filename)

	fmt.Println("Done!")
}

func parseFlags() (day int, cookie string) {
	today := time.Now()
	flag.IntVar(&day, "day", today.Day(), "day number to fetch, 1-25")
	// defaults to env variable
	flag.StringVar(&cookie, "cookie", os.Getenv("AOC_SESSION_COOKIE"), "AOC session cookie")
	flag.Parse()

	if day > 25 || day < 1 {
		log.Fatalf("day out of range: %d", day)
	}

	if cookie == "" {
		log.Fatalf("no session cookie set on flag or env var (AOC_SESSION_COOKIE)")
	}

	return day, cookie
}

func main() {
	day, cookie := parseFlags()
	getInput(day, cookie)
}
