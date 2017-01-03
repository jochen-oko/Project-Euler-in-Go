package main

import "testing"

func TestIsDividedByAny(t *testing.T) {

	if !isDividedByAny(12, []int{3, 9}) {
		t.Error("12 must be dividable by 3")
	}

	if !isDividedByAny(12, []int{4}) {
		t.Error("12 must be dividable by 4")
	}

	if isDividedByAny(11, []int{4}) {
		t.Error("11 must not be dividable by 4")
	}

	if isDividedByAny(11, []int{}) {
		t.Error("11 must not be dividable by nothing")
	}

	if isDividedByAny(11, []int{0}) {
		t.Error("11 must not be dividable by 0")
	}

	if !isDividedByAny(11, []int{0, 1}) {
		t.Error("11 must be dividable by 1")
	}

	if !isDividedByAny(11, []int{0, 1}) {
		t.Error("11 must be dividable by 1")
	}

}

func TestGetSumOfDivider(t *testing.T) {
	var r = getSumOfDividers(99, []int{3, 5})
	if r != 2318 {
		t.Error("The sum of all numbers between 1 and 99 that are dividable by 3 or 5 should be '2318', not:", r)
	}
	var r2 = getSumOfDividers(9, []int{3, 5})
	if r2 != 23 {
		t.Error("The sum of all numbers between 1 and 99 that are dividable by 3 or 5 should be '23', not:", r2)
	}
	var r3 = getSumOfDividers(999, []int{3, 5})
	if r3 != 233168 {
		t.Error("The sum of all numbers between 1 and 999 that are dividable by 3 or 5 should be '233168', not:", r3)
	}
}
