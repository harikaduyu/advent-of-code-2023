package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func readFile(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("File reading error %s", err)
	}
	return string(data)

}

func dirName() string {
	_, cur_filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("getting calling function")
	}
	return filepath.Dir(cur_filename)
}

func ReadInput() string {

	day, err := getDayFromFile()
	if err != nil {
		panic(err)
	}
	// write to file
	dirname := dirName()
	filename := filepath.Join(dirname, "..", fmt.Sprintf("/inputs/day%02d/Q_input.txt", day))
	return readFile(filename)
}

func ReadExampleInput(part int) string {
	day, err := getDayFromFile()

	if err != nil {
		panic(err)
	}

	dirname := dirName()
	filename := filepath.Join(dirname, "..", fmt.Sprintf("/inputs/day%02d/example_%d.txt", day, part))
	return readFile(filename)
}

func getDayFromFile() (int, error) {
	// Use runtime.Caller to get information about the calling function
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return 0, fmt.Errorf("Could not retrieve caller information")
	}

	// Use runtime.FuncForPC to get information about the calling function
	callerFunc := runtime.FuncForPC(pc)
	if callerFunc == nil {

		return 0, fmt.Errorf("Could not retrieve caller function")
	}

	// Get the file path of the calling function
	file_path, _ := callerFunc.FileLine(pc)
	// extract day number from file path
	fileSlice := strings.Split(filepath.ToSlash(file_path), "/")
	day_string := fileSlice[len(fileSlice)-2]
	day, err := strconv.Atoi(day_string[len(day_string)-2:])

	if err != nil {
		return 0, fmt.Errorf(err.Error())
	}

	return day, nil
}
