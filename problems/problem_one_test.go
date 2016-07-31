package problems_test

import (
	"testing"

	"github.com/Tolk-Haggard/project-goeuler/problems"
	"github.com/stretchr/testify/assert"
)

func Test_ProblemOneReturnsExampleAnswer(t *testing.T) {
	testObject := problems.NewProblemOne(10)

	assert.Equal(t, 23, testObject.Solve())

}

func Test_ProblemOneReturnsSumOfMultiplesOf3And5Under20(t *testing.T) {
	// sum(3, 5, 6, 9, 10, 12, 15, 18) = 78
	testObject := problems.NewProblemOne(20)

	assert.Equal(t, 78, testObject.Solve())

}

func Test_ProblemOneReturnsSumOfMultiplesOf3And5Under1000(t *testing.T) {
	testObject := problems.NewProblemOne(1000)

	assert.Equal(t, 233168, testObject.Solve())

}
