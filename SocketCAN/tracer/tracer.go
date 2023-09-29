package tracer

import (
	"context"
	"log"
	"os"

	"go.einride.tech/can/pkg/socketcan"
)

func Trace() {
	file, err := os.OpenFile("can.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file: ", err)
	}
	defer file.Close()

	log.SetOutput(file) // Set the log output to the file

	conn, err := socketcan.DialContext(context.Background(), "can", "can0")
	if err != nil {
			panic(err)
	}

	rx := socketcan.NewReceiver(conn)
	for rx.Receive() {
		frame := rx.Frame()

		log.Println(frame)
	}

	if rx.Err() != nil {
			panic(err)
	}
}

// func setupFile() {
// 	file, err := os.OpenFile("can.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
// 	if err != nil {
// 		log.Fatal("Failed to open log file: ", err)
// 	}
// 	defer file.Close()

// 	log.SetOutput(file) // Set the log output to the file
// }