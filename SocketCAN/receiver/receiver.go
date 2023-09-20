package receiver

import (
	"context"
	"fmt"

	"go.einride.tech/can/pkg/socketcan"
)

func Receive() {
	// _ := recv.Receive()
	// frame := recv.Frame()

	// var auxMsg *etruckcan.Auxiliary
	// _ = auxMsg.UnmarshalFrame(frame)
	conn, _ := socketcan.DialContext(context.Background(), "can", "can0")

	recv := socketcan.NewReceiver(conn)
	for recv.Receive() {
		frame := recv.Frame()
		fmt.Println(frame.String())
	}
}
