package problems

type ProblemOne struct {
}

func (p ProblemOne) Solve() int {
	return 23
}

func NewProblemOne(ceiling int) Solver {
	return &ProblemOne{}
}
