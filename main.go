package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi"
)

func main() {
	go printRunTime()
	//go blinkLED()

	cmd := exec.Command("sudo $HOME/Documents/C++/433Utils/RPi_utils/RFSniffer")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(stdout)

}

func printRunTime() {
	start := time.Now()
	fmt.Println("FKV1HUB-Server\n------------------")
	for {
		fmt.Printf("Server runtime: %.F seconds\r", time.Since(start).Seconds())
		time.Sleep(1 * time.Second)
	}
}

func blinkLED() {
	embd.InitGPIO()
	defer embd.CloseGPIO()

	pin, err := embd.NewDigitalPin(10)
	if err != nil {
		fmt.Println("Something went wrong when assigning new pin")
		fmt.Println(err)
	}
	pin.SetDirection(embd.Out)
	for {
		pin.Write(embd.High)
		time.Sleep(2 * time.Second)
		pin.Write(embd.Low)
		time.Sleep(2 * time.Second)
	}
}
