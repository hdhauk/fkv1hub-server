package main

/*
func blinkLED() {
	embd.InitGPIO()
	defer embd.CloseGPIO()

	pin, err := embd.NewDigitalPin(10)
	if err != nil {
		fmt.Println("Something went wrong when assigning new pin")
		fmt.Println(err)
	}
	pin.SetDirection(embd.Out)
	for {
		pin.Write(embd.High)
		time.Sleep(2 * time.Second)
		pin.Write(embd.Low)
		time.Sleep(2 * time.Second)
	}
}
*/
