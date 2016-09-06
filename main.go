package main

import (
	"fmt"
	"time"

	_ "github.com/kidoman/embd/host/rpi"
)

func main() {
	go printRunTime()
	//go blinkLED()
	listenAndRespond(GPIO18, dummyHandler)

}

func printRunTime() {
	start := time.Now()
	fmt.Println("FKV1HUB-Server\n------------------")
	for {
		fmt.Printf("Server runtime: %.F seconds\r", time.Since(start).Seconds())
		time.Sleep(1 * time.Second)
	}
}

func dummyHandler(code uint) {
	fmt.Printf("Recived code: %d\n", code)
}
