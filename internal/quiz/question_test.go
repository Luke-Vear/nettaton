package quiz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQuestion(t *testing.T) {
	nq := NewQuestion("", "", "")
	assert.NotNil(t, nq)
}

func TestRandomIP(t *testing.T) {
	iter := 100
	seen := make(map[string]bool)

	for i := 0; i < iter; i++ {
		rip, _ := randomIP()
		o1 := rip[:3]
		seen[o1] = true
		if len(seen) == 3 {
			return
		}
	}
	t.Errorf("after %v iterations only the following were found: %v", iter, seen)

}
func TestRandomNetwork(t *testing.T) {
	iter := len(networks) * len(networks)
	seen := make(map[string]bool)

	for i := 0; i < iter; i++ {
		seen[randomNetwork(12)] = true
		if len(seen) == len(networks) {
			return
		}
	}
	t.Errorf("after %v iterations only the following were found: %v", iter, seen)
}
func TestRandomQuestionKind(t *testing.T) {
	iter := len(solvers) * len(solvers)
	seen := make(map[string]bool)

	for i := 0; i < iter; i++ {
		seen[randomQuestionKind()] = true
		if len(seen) == len(solvers) {
			return
		}
	}
	t.Errorf("after %v iterations only the following were found: %v", iter, seen)
}
