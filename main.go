package main

import (
	"fmt"
	"os"

	"github.com/tejal29/gosink/pkg/runner"
)

func main() {
	// Running in Sequence
	fmt.Println("Running in Sequence [3, 1, 5,8]\n")
	runner.Run(false, []int{3, 1, 5, 8}, os.Stdout)

	// Running in Parallel
	fmt.Println("Running in Parallel [3, 1, 5,8]\n")
	runner.Run(true, []int{3, 1, 5, 8}, os.Stdout)
	fmt.Println("\nBye Bye!")
}
