package main

import (
	"fmt"
	"time"

	"github.com/hdhauk/fkv1hub-server/RPi433Mhz"
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
func (m *Message) Encode() (int, error) {
	if m.DeviceID <= 0 {
		err := fmt.Errorf("DeviceID not set/or invalid. DeviceID = %d", m.DeviceID)
		return 0, err
	}
	if m.OpCode <= 0 {
		err := fmt.Errorf("OpCode not set/or invalid. OpCode = %d", m.OpCode)
		return 0, err
	}
	return m.DeviceID*1000000 + m.OpCode*10000 + m.Payload, nil
}

// Decode parses a raw message onto an existing Message object
func (m *Message) Decode(rawCode int) error {
	m.Payload = rawCode % 10000
	m.OpCode = ((rawCode - m.Payload) % 1000000) / 10000
	m.DeviceID = ((rawCode - m.Payload - m.OpCode) % 100000000) / 1000000
	if m.OpCode <= 0 {
		return fmt.Errorf("Cannot parse OpCode. Opcode =  %d", m.OpCode)
	}
	if m.DeviceID <= 0 {
		return fmt.Errorf("Cannot parse DeviceID. DeviceID =  %d", m.DeviceID)
	}
	return nil
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
type RxCodeHandler func(int)

func listenAndRespond(pin int, handler RxCodeHandler) {
	for {
		code := RPi433Mhz.Listen(pin)
		go handler(code)
		time.Sleep(500 * time.Millisecond) // Make sure signal has died out...
	}
}
