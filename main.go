/*
main.go - Entrypoint to the CLI application.

CLI Application for managing viewing of Euler Problem results based on question number
from the site: https://projecteuler.net/archives

Author: Matthew Sunner, 2025
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	probs "github.com/mattsunner/eulersolves/internal"
)

func printMenu() {
	fmt.Println("Welcome to the Euler Problem Solver CLI!")
	fmt.Println("Please select a problem to view the answer for: ")
	fmt.Println("1. Problem 1")
	fmt.Println("2. Problem 2")
}

func getUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(prompt + ": ")
	input, _ := reader.ReadString('\n')

	return strings.TrimSpace(input)
}

func main() {
	printMenu()
	choice := getUserInput("Enter your choice: ")

	switch strings.ToLower(choice) {
	case "1":
		fmt.Println("You selected Problem 1.")
		res := probs.ProbOne(0, 1000)
		fmt.Println("The answer is: ", res)
		return
	case "2":
		fmt.Println("You selected Problem 2.")
		res := probs.ProbTwo(4000000)
		fmt.Println("The answer is: ", res)
	default:
		fmt.Println("Invalid choice. Please try again.")
	}
}
