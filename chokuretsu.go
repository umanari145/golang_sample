package main

import (
	"log"
	"time"
)

func main() {
	log.Print("started")

	log.Print("sleep1 start")
	time.Sleep(1 * time.Second)
	log.Print("sleep1 end")

	log.Print("sleep2 start")
	time.Sleep(2 * time.Second)
	log.Print("sleep2 end")

	log.Print("sleep3 start")
	time.Sleep(3 * time.Second)
	log.Print("sleep3 end")
}
