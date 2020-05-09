// Inspired by unix's `uniq` command

// Prints the text of each line that appears more than once
// in the standard input, preceded by its count.

// Written to be equivalent to `python`'s `streaming` mode -- where
// data is read into memory per each iteration, and not at once.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg) //`arg` is a file
			if err != nil {
				fmt.Fprint(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	// Iterate through the map, printing if the value
	// is more than one. (i.e. not 0)
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

// using `test.txt` as input, should return:
// 3       olumide
// 5       oluchi
