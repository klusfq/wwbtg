package main

import (
	"log"
)

func main() {

	run()

	log.Println("after run")
}

func run() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("cover Panic: ", err)
		}
	}()

	// log.Fatal("this is log fatal")
	log.Panic("this is log panic")
}
