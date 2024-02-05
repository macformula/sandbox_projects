package sender

import (
	"context"
	"github.com/macformula/sandbox_projects/output/CANBMScan"

	// CANBMScan "github.com/macformula/sandbox_projects/output/CANBMScan"
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
	// packMsg2 := CANBMScan.NewContactor_Feedback().SetPack_Negative_Feedback(false)
	// frame2 := packMsg2.Frame()
	// Gives me the signal: CANBMScan.NewContactor_Feedback().Pack_Negative_Feedback()

	var frames []can.Frame

	for i := 0; i <= 7; i++ {
		// Create a new can.Frame for each iteration with the appropriate binary representation
		//frame := can.Frame{
		//	ID:     1574,
		//	Length: 1,
		//	Data:   can.Data{byte(i)},
		//}
		frame := frame1
		frames = append(frames, frame)
	}

	tx := socketcan.NewTransmitter(conn)
	// Print the frames to verify the result
	for _, frame := range frames {
		if err := tx.TransmitFrame(context.Background(), frame); err != nil {
			panic(err)
		}
	}

	// frame := can.Frame{
	// 	ID:     1572,
	// 	Length: 1,
	// 	Data:   can.Data{0b111},
	// }

	// if err := tx.TransmitFrame(context.Background(), frame1); err != nil {
	// 	panic(err)
	// }
	// if err := tx.TransmitFrame(context.Background(), frame2); err != nil {
	// 	panic(err)
	// }
}
