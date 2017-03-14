package main

import (
	"fmt"
	"time"

	"github.com/hdhauk/fkv1hub-server/rpi433mhz"
)

func main() {
	tx := make(chan int)
	rx := make(chan int)
	cfg := rpi433mhz.Config{
		TxCh: tx,
		RxCh: rx,
	}
	r := rpi433mhz.NewRCSwitch()
	r.Init(cfg)
	for {
		select {
		case i := <-rx:
			fmt.Println(i)
		}
	}
}

func printRunTime() {
	start := time.Now()
	fmt.Println("FKV1HUB-Server\n------------------")
	for {
		fmt.Printf("Server runtime: %.F seconds\r", time.Since(start).Seconds())
		time.Sleep(1 * time.Second)
	}
}

func dummyHandler(code int) {
	fmt.Printf("Recived code: %d\n", code)
}
