/**
 * Solution for Problem 01 from Project Euler (https://projecteuler.net/problem=1)
 * If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
 * The sum of these multiples is 23.
 * Find the sum of all the multiples of 3 or 5 below 1000.
 */
package main

import (
	"fmt"
)

/**
 * Checks if 'number' is divided by any in the slice 'dividers'.
 * Skipps divisions by 0.
 */
func isDividedByAny(number int, dividers []int) bool {
	for i := 0; i < len(dividers); i++ {
		if dividers[i] == 0 {
			continue
		}
		if number%dividers[i] == 0 {
			return true
		}
	}
	return false
}

/**
 * Iterates from 1 to 'upper' and sums up all numbers that are
 * dividable by any of the 'dividers' provided.
 */
func getSumOfDividers(upper int, dividers []int) int {
	sum := 0
	for i := 1; i <= upper; i++ {
		if isDividedByAny(i, dividers) {
			sum += i
		}
	}
	return sum
}

func main() {
	fmt.Println(getSumOfDividers(999, []int{3, 5}))
}
