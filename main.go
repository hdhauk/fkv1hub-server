package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"time"

	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi"
)

func main() {
	go printRunTime()
	//go blinkLED()
	listenAndServe(1, dummyHandler)

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

func listenAndRespond(pin int, handler func(int)) {
	for {
		code := listen(pin)
		go handler(code)
		time.Sleep(500 * time.Millisecond) // Make sure signal has died out...
	}
}

func listen(pin int) int {
	// For correct pinout: https://projects.drogon.net/raspberry-pi/wiringpi/pins/
	pinStr := strconv.Itoa(pin)
	out, err := exec.Command("/home/pi/C++/433Utils/RPi_utils/SingleSniffer ", pinStr).Output()
	if err != nil {
		log.Fatal(err)
	}
	code, err := binary.ReadVarint(out)
	if err != nil {
		return -1
	}
	return code
}

func dummyHandler(code int) {
	fmt.Printf("Recived code: %d\n", code)
}
