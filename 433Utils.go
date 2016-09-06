package main

import (
	"fmt"
	"time"

	"github.com/hdhauk/fkv1hub-server/RPi433Mhz"
)

// Message defines the messages that can be sendt in the system.
type Message struct {
	// Messages needs to bes parsed into an 8-digit code before sending, using
	// the format described below:
	// 	Message: XXYYZZZZ
	//   ______________________________________
	//	|	DeviceID		OpCode		Payload				 |
	//	| 2 digits		2digits				4 digits	 |
	//	|	XX					YY						ZZZZ			 |
	//   ‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
	//	DeviceID and OpCode must be greater than 0
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
