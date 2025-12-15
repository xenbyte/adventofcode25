package day01

import (
	"aoc25/internal/aoc"
	"aoc25/internal/util"
	"fmt"
	"log"
	"strconv"
)

type sol struct{}

func (sol) Day() int { return 1 }

func (sol) Part1(input string) (any, error) {
	chars := util.ParseInput(input)
	var count = 0
	var position = 50
	for _, ch1 := range chars {
		dir := string(ch1[0])
		numStr := ch1[1:]
		n, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatal("Invalid number:", err)
		}
		calculatePosition(dir, n, &position)
		if position == 0 {
			count++
		}
	}

	return count, nil
}

func (sol) Part2(input string) (any, error) {
	lines := util.ParseInput(input)
	pos := 50
	count := 0
	for _, line := range lines {
		dir := line[:1]
		n, err := strconv.Atoi(line[1:])
		if err != nil {
			return nil, fmt.Errorf("invalid number in line %q: %v", line, err)
		}
		count += clicksToZero(dir, n, &pos)
	}
	return count, nil
}

func init() { aoc.Register(sol{}) }

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

func clicksToZero(dir string, move int, pos *int) int {
	p := *pos % 100
	if p < 0 || move < 0 {
		log.Fatalf("Invalid pos/move: p=%d move=%d ", p, move)
	}

	switch dir {
	case "R":
		first := (100 - p) % 100
		if first == 0 {
			first = 100
		}
		hits := 0
		if move >= first {
			hits = 1 + (move-first)/100
		}
		*pos = (p + move) % 100
		return hits

	case "L":
		first := p
		if first == 0 {
			first = 100
		}
		hits := 0
		if move >= first {
			hits = 1 + (move-first)/100
		}
		*pos = (p - (move % 100) + 100) % 100
		return hits
	default:
		log.Fatalf("Invalid dir %q", dir)
		return 0
	}
}
