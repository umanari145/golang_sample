package parallel

import (
	"log"
	"time"
)

func heiretsu3() {
	log.Print("started")
	//https://qiita.com/suin/items/82ecb6f63ff4104d4f5d
	sleep_finish := make(chan bool)

	sleepfuncs := []func(){
		func() {
			log.Print("sleep1 start")
			time.Sleep(1 * time.Second)
			log.Print("sleep1 end")
			sleep_finish <- true
		},
		func() {
			log.Print("sleep2 start")
			time.Sleep(2 * time.Second)
			log.Print("sleep2 end")
			sleep_finish <- true
		},
		func() {
			log.Print("sleep3 start")
			time.Sleep(3 * time.Second)
			log.Print("sleep3 end")
			sleep_finish <- true
		},
	}

	//for文でぶん回す
	for _, sleepfunc := range sleepfuncs {
		go sleepfunc()
	}

	//終わらせるチャネル
	for i := 1; i <= 3; i++ {
		<-sleep_finish
	}

	log.Print("all_finished")
}
