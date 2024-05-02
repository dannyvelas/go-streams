package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func producer(ch chan<- []byte) {
	defer close(ch)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Bytes()
		ch <- line
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func consumer(ch <-chan []byte) {
	for line := range ch {
		fmt.Printf("%d\n", len(line))
	}
}

func main() {
	byteCh := make(chan []byte)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() { producer(byteCh); wg.Done() }()
	go func() { consumer(byteCh); wg.Done() }()

	wg.Wait()
}
