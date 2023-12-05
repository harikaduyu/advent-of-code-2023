package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/harikaduyu/advent-of-code-2023/go/utils"
)

var wordToDigit = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"zero":  "0",
}

func extractDigitsFromWords(line string) string {

	for _, word := range utils.GetKeys(wordToDigit) {
		// 😈 replace the word digit with number but keep 1st and last character
		// because sometimes they overlap, like "eightwo"
		replaceString := string([]rune(word)[0]) + wordToDigit[word] + word[len(word)-1:]
		line = strings.Replace(line, word, replaceString, -1)
	}
	return line

}

func findFirstAndLastDigits(line string) (firstDigit, lastDigit int, found bool) {
	found = false

	for _, char := range line {
		if unicode.IsDigit(char) {
			if !found {
				firstDigit = int(char - '0')
				found = true
			}
			lastDigit = int(char - '0')
		}
	}

	return firstDigit, lastDigit, found
}

func result(input string, part int) int {
	lines := strings.Fields(input)
	total := 0
	for _, line := range lines {
		if part == 2 {
			line = extractDigitsFromWords(line)
		}
		first, last, found := findFirstAndLastDigits(line)
		if !found {
			txt := fmt.Sprintf("No digits in line:%s", line)
			panic(txt)
		}
		total += (first * 10) + last
	}
	return total
}

func main() {
	input := utils.ReadInput(1)
	result_1 := result(input, 1)
	fmt.Println("Part1:", result_1)
	result_2 := result(input, 2)
	fmt.Println("Part2:", result_2)
}
