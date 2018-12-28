package main

import (
	"github.com/SQLek/voip/pkg/sip"
	"log"
	"os"
)

func main() {

	toStr := "localhost:5060"
	if len(os.Args) > 1 {
		toStr = os.Args[1]
	}

	c, err := sip.Dial(toStr)
	if err != nil {
		log.Fatal("Dialog open failed! ", err)
	}

	err = c.DemoWrite()
	if err != nil {
		log.Fatal("Dialog write failed! ", err)
	}

	err = c.DemoRead()
	if err != nil {
		log.Fatal("Dialog read failed! ", err)
	}

}
