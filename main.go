/*
main.go - Entrypoint to the CLI application.

CLI Application for managing viewing of Euler Problem results based on question number
from the site: https://projecteuler.net/archives

Author: Matthew Sunner, 2025
*/
package main

import "fmt"

func Sum(intSlice []int) int {
	sum := 0
	for _, v := range intSlice {
		sum += v
	}
	return sum
}

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
}

func main() {
	fmt.Println("Hello World!")

}
