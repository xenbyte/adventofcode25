package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"aoc25/internal/aoc"
	_ "aoc25/solutions/day01"
)

func main() {
	day := flag.Int("day", 1, "day number")
	part := flag.Int("part", 1, "part number (1 or 2)")
	file := flag.String("file", "", "input file (defaults to solutions/dayXX/input.txt)")
	example := flag.Bool("example", false, "use example.txt")
	flag.Parse()

	defaultFile := fmt.Sprintf("solutions/day%02d/%s", *day, ternary(*example, "example.txt", "input.txt"))
	path := *file
	if path == "" {
		path = defaultFile
	}

	input, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("read %s: %v", path, err)
	}

	sol, ok := aoc.Get(*day)
	if !ok {
		log.Fatalf("no solution for day %d", *day)
	}

	var out any
	switch *part {
	case 1:
		out, err = sol.Part1(string(input))
	case 2:
		out, err = sol.Part2(string(input))
	default:
		log.Fatalf("part must be 1 or 2")
	}
	if err != nil {
		log.Fatalf("solve: %v", err)
	}
	fmt.Println(out)
}

func ternary[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}
