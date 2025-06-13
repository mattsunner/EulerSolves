/*
probs - Package containing the logic for solving Euler problems.
This package contains the functions to solve the problems and return the results.

Author: Matthew Sunner, 2025
*/

package probs

func Sum(intSlice []int) int {
	sum := 0
	for _, v := range intSlice {
		sum += v
	}
	return sum
}

// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
// The sum of these multiples is 23.
//
// Find the sum of all the multiples of 3 or 5 below 1000.
//
// https://projecteuler.net/problem=1

func ProbOne(countOf int, maxOf int) int {
	solutionsList := make([]int, 0, countOf)

	for i := 0; i < maxOf; i++ {
		if i%3 == 0 || i%5 == 0 {
			solutionsList = append(solutionsList, i)
		}
	}

	if len(solutionsList) == 0 {
		return 0
	}

	return Sum(solutionsList)

	// Answer: 233,168
}
