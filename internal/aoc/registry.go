package aoc

type Solution interface {
	Day() int
	Part1(input string) (any, error)
	Part2(input string) (any, error)
}
