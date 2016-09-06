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

func TestDecode(t *testing.T) {
	var tests = []struct {
		input int
		want  Message
	}{
		{11223333, Message{11, 22, 3333}},
		{22331234, Message{22, 33, 1234}},
	}

	for _, test := range tests {
		msg := Message{0, 0, 0}
		msg.Decode(test.input)
		if msg != test.want {
			t.Errorf(".Decode(%d) = Message{%d, %d, %d}. Want Message{%d, %d, %d}",
				test.input,
				msg.DeviceID, msg.OpCode, msg.Payload,
				test.want.DeviceID, test.want.OpCode, test.want.Payload)
		}
	}
}
