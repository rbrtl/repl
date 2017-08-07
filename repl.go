package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

type state struct{}

func main() {
	log.SetFlags(0)
	log.SetOutput(ioutil.Discard)
	ctlChan := make(chan state)
	log.Println("init ctlChan")

	quitChan, lexcChan, inChan := initInput(ctlChan)
	parcChan, lexChan := initLex(lexcChan, inChan)
	evlcChan, parChan := initParse(parcChan, lexChan)
	outcChan, valChan := initEval(evlcChan, parChan)
	initOutput(ctlChan, outcChan, valChan)

	log.Println("send control signal")
	ctlChan <- state{}

	<-quitChan
	log.Println("quitChan recv")
	fmt.Println("Bye! :)")
}
