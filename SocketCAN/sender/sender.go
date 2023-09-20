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

	frame := can.Frame{}
	tx := socketcan.NewTransmitter(conn)
	_ = tx.TransmitFrame(context.Background(), frame)
}
