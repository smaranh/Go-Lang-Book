package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("Accepting inputs")

	// ctrl+D to exit the loop in Mac/Linux and ctrl+Z to exit in Windows
	for input.Scan() {
		counts[input.Text()]++
	}
	if err := input.Err(); err != nil {
		fmt.Println("There was an error, ", err)
	}

	fmt.Println("Printing output")
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("Count: %d\t Line: %s\n", n, line)
		}
	}
}
