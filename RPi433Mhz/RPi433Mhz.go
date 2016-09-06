package RPi433Mhz

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

// Listen for 433Mhz codes on specified pin
func Listen(pin int) int {
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
