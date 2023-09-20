package main

import (
	"fmt"
	"os"

	"github.com/macformula/sandbox_projects/receiver"
	"github.com/macformula/sandbox_projects/sender"
	"go.einride.tech/can/pkg/candevice"
)

// Call with arg either "receiver" or "sender"
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please input 'sender' or 'receiver'")
		return
	}

	action := os.Args[1]
	switch action {
	case "sender":
		setup()
		sender.Send()
	case "receiver":
		setup()
		receiver.Receive()
	default:
		fmt.Println("Invalid input. Please input 'sender' or 'receiver'")
	}
}

func setup() {
	fmt.Println("Setting up can0")
	d, _ := candevice.New("can0")
	_ = d.SetBitrate(250000)
	_ = d.SetUp()
	defer d.SetDown()
	fmt.Println("Done can0 setup")
}