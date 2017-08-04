package main

import (
	"log"
)

type astNode struct {
	value token
	edges []*astNode
}

func parse(ctl <-chan state, in <-chan []token) (<-chan state, <-chan *astNode) {
	outctl := make(chan state)
	out := make(chan *astNode)
	var rootNode *astNode
	go func() {
		defer close(out)
		log.Println("init parse")

		for {
			aState, ts := <-ctl, <-in
			rootNode = &astNode{edges: make([]*astNode, len(ts))}
			for i, t := range ts {
				n := &astNode{value: t}
				rootNode.edges[i] = n
			}

			log.Println("send node")
			outctl <- aState
			out <- rootNode
		}
	}()
	return outctl, out
}
