package main

import "testing"

func TestGetFactors(t *testing.T) {
	verifyExpectancy(1, []uint{}, t)
	verifyExpectancy(2, []uint{1}, t)
	verifyExpectancy(64, []uint{1, 2, 4, 8, 16, 32}, t)
}

func verifyExpectancy(argument uint, expected []uint, t *testing.T) {
	got := GetFactorsExcluding(argument)
	correct := true
	if len(got) == len(expected) {
		for index, value := range expected {
			correct = correct && (value == got[index])
		}
	} else {
		correct = false
	}

	if !correct {
		t.Errorf("GetFactors(%d) = %v; want %v", argument, got, expected)
	}
}
