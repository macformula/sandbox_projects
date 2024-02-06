package receiver

import (
	"context"
	"go.einride.tech/can"
	"go.einride.tech/can/pkg/generated"
	"go.einride.tech/can/pkg/socketcan"
	"net"
)

type CANClient struct {
	md   MessagesDescriptor
	rx   *socketcan.Receiver
	stop chan int
}

type MessagesDescriptor interface {
	UnmarshalFrame(f can.Frame) (generated.Message, error)
}

func NewCANClient(messages MessagesDescriptor, conn net.Conn) CANClient {
	c := CANClient{
		md: messages,
		rx: socketcan.NewReceiver(conn),
	}
	return c
}

func (c *CANClient) receive(ctx context.Context, msgToRead ...generated.Message) {
	for c.rx.Receive() {
		frame := c.rx.Frame()
		msg, err := c.md.UnmarshalFrame(frame)
		if err != nil {
			return err
		}

		if frame.ID == msgToRead.Frame().ID {
			msgToRead = msg
		}
	}
}

func (c *CANClient) Open() {
	// ...
	go c.receive()
}
