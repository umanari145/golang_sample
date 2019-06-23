package parallel

import (
	"log"
	"time"
)

func heiretsu4() {
	log.Print("started")
	//https://qiita.com/TakaakiFuruse/items/241578174fd2f00aaa8a
	sleep_finish := make(chan bool)

	//稼働
	go sleepfunc1(sleep_finish)
	go sleep2func2(sleep_finish)

	//終わらせるチャネル
	for i := 1; i <= 2; i++ {
		<-sleep_finish
	}
	log.Print("all_finished")
}

func sleepfunc1(sleep_finish chan bool) {
	log.Print("sleep1 start")
	time.Sleep(1 * time.Second)
	log.Print("sleep1 end")
	sleep_finish <- true
}

func sleep2func2(sleep_finish chan bool) {
	log.Print("sleep2 start")
	time.Sleep(2 * time.Second)
	log.Print("sleep2 end")
	sleep_finish <- true
}
