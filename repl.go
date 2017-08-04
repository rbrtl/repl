package main

import (
	"fmt"
	"log"
)

type state struct{}

func main() {
	ctlChan := make(chan state)
	log.Println("init ctlChan")

	quitChan, lexcChan, inChan := input(ctlChan)
	parcChan, lexChan := lex(lexcChan, inChan)
	evlcChan, parChan := parse(parcChan, lexChan)
	outcChan, valChan := eval(evlcChan, parChan)
	output(ctlChan, outcChan, valChan)

	log.Println("send control signal")
	ctlChan <- state{}

	<-quitChan
	log.Println("quitChan recv")
	fmt.Println("Bye! :)")
}
