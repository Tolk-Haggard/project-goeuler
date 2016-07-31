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
