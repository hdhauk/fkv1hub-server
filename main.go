package main

import (
	"fmt"
	"time"

	"github.com/hdhauk/fkv1hub-server/RPi433Mhz"
)

func main() {
	go printRunTime()
	listenAndRespond(RPi433Mhz.GPIO18, dummyHandler)

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
