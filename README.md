# Advent of code 2023

I started learning Go and AOC seemed like a good opportunity to exercise. 

Got some inspiration for the folder structure and download input scripts from [alexchao](https://github.com/alexchao26/advent-of-code-go).

Either set your `AOC_SESSION_COOKIE` env variable or add it into the make command as an argument like so:
```bash
make input DAY=5 AOC_SESSION_COOKIE=your_cookie
```
which will download your puzzle input into `./inputs/day05/Q_input.txt`

To run all the tests:

```bash
go test ./...
```