package main

import (
	"fmt"
	"os"

	"github.com/macformula/sandbox_projects/receiver"
	"github.com/macformula/sandbox_projects/sender"
	"github.com/macformula/sandbox_projects/simulator"
	"github.com/macformula/sandbox_projects/tracer"
	"go.einride.tech/can/pkg/candevice"
)

// Call with arg either "receiver" or "sender"
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please input 'sender', 'receiver' or 'tracer'")
		return
	}

	action := os.Args[1]
	done := make(chan struct{})
	switch action {
	case "sender":
		setup()
		sender.Send()
	case "receiver":
		setup()
		go receiver.Receive("Receiver 1", done)
		go receiver.Receive("Receiver 2", done)
		<-done
		<-done
	case "tracer":
		setup()
		tracer.Trace()
	case "simulator":
		setup()
		simulator.Simulate()
	default:
		fmt.Println("Invalid input. Please input 'sender', 'receiver' or 'tracer'")
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
