package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func initInput(in <-chan state) (<-chan state, <-chan state, <-chan string) {
	quitctl := make(chan state)
	outctl := make(chan state)
	out := make(chan string)
	go input(in, quitctl, outctl, out)
	return quitctl, outctl, out
}

func input(in <-chan state, quitctl <-chan state, outctl chan state, out chan string) {
	defer close(out)
	var scanner bufio.Scanner
	log.Println("init input")
	for {
		log.Println("wait for, and receive state")
		aState := <-in
		log.Println("received control signal")

		log.Println("prompt for input")
		fmt.Print("> ")

		scanner = *bufio.NewScanner(os.Stdin)

		if ok := scanner.Scan(); ok {
			log.Println("read input line")

			// For some reason, sending the state must happen before
			// the sending of the scanned text, or we panic (deadlock)
			// TODO: investigate
			outctl <- aState
			out <- scanner.Text()
		}
	}
}
