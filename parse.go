package main

import (
	"log"
)

type astNode struct {
	value token
	edges []*astNode
}

func initParse(ctl <-chan state, in <-chan []token) (<-chan state, <-chan *astNode) {
	outctl := make(chan state)
	out := make(chan *astNode)
	go parse(ctl, in, outctl, out)
	return outctl, out
}

func parse(ctl <-chan state, in <-chan []token, outctl chan state, out chan *astNode) {
	defer close(out)
	log.Println("init parse")

	var rootNode *astNode

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
}
