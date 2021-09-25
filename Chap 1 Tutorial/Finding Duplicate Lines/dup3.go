package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	count := make(map[string]int)

	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
		}

		for _, line := range strings.Split(string(data), "\n") {
			count[line]++
		}
	}

	for line, value := range count {
		if value > 1 {
			fmt.Printf("Count: %d\t Line: %s\n", value, line)
		}
	}
}
