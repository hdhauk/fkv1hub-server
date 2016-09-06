package main

import "testing"

func TestEncode(t *testing.T) {
	var tests = []struct {
		input Message
		want  int
	}{
		{Message{11, 22, 3333}, 11223333},
		{Message{22, 33, 4444}, 22334444},
	}

	for _, test := range tests {
		if got, _ := test.input.Encode(); got != test.want {
			t.Errorf("(Message{%d, %d, %d}).Encode() = %d. Want %d",
				test.input.DeviceID, test.input.OpCode, test.input.Payload, got, test.want)
		}
	}
}
