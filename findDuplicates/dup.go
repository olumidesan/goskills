// Inspired by unix's `uniq` command

// Prints the text of each line that apperars more than once
// in the standard input, preceded by its count

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() { // For each input
		// Set the key to the input and increase the value
		counts[input.Text()]++
	}
	// Iterate through the map, printing if the value
	// is more than one. (i.e. not 0)
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
