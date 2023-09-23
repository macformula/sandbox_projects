package receiver

import (
	"context"
	"fmt"

	"go.einride.tech/can/pkg/socketcan"
	CANAMKInvertercan "github.com/macformula/sandbox_projects/output/CANAMKInvertercan"
	CANBMScan "github.com/macformula/sandbox_projects/output/CANBMScan"
)

func Receive() {
	conn, err := socketcan.DialContext(context.Background(), "can", "can0")
	if err != nil {
			panic(err)
	}

	// Messages
	AMK1_SetPoints1 := CANAMKInvertercan.NewAMK1_SetPoints1()
	Pack_SOC := CANBMScan.NewPack_SOC()
	
	rx := socketcan.NewReceiver(conn)
	for rx.Receive() {
		frame := rx.Frame()
		// fmt.Println(frame.String())

		// Check signal and unmarshal
		switch frame.ID {
		case CANBMScan.Messages().Pack_SOC.ID:
			if err := Pack_SOC.UnmarshalFrame(frame); err != nil {
				panic(err)
			}
			// UnmarshalFrame sets all signals under the message
			fmt.Printf("\t%s\n", "Pack_SOC")
			fmt.Printf("\t\t%.2f\n", Pack_SOC.Pack_SOC())
			fmt.Printf("\t\t%.2f\n", Pack_SOC.Maximum_Pack_Voltage())

		case CANAMKInvertercan.Messages().AMK1_SetPoints1.ID:
			if err := AMK1_SetPoints1.UnmarshalFrame(frame); err != nil {
				panic(err)
			}
			fmt.Printf("\t%s\n", "AMK1_SetPoints1")
			fmt.Printf("\t\t%d\n", AMK1_SetPoints1.AMK_TargetVelocity())
			fmt.Printf("\t\t%d\n", AMK1_SetPoints1.AMK_TorqueLimitPositiv())
			fmt.Printf("\t\t%d\n", AMK1_SetPoints1.AMK_TorqueLimitNegativ())
			fmt.Printf("\t\t%t\n", AMK1_SetPoints1.AMK_bErrorReset())
			fmt.Printf("\t\t%t\n", AMK1_SetPoints1.AMK_bEnable())
			fmt.Printf("\t\t%t\n", AMK1_SetPoints1.AMK_bDcOn())
			fmt.Printf("\t\t%t\n", AMK1_SetPoints1.AMK_bInverterOn())
		}
	}
	if rx.Err() != nil {
			panic(err)
	}
}