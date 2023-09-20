package sender

import (
	"context"

	"go.einride.tech/can"
	"go.einride.tech/can/pkg/socketcan"
)

func Send() {
	// auxMsg := etruckcan.NewAuxiliary().SetHeadLights(etruckcan.Auxiliary_HeadLights_LowBeam)
	// frame := auxMsg.Frame()
	conn, _ := socketcan.DialContext(context.Background(), "can", "can0")

	var data *can.Data
	data.SetBit(0, true)
	data.SetBit(1, true)
	data.SetBit(2, true)
	data.SetBit(3, true)
	data.SetBit(4, true)
	data.SetBit(5, true)
	data.SetBit(6, true)
	data.SetBit(7, true)
	frame := can.Frame{
		ID:     456,
		Length: 1,
		Data:   *data,
	}
	tx := socketcan.NewTransmitter(conn)
	_ = tx.TransmitFrame(context.Background(), frame)
}
