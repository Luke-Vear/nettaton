package subnetquiz

import (
	"regexp"
	"testing"
)

func TestGetRandomIP(t *testing.T) {

	rgx := `\b(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])\.(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])\.(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])\.(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])\b`

	got := RandomIP()

	if match, _ := regexp.MatchString(rgx, got); !match {
		t.Errorf("RandomIP() = %v, does not match regex\n", got)
	}
}

func TestGetRandomNetwork(t *testing.T) {
	var allPossNets []string
	for _, v := range Networks {
		allPossNets = append(allPossNets, v.netmask, v.prefix)
	}

	iterations := 2048
	if isRandStructError(RandomNetwork, allPossNets, iterations) {
		t.Error("getRandomNetwork() not producing all possible random networks.")
	}
}

func TestGetRandomQuestionKind(t *testing.T) {

	qkinds := make([]string, 0)
	for k := range Questions {
		qkinds = append(qkinds, k)
	}

	iterations := 128
	if isRandStructError(RandomQuestionKind, qkinds, iterations) {
		t.Error("getQuestionKind() not producing all possible random question kinds.")
	}
}

func isRandStructError(genFunc func() string, allPermutations []string, iter int) bool {

	generatedOutputs := make([]string, iter)
	for i := 0; i < iter; i++ {
		generatedOutputs[i] = genFunc()
	}

	var totalUnique int
	for _, v1 := range allPermutations {
		for _, v2 := range generatedOutputs {
			if v1 == v2 {
				totalUnique++
				break
			}
		}
	}
	return totalUnique != len(allPermutations)
}
