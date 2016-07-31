package problems

type ProblemOne struct {
	ceiling int
}

func (p ProblemOne) Solve() int {
	var answer int
	for i := 0; i < p.ceiling; i++ {
		if i%3 == 0 || i%5 == 0 {
			answer += i
		}
	}
	return answer
}

func NewProblemOne(ceiling int) Solver {
	return &ProblemOne{ceiling: ceiling}
}
