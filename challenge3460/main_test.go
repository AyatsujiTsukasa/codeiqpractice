package main

import (
	"math/big"
	"testing"
)

func TestSolve(t *testing.T) {
	testSet := []string{"2 1", "3 2"}
	ansSet := []string{"16", "72"}
	for i := 0; i < len(testSet); i++ {
		ans := solve(testSet[i])
		if ans != ansSet[i] {
			t.Errorf("got %v\n want %v", ans, ansSet[i])
		}
	}
}

var (
	expected = big.NewInt(137846528820)
	forty    = big.NewInt(40)
	twenty   = big.NewInt(20)
)

func TestConbination(t *testing.T) {
	ans := combination(forty, twenty)
	if ans.Cmp(expected) != 0 {
		t.Errorf("got %v\n want %v", ans, 137846528820)
	}
}
