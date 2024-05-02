package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(len(line))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading standard input:", err)
	}
}
