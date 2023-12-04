import os
from typing import List
from utils import read_input_lines


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
    example_1 = part1(read_input_lines(example=1))
    assert example_1 == 142

    lines = read_input_lines()
    result = part1(lines)
    print("Part 1: ", result)

    example_2 = part2(read_input_lines(example=2))
    assert example_2 == 281

    lines = read_input_lines()
    result = part2(lines)
    print("Part 2: ", result)
