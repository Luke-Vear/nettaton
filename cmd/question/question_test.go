package main

import (
	"strings"
	"testing"
)

// TODO: make this a regex
func TestGetRandomIP(t *testing.T) {
	got := getRandomIP()
	if strings.Count(got, ".") != 3 {
		t.Errorf("getRandomIP() = %v, this should have three dots", got)
	}
}

func TestGetRandomNetwork(t *testing.T) {
	var allPossNets []string
	for _, v := range networks {
		allPossNets = append(allPossNets, v.netmask, v.prefix)
	}

	iterations := 2048
	if isRandStructError(getRandomNetwork, allPossNets, iterations) {
		t.Error("getRandomNetwork() not producing all possible random question kinds.")
	}
}

func TestGetRandomQuestionKind(t *testing.T) {
	iterations := 128
	if isRandStructError(getRandomQuestionKind, qkinds, iterations) {
		t.Error("getQuestionKind() not producing all possible random question kinds.")
	}
}

// TODO: test Handle
//func TestHandle(t *testing.T) {}

// utility function
func isRandStructError(genFunc func() string, expected []string, iter int) bool {
	var ranGens []string
	for i := 0; i < iter; i++ {
		ranGens = append(ranGens, genFunc())
	}

	correct := make(chan bool, len(expected))
	for _, v1 := range expected {
		for _, v2 := range ranGens {
			if v1 == v2 {
				correct <- true
				break
			}
		}
	}
	return len(correct) != len(expected)
}
