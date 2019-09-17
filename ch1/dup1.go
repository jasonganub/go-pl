// Dup1 prints the text of each line that appears more than once
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	// Ignore potential errors input.Err()
	fmt.Println("Count")
	for text, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, text)
		}
	}
}