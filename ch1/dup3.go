// Count duplicate text via reading the entire file first and then splitting rather than streaming
package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		for _, text := range strings.Split(string(data), "\n") {
			counts[string(text)]++
		}
	}

	fmt.Println("Counts")
	for text, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, text)
		}
	}
}