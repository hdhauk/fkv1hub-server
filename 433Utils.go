package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// Message defines the messages that can be sendt in the system.
// Messages needs to bes parsed into an 8-digit code before sending, using
// the format described below:
// 	Message: XXYYZZZZ
//   ______________________________________
//	|	DeviceID		OpCode		Payload				 |
//	| 2 digits		2digits				4 digits	 |
//	|	XX					YY						ZZZZ			 |
//   ‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
//	DeviceID and OpCode must be greater than 0
type Message struct {
	DeviceID int // 2 digits
	OpCode   int // 2 digits
	Payload  int // 4 digits
}

// Encode returns the message as a single int, and and error if OpCode
// or DeviceID is not set.
// TODO: Write test for this
func (m *Message) Encode() (int, error) {
	if m.DeviceID <= 0 {
		err := fmt.Errorf("DeviceID not set/or invalid. DeviceID = %d", m.DeviceID)
		return 0, err
	}
	if m.OpCode <= 0 {
		err := fmt.Errorf("OpCode not set/or invalid. OpCode = %d", m.OpCode)
		return 0, err
	}
	return m.DeviceID*100000 + m.OpCode*10000 + m.Payload, nil
}

// Decode parses a raw message onto an existing Message object
func (m *Message) Decode(rawCode int) error {

}

// OpCodes
const (
	// Temperature
	GetTemp = 10
	SetTemp = 11

	// Humidity
	GetHumid = 20
)

// For using labels on raspberry-pi breakout board
const (
	SDA    = 8
	SCL    = 9
	GPIO04 = 7
	GPIO17 = 0
	GPIO27 = 2
	GPIO22 = 3
	MOSI   = 12
	MISO   = 13
	SCLK   = 14

	TxD    = 15
	RxD    = 16
	GPIO18 = 1
	GPIO23 = 4
	GPIO24 = 5
	GPIO25 = 6
	CE0    = 10
	CE1    = 11
)

// RxCodeHandler is a function that run concurrently when a 433Mhz
// 8-bit code is recieved. It takes the code as an input
type RxCodeHandler func(uint)

func listenAndRespond(pin int, handler RxCodeHandler) {
	for {
		code := listen(pin)
		go handler(code)
		time.Sleep(500 * time.Millisecond) // Make sure signal has died out...
	}
}

func listen(pin int) int {
	// For correct pinout: https://projects.drogon.net/raspberry-pi/wiringpi/pins/
	// or use constants declared above
	pinStr := strconv.Itoa(pin)
	// TODO: Configure this to go in a config file
	// TODO: Test that choosing a diffrent pin works
	out, err := exec.Command("/home/pi/Documents/C++/433Utils/RPi_utils/SingleSniffer", pinStr).Output()
	if err != nil {
		log.Fatal(err)
	}
	code, err := strconv.Atoi(strings.TrimSpace(string(out)))
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return code
}
