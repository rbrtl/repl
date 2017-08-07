package main

import "strings"
import "log"

type token string

func initLex(ctl <-chan state, in <-chan string) (<-chan state, <-chan []token) {
	outctl := make(chan state)
	out := make(chan []token)
	go lex(ctl, in, outctl, out)
	return outctl, out
}

func lex(ctl <-chan state, in <-chan string, outctl chan state, out chan []token) {
	defer close(out)
	log.Println("init lex")
	for {
		aState, s := <-ctl, <-in

		ss := strings.Split(s, " ")

		log.Println("tokenise input")
		t := make([]token, len(ss))
		for i, it := range ss {
			t[i] = token(it)
		}

		log.Println("send tokens")
		outctl <- aState
		out <- t
	}
}
