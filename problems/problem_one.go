package problems

type ProblemOne struct {
	ceiling int
	factors []int
}

func (p ProblemOne) Solve() int {
	var answer int
	for i := 0; i < p.ceiling; i++ {
		if i%p.factors[0] == 0 || i%p.factors[1] == 0 {
			answer += i
		}
	}
	return answer
}

func NewProblemOne(ceiling int, factors ...int) Solver {
	return &ProblemOne{ceiling: ceiling, factors: factors}
}
