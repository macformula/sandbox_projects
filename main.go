package main

import (
	"context"
	"go.einride.tech/can/pkg/socketcan"
	"net"
	"os"
	"time"
)

const (
	log1 = "log1.txt"
	log2 = "log2.txt"
	can0 = "can0"
)

type Receiver struct {
	receiver *socketcan.Receiver
	ctx      context.Context
	file     *os.File
}

func main() {
	var l1, l2 *os.File
	var err error
	var ctx context.Context = context.Background()

	l1, err = os.Create(log1)
	if err != nil {
		panic(err)
	}
	defer l1.Close()

	l2, err = os.Create(log2)
	if err != nil {
		panic(err)
	}
	defer l2.Close()

	conn1, err := socketcan.DialContext(ctx, "can", can0)
	if err != nil {
		panic(err)
	}

	conn2, err := socketcan.DialContext(ctx, "can", can0)
	if err != nil {
		panic(err)
	}

	var receive1 Receiver = Receiver{
		file:     l1,
		receiver: getReceiver(conn1, ctx),
		ctx:      ctx,
	}

	var receive2 Receiver = Receiver{
		file:     l2,
		receiver: getReceiver(conn2, ctx),
		ctx:      ctx,
	}

	go receive1.receive()
	go receive2.receive()

	time.Sleep(5 * time.Second)

	ctx.Done()

}

func getReceiver(conn net.Conn, ctx context.Context) *socketcan.Receiver {
	return socketcan.NewReceiver(conn)
}

func (r *Receiver) receive() {
	select {
	case <-r.ctx.Done():
		return
	default:
		if r.receiver.Receive() {
			frame := r.receiver.Frame()

			r.file.Write(frame.Data[:])
		}
	}
}
