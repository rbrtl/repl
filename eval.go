package main

import "strings"
import "log"

func eval(ctl <-chan state, in <-chan *astNode) (<-chan state, <-chan string) {
	outctl := make(chan state)
	out := make(chan string)
	go func() {
		defer close(out)
		log.Println("init eval")

		for {
			aState, n := <-ctl, <-in
			vals := make([]string, len(n.edges))
			for i, e := range n.edges {
				vals[i] = string(e.value)
			}
			outctl <- aState
			out <- strings.Join(vals, " ")
		}
	}()
	return outctl, out
}
