package main

import (
	"fmt"
	"log"
	"strings"
)

func initEval(ctl <-chan state, in <-chan *astNode) (<-chan state, <-chan string) {
	outctl := make(chan state)
	out := make(chan string)
	go eval(ctl, in, outctl, out)
	return outctl, out
}

func eval(ctl <-chan state, in <-chan *astNode, outctl chan state, out chan string) {
	defer close(out)
	log.Println("init eval")

	for {
		aState, n := <-ctl, <-in
		vals := make([]string, len(n.edges))
		for i, e := range n.edges {
			vals[i] = fmt.Sprintf("\"%v\"", e)
		}
		outctl <- aState
		out <- strings.Join(vals, ":")
	}
}
