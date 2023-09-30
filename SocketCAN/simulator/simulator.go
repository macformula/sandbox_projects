package simulator

import (
	"context"
	"os"
	"fmt"
	"bufio"
	"strconv"
	"strings"

	"go.einride.tech/can"
	"go.einride.tech/can/pkg/socketcan"
)

func parseCANFrameInfo(str string) (can.Frame, error) {
	var frame can.Frame
	var id uint32
	var length uint8

	parts := strings.Split(str, ",")
	if len(parts) != 3 {
		return frame, fmt.Errorf("invalid format")
	}

	// Parse CAN ID
	tempId, err := strconv.ParseUint(parts[0], 10, 32)
	if err != nil {
		return frame, fmt.Errorf("failed to parse CAN ID: %v", err)
	}
	id = uint32(tempId)

	// Parse length
	tempLength, err := strconv.ParseUint(parts[1], 10, 8)
	if err != nil {
		return frame, fmt.Errorf("failed to parse length: %v", err)
	}
	length = uint8(tempLength)
	bytes := make([]byte, length)

	// Parse data
	dataStr := strings.TrimSpace(parts[2])
	dataStr = strings.TrimPrefix(dataStr, "[")
	dataStr = strings.TrimSuffix(dataStr, "]")
	dataBytes := strings.Split(dataStr, " ")
	for i, byteStr := range dataBytes {
		if uint8(i) == length {break}

		byteVal, err := strconv.ParseUint(byteStr, 10, 8)
		if err != nil {
			return frame, fmt.Errorf("failed to parse data byte: %v", err)
		}
		bytes[i] = byte(byteVal)
	}

	data := can.Data{} // Initialize a can.Data array
	copy(data[:], bytes)

	return can.Frame{
		ID:     id,
		Length: length,
		Data:   data,
	}, nil
}


func Simulate() {
	filePath := "can.log" // Path to the log file

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	conn, err := socketcan.DialContext(context.Background(), "can", "can0")
	if err != nil {
		fmt.Printf("here1: %v\n", err)
		panic(err)
	}
	tx := socketcan.NewTransmitter(conn)

	for scanner.Scan() {
		line := scanner.Text()
		frame, err := parseCANFrameInfo(line)
		if err != nil {
			fmt.Printf("here2: %v\n", err)
			fmt.Println(frame)
			panic(err)
		}
	
		if err := tx.TransmitFrame(context.Background(), frame); err != nil {
			fmt.Printf("here3: %v\n", err)
			panic(err)
		}
	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading log file: %v\n", err)
	}
}
