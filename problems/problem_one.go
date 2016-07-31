package problems

type ProblemOne struct {
	ceiling int
	factors []int
}

func (p ProblemOne) Solve() int {
	var answer int
	for i := 0; i < p.ceiling; i++ {
		for _, factor := range p.factors {
			if i%factor == 0 {
				answer += i
				break
			}
		}
	}
	return answer
}

func NewProblemOne(ceiling int, factors ...int) Solver {
	return &ProblemOne{ceiling: ceiling, factors: factors}
}
