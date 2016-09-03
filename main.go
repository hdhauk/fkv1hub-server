package main

import (
	"time"

	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/all"
)

func main() {
	embd.InitGPIO()
	defer embd.CloseGPIO()

	embd.SetDirection(10, embd.Out)
	for {
		embd.DigitalWrite(10, embd.High)
		time.Sleep(1 * time.Second)
		embd.DigitalWrite(10, emdb.Low)
		time.Sleep(1 * time.Second)
	}

}
