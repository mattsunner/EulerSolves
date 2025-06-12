/*
main.go - Entrypoint to the CLI application.

CLI Application for managing viewing of Euler Problem results based on question number
from the site: https://projecteuler.net/archives

Author: Matthew Sunner, 2025
*/
package main

import (
	"fmt"

	probs "github.com/mattsunner/eulersolves/internal"
)

func main() {
	val := probs.ProbOne(0, 10)
	fmt.Println("The result of ProbOne is:", val)
}
