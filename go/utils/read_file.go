package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func readFile(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("File reading error %s", err)
	}
	// fmt.Println("Contents of file:")
	// fmt.Println(string(data))
	return string(data)

}

func dirName() string {
	_, cur_filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("getting calling function")
	}
	return filepath.Dir(cur_filename)
}

func ReadInput(day int) string {
	// write to file
	dirname := dirName()
	filename := filepath.Join(dirname, "../..", fmt.Sprintf("/inputs/day%02d.txt", day))
	return readFile(filename)
}

func ReadExampleInput(day int, example int) string {

	dirname := dirName()
	filename := filepath.Join(dirname, "..", fmt.Sprintf("/day%02d/example_%d.txt", day, example))
	return readFile(filename)
}
