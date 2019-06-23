package parallel

import (
	"log"
	"time"
)

func heiretsu() {
	log.Print("started")
	//https://qiita.com/suin/items/82ecb6f63ff4104d4f5d
	sleep1_finish := make(chan bool)
	sleep2_finish := make(chan bool)
	sleep3_finish := make(chan bool)

	go func() {
		log.Print("sleep1 start")
		time.Sleep(1 * time.Second)
		log.Print("sleep1 end")
		sleep1_finish <- true
	}()

	go func() {
		log.Print("sleep2 start")
		time.Sleep(2 * time.Second)
		log.Print("sleep2 end")
		sleep2_finish <- true
	}()

	go func() {
		log.Print("sleep3 start")
		time.Sleep(3 * time.Second)
		log.Print("sleep3 end")
		sleep3_finish <- true
	}()

	<-sleep1_finish
	<-sleep2_finish
	<-sleep3_finish

	log.Print("all_finished")
}
