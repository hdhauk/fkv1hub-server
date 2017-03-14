package rpi433mhz

/*
#include "../rc-switch/RCSwitch.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import (
	"log"
	"os"
	"time"
)

// RCSwitch is the driver handle.
type RCSwitch struct {
	mySwitch C.RCSwitch
	txCh     chan int
	rxCh     chan int
	abortCh  chan struct{}
	logger   *log.Logger
}

// NewRCSwitch returns a new RCSwitch handle.
func NewRCSwitch() *RCSwitch {
	var ret RCSwitch
	ret.mySwitch = C.RCSwitch()
	return &ret
}

func (r *RCSwitch) available() bool {
	return r.mySwitch.available()
}

func (r *RCSwitch) getReceivedValue() int {
	return r.mySwitch.getReceivedValue()
}

func (r *RCSwitch) resetAvailable() {
	r.mySwitch.resetAvailable()
}

func (r *RCSwitch) rx() {
	r.logger.Println("[INFO] Starting listener.")
	for {
		select {
		case <-r.abortCh:
			r.logger.Println("[INFO] Stopping listener.")
			return
		case <-time.After(50 * time.Millisecond):
		}

		if r.available() {
			value := r.getReceivedValue()
			if value == 0 {
				r.logger.Println("[WARN] Unknown encoding.")
			} else {
				r.rxCh <- value
			}
			r.resetAvailable()
		}
	}
}

// Config defines configuration.
type Config struct {
	TxCh   chan int
	RxCh   chan int
	Logger *log.Logger
}

// Init initiates the driver.
func (r *RCSwitch) Init(c Config) {
	if c.Logger != nil {
		r.logger = log.New(os.Stdout, "[rpi433mhz]", log.LstdFlags)
	} else {
		r.logger = c.Logger
	}

	r.rxCh = c.RxCh
	r.txCh = c.TxCh

	go r.rx()

}
