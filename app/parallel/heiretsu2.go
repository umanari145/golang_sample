package parallel

import (
	"log"
	"time"
)

func heiretsu2() {
	log.Print("started")
	//https://qiita.com/suin/items/82ecb6f63ff4104d4f5d
	sleep_finish := make(chan bool)

	go func() {
		log.Print("sleep1 start")
		time.Sleep(1 * time.Second)
		log.Print("sleep1 end")
		sleep_finish <- true
	}()

	go func() {
		log.Print("sleep2 start")
		time.Sleep(2 * time.Second)
		log.Print("sleep2 end")
		sleep_finish <- true
	}()

	go func() {
		log.Print("sleep3 start")
		time.Sleep(3 * time.Second)
		log.Print("sleep3 end")
		sleep_finish <- true
	}()

	//もし一つだけチャネルだと1つが終わった時点で終了してしまう
	//<-sleep_finish

	//上記の理由からチェネル分のfinishを作る
	for i := 1; i <= 3; i++ {
		<-sleep_finish
	}

	log.Print("all_finished")
}
