package day01

import (
	"aoc25/internal/aoc"
	"log"
	"strconv"
	"strings"
)

type sol struct{}

func (sol) Day() int { return 1 }

func (sol) Part1(input string) (any, error) {
	chars := parseInput(input)
	var count = 0
	var position = 50
	for _, ch1 := range chars {
		direction := string(ch1[0])
		numStr := ch1[1:]
		n, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatal("Invalid number:", err)
		}
		calculatePosition(direction, n, &position)
		if position == 0 {
			count++
		}
	}

	return count, nil
}

func (sol) Part2(input string) (any, error) {
	// TODO: parse + solve
	return nil, nil
}

func init() { aoc.Register(sol{}) }

func parseInput(input string) []string {
	var chars []string
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		chars = append(chars, line)
	}
	return chars
}

func calculatePosition(direction string, move int, position *int) {
	if move < 0 {
		log.Fatal("Invalid move:", move)
	}
	move %= 100
	switch direction {
	case "L":
		*position = (*position - move + 100) % 100
	case "R":
		*position = (*position + move) % 100
	default:
		log.Fatal("Invalid direction:", direction)
	}
}
