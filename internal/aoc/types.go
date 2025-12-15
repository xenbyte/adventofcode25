package aoc

var register = make(map[int]Solution)

func Register(s Solution) {
	register[s.Day()] = s
}

func Get(day int) (Solution, bool) {
	sol, ok := register[day]
	return sol, ok
}
