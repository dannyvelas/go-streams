package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sync"
)

func producer(writer *io.PipeWriter) {
	defer writer.Close()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if _, err := writer.Write([]byte(line + "\n")); err != nil {
			fmt.Fprintf(os.Stderr, "error writing to writer: %v\n", err)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading standard input: %v\n", err)
	}
}

func consumer(reader *io.PipeReader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Bytes()
		fmt.Printf("%d\n", len(line))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading from reader: %v\n", err)
	}
}

func main() {
	reader, writer := io.Pipe()
	var wg sync.WaitGroup
	wg.Add(2)

	go func() { producer(writer); wg.Done() }()
	go func() { consumer(reader); wg.Done() }()

	wg.Wait()
}
