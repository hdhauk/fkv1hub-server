package RPi433Mhz

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

// Breakout board GPIO pins
const (
	// Mapping according to: https://projects.drogon.net/raspberry-pi/wiringpi/pins/

	// Left Side
	SDA    = 8
	SCL    = 9
	GPIO04 = 7
	GPIO17 = 0
	GPIO27 = 2
	GPIO22 = 3
	MOSI   = 12
	MISO   = 13
	SCLK   = 14

	// Right Side
	TxD    = 15
	RxD    = 16
	GPIO18 = 1
	GPIO23 = 4
	GPIO24 = 5
	GPIO25 = 6
	CE0    = 10
	CE1    = 11
)

// Location of RPi_utils
// TODO: Configure this to go in a config file
const utilsPath = "/home/pi/Documents/C++/433Utils/RPi_utils/"

// Listen for 433Mhz codes on specified GPIO-pin
func Listen(gpio int) int {
	pinStr := strconv.Itoa(gpio)
	// TODO: Test that choosing a diffrent gpio works
	commandStr := fmt.Sprintf("%s%s", utilsPath, "SingleSniffer")
	out, err := exec.Command(commandStr, pinStr).Output()
	//out, err := exec.Command("/home/pi/Documents/C++/433Utils/RPi_utils/SingleSniffer", pinStr).Output()
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

// Send an (up to) 8-digit code over 433Mhz on specified GPIO-pin
func Send(gpio, code int) error {
	pinStr := strconv.Itoa(gpio)
	codeStr := fmt.Sprintf("%08d", code)

	commandStr := fmt.Sprintf("%s%s", utilsPath, "CodeSend")
	out, err := exec.Command(commandStr, codeStr, pinStr).Output()
	if err != nil {
		log.Println(err)
		return err
	}
	if string(out) != fmt.Sprintf("sending code[%s]", codeStr) {
		log.Println(out)
		return fmt.Errorf("Output indicates a problem.\nOutput: %s", out)
	}
	return nil
}
