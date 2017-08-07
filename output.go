package main

import "fmt"
import "log"

func initOutput(ctl chan<- state, s <-chan state, in <-chan string) {
	go output(ctl, s, in)
}

func output(ctl chan<- state, s <-chan state, in <-chan string) {
	log.Println("init output")
	for {
		aState, value := <-s, <-in
		log.Println("output recv")
		fmt.Println(value)

		log.Println("send control state")
		ctl <- aState
	}
}
