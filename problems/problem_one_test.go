package problems_test

import (
	"testing"

	"github.com/Tolk-Haggard/project-goeuler/problems"
	"github.com/stretchr/testify/assert"
)

func Test_ProblemOneReturnsExampleAnswer(t *testing.T) {
	ceiling := 10
	factors := []int{3, 5}
	testObject := problems.NewProblemOne(ceiling, factors[0], factors[1])

	assert.Equal(t, 23, testObject.Solve())

}

func Test_ProblemOneReturnsSumOfMultiplesOf3And5Under20(t *testing.T) {
	// sum(3, 5, 6, 9, 10, 12, 15, 18) = 78
	ceiling := 20
	factors := []int{3, 5}
	testObject := problems.NewProblemOne(ceiling, factors[0], factors[1])

	assert.Equal(t, 78, testObject.Solve())

}

func Test_ProblemOneReturnsSumOfMultiplesOf3And5Under1000(t *testing.T) {
	ceiling := 1000
	factors := []int{3, 5}
	testObject := problems.NewProblemOne(ceiling, factors[0], factors[1])

	assert.Equal(t, 233168, testObject.Solve())

}

func Test_ProblemOneReturnsSumOfMultiplesOfNFactors(t *testing.T) {
	ceiling := 10
	factors := []int{3, 5, 7}
	testObject := problems.NewProblemOne(ceiling, factors[0], factors[1], factors[2])

	assert.Equal(t, 30, testObject.Solve())

}

func Test_ProblemOneReturnsSumOfMultiplesOfNFactorsUnder22(t *testing.T) {
	// sum(3, 5, 6, 7, 9, 10, 12, 14, 15, 18, 19, 20, 21) = 159
	ceiling := 22
	factors := []int{3, 5, 7, 19}
	testObject := problems.NewProblemOne(ceiling, factors[0], factors[1], factors[2], factors[3])

	assert.Equal(t, 159, testObject.Solve())

}
