package main

import "testing"

func TestEncode(t *testing.T) {
	var tests = []struct {
		input Message
		want  int
	}{
		{Message{11, 22, 3333}, 11223333},
		{Message{22, 33, 4444}, 22334444},
		{Message{10, GetTemp, 0000}, 10100000},
		{Message{11, SetTemp, 2222}, 11112222},
		{Message{22, GetHumid, 0000}, 22200000},
		{Message{1, 1, 0}, 1010000},
		{Message{0, 0, 0}, 0},
	}

	for _, test := range tests {
		if got, _ := test.input.Encode(); got != test.want {
			t.Errorf("(Message{%d, %d, %d}).Encode() = %d. Want %d",
				test.input.DeviceID, test.input.OpCode, test.input.Payload, got, test.want)
		}
	}
}
