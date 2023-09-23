package sender

import (
	"context"

	CANBMScan "github.com/macformula/sandbox_projects/output/CANBMScan"
	"go.einride.tech/can"
	"go.einride.tech/can/pkg/socketcan"
)

func Send() {
	// auxMsg := etruckcan.NewAuxiliary().SetHeadLights(etruckcan.Auxiliary_HeadLights_LowBeam)
	// frame := auxMsg.Frame()

	conn, err := socketcan.DialContext(context.Background(), "can", "can0")
	if err != nil {
		panic(err)
	}

	packMsg1 := CANBMScan.NewContactor_Feedback().SetPack_Negative_Feedback(true)
	frame1 := packMsg1.Frame()
	packMsg2 := CANBMScan.NewContactor_Feedback().SetPack_Negative_Feedback(false)
	frame2 := packMsg2.Frame()
	// Gives me the signal: CANBMScan.NewContactor_Feedback().Pack_Negative_Feedback()

	frame := can.Frame{
		ID:     0x123,
		Length: 8,
		Data:   can.Data{'h', 'i', ' ', 'w', 'o', 'r', 'l', 'd'},
	}
	tx := socketcan.NewTransmitter(conn)
	if err := tx.TransmitFrame(context.Background(), frame); err != nil {
		panic(err)
	}
	if err := tx.TransmitFrame(context.Background(), frame1); err != nil {
		panic(err)
	}
	if err := tx.TransmitFrame(context.Background(), frame2); err != nil {
		panic(err)
	}
}
