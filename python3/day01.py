import os
from typing import List
from utils import readlines


def part1(lines: List[str]) -> int:
    result = 0
    for line in lines:
        numbers = []
        for char in line:
            if char.isnumeric():
                numbers.append(int(char))
        if not len(numbers):
            raise ValueError("The line has no numeric characters")
        result += (numbers[0] * 10) + numbers[-1]
    return result


WORDS_TO_DIGITS = {
    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9,
}


def part2(lines: List[str]) -> int:
    result = 0
    for line in lines:
        numbers = []
        for i, char in enumerate(line):
            if char.isnumeric():
                number = int(char)
                numbers.append(number)
                continue
            # check with a window of length 3,4 and 5
            # to find possible digit words which can only be
            # of length 3,4 or 5
            for j in 3, 4, 5:
                possible_word = line[i : i + j]
                if possible_word in WORDS_TO_DIGITS:
                    number = WORDS_TO_DIGITS[possible_word]
                    numbers.append(number)

        if not len(numbers):
            raise ValueError("The line has no numeric characters")
        nums = (numbers[0] * 10) + numbers[-1]
        result += nums
    return result


if __name__ == "__main__":
    test_lines = ["1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"]
    test = part1(test_lines)
    print("test: ", test)

    day_string = __file__.split("/")[-1].split(".py")[0]
    lines = readlines(day_string)
    result = part1(lines)
    print("result: ", result)

    part2_test_lines = [
        "two1nine",
        "eightwothree",
        "abcone2threexyz",
        "xtwone3four",
        "4nineeightseven2",
        "zoneight234",
        "7pqrstsixteen",
    ]

    test = part2(part2_test_lines)
    print("part 2 test: ", test)
    lines = readlines(day_string)
    result_2 = part2(lines)
    print("result_2: ", result_2)
