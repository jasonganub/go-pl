// Dup1 prints the text of each line that appears more than once from a list of files or via input
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	dupFiles := make(map[string][]string)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, "", dupFiles)
	} else {
		for _, filename := range files {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, filename, dupFiles)
			f.Close()
		}
	}

	fmt.Println("Counts")
	for text, count := range counts {
		fmt.Printf("%d\t%s\tfiles: %s\n", count, text, strings.Join(dupFiles[text], ","))
	}
}

func countLines(f *os.File, counts map[string]int, filename string, dupFiles map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++

		containsTextAlready := false
		for _, filenameInDupFiles := range dupFiles[input.Text()] {
			if filenameInDupFiles == filename {
				containsTextAlready = true
			}
		}

		if counts[input.Text()] > 1 && !containsTextAlready {
			dupFiles[input.Text()] = append(dupFiles[input.Text()], filename)
		}
	}
	// Ignoring potential errors from input.Err()
}