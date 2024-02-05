package receiver

import (
	"context"
	"fmt"

	CANAMKInvertercan "github.com/macformula/sandbox_projects/output/CANAMKInvertercan"
	"go.einride.tech/can/pkg/socketcan"
)

func Receive(name string, done chan struct{}) {
	conn, err := socketcan.DialContext(context.Background(), "can", "can0")
	if err != nil {
		panic(err)
	}

	// Messages
	//AMK1_SetPoints1 := CANAMKInvertercan.NewAMK1_SetPoints1()
	//AMK1_ActualValues2 := CANAMKInvertercan.NewAMK1_ActualValues2()
	//Pack_SOC := CANBMScan.NewPack_SOC()
	//Pack_State := CANBMScan.NewPack_State()
	//Contactor_Feedback := CANBMScan.NewContactor_Feedback()

	i := CANAMKInvertercan.Messages()
	rx := socketcan.NewReceiver(conn)
	for rx.Receive() {
		frame := rx.Frame()
		// fmt.Println(frame.String())
		message, err := i.UnmarshalFrame(frame)
		if err != nil {
			return
		}

		// Check signal and unmarshal
		fmt.Println(message)
		fmt.Println("HERE")
		//switch frame.ID {
		//case CANBMScan.Messages().Pack_SOC.ID:
		//	if err := Pack_SOC.UnmarshalFrame(frame); err != nil {
		//		panic(err)
		//	}
		//	// UnmarshalFrame sets all signals under the message
		//	fmt.Printf("\t%s\n", "Pack_SOC")
		//	fmt.Printf("\t\t%.2f\n", Pack_SOC.Pack_SOC())
		//	fmt.Printf("\t\t%.2f\n", Pack_SOC.Maximum_Pack_Voltage())
		//
		//case CANAMKInvertercan.Messages().AMK1_SetPoints1.ID:
		//	if err := AMK1_SetPoints1.UnmarshalFrame(frame); err != nil {
		//		panic(err)
		//	}
		//	fmt.Printf("\t%s\n", "AMK1_SetPoints1")
		//	fmt.Printf("\t\t%d\n", AMK1_SetPoints1.AMK_TargetVelocity())
		//	fmt.Printf("\t\t%d\n", AMK1_SetPoints1.AMK_TorqueLimitPositiv())
		//	fmt.Printf("\t\t%d\n", AMK1_SetPoints1.AMK_TorqueLimitNegativ())
		//	fmt.Printf("\t\t%t\n", AMK1_SetPoints1.AMK_bErrorReset())
		//	fmt.Printf("\t\t%t\n", AMK1_SetPoints1.AMK_bEnable())
		//	fmt.Printf("\t\t%t\n", AMK1_SetPoints1.AMK_bDcOn())
		//	fmt.Printf("\t\t%t\n", AMK1_SetPoints1.AMK_bInverterOn())
		//case CANAMKInvertercan.Messages().AMK1_ActualValues2.ID:
		//	if err := AMK1_ActualValues2.UnmarshalFrame(frame); err != nil {
		//		panic(err)
		//	}
		//	fmt.Printf("\t%s\n", "AMK1_ActualValues2")
		//	fmt.Printf("\t\tAMK_TempMotor: %0.2f\n", AMK1_ActualValues2.AMK_TempMotor())
		//	fmt.Printf("\t\tAMK_TempInverter: %0.2f\n", AMK1_ActualValues2.AMK_TempInverter())
		//	fmt.Printf("\t\tAMK_TempIGBT: %0.2f\n", AMK1_ActualValues2.AMK_TempIGBT())
		//	fmt.Printf("\t\tAMK_ErrorInfo: %d\n", AMK1_ActualValues2.AMK_ErrorInfo())
		//case CANBMScan.Messages().Pack_State.ID:
		//	if err := Pack_State.UnmarshalFrame(frame); err != nil {
		//		panic(err)
		//	}
		//	fmt.Printf("\t%s\n", "Pack_State")
		//	fmt.Printf("\t\tPack_Current: %0.2f\n", Pack_State.Pack_Current())
		//	fmt.Printf("\t\tPack_Inst_Voltage: %0.2f\n", Pack_State.Pack_Inst_Voltage())
		//	fmt.Printf("\t\tAvg_Cell_Voltage: %0.2f\n", Pack_State.Avg_Cell_Voltage())
		//	fmt.Printf("\t\tPopulated_Cells: %d\n", Pack_State.Populated_Cells())
		//case CANBMScan.Messages().Contactor_Feedback.ID:
		//	if err := Contactor_Feedback.UnmarshalFrame(frame); err != nil {
		//		panic(err)
		//	}
		//	fmt.Printf("\t%s\n", "Contactor_Feedback")
		//	fmt.Printf("\t\tPack_Precharge_Feedback: %t\n", Contactor_Feedback.Pack_Precharge_Feedback())
		//	fmt.Printf("\t\tPack_Negative_Feedback: %t\n", Contactor_Feedback.Pack_Negative_Feedback())
		//	fmt.Printf("\t\tPack_Positive_Feedback: %t\n", Contactor_Feedback.Pack_Positive_Feedback())
		//default:
		//	fmt.Print(frame.ID)
		//}
	}
	if rx.Err() != nil {
		panic(err)
	}
	done <- struct{}{}
}
