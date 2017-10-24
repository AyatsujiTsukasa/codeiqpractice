package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	testSet := []string{"307", "456", "545", "165", "497"}
	ansSet := []string{"3257", "0", "1835", "6061", "2012"}
	for i := 0; i < len(testSet); i++ {
		ans := solve(testSet[i])
		if ans != ansSet[i] {
			t.Errorf("got %v\n want %v", ans, ansSet[i])
		}
	}
}
