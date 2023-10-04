package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"go.uber.org/zap"

	"can/logger/cantracer"
	"time"

	"go.uber.org/zap/zapcore"
)

const (
	_tracerPeriod = 1 * time.Millisecond
	_sleepTime    = 3 * time.Second
)

func main() {
	fmt.Print("starting program")
	rawJSON, err := os.ReadFile("config.json")
	if err != nil {
			panic(err)
	}
	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}

  cfg.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	cfg.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder
    
    // cfg.OutputPaths = []string{"app.log"}
	logger := zap.Must(cfg.Build())
  // logger = logger.Named("can_tracer")

	t := cantracer.NewTracer(_tracerPeriod, logger, "can0")

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	// defer cancel()
	ctx := context.Background()

	// Start tracing
	err = t.StartTrace(ctx)
	if err != nil {
		panic(err)
	}

	// Do other stuff here
	time.Sleep(_sleepTime)

	// Stop tracing
	t.StopTrace()

	// Add delay to see all logs
	time.Sleep(1 * time.Second)
}
