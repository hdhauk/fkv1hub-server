#include "../rc-switch/RCSwitch.h"
#include <stdio.h>
#include <stdlib.h>

RCSwitch mySwitch;

int init() {
  int PIN = 2;
  if (wiringPiSetup() == -1) {
    return 1
  }

  int pulseLength = 0;

  mySwitch = RCSwitch();
  if (pulseLength != 0)
    mySwitch.setPulseLength(pulseLength);

  mySwitch.enableReceive(PIN);
  return 0
}
