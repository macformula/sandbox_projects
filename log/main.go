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
	_sleepTime    = 5 * time.Second
)

func main() {
	fmt.Println("Starting program")
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


// func main2() {
// 	// The zap.Config struct includes an AtomicLevel. To use it, keep a
// // reference to the Config.
// // rawJSON := []byte(`{
// // 	"level": "debug",
// // 	"outputPaths": ["stdout"],
// // 	"errorOutputPaths": ["stderr"],
// // 	"encoding": "json",
// // 	"encoderConfig": {
// // 		"messageKey": "message",
// // 		"levelKey": "level",
// // 		"levelEncoder": "lowercase"
// // 	}
// // }`)

// rawJSON, err := os.ReadFile("config.json")
// if err != nil {
// 		panic(err)
// }
// var cfg zap.Config
// if err := json.Unmarshal(rawJSON, &cfg); err != nil {
// 	panic(err)
// }

// cfg.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
// cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
// cfg.EncoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
// cfg.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	
// 	// cfg.OutputPaths = []string{"app.log"}
// logger := zap.Must(cfg.Build())
// logger = logger.Named("can_tracer")
	
// defer logger.Sync()

// 	// Append custom fields to the log
// canData := map[string]int{
// 	"speed": 60,
// 	"fuel":  70,
// }
// logger.Info("info logging disabled", zap.Any("canData", canData))

// // logger.Info("info logging disabled", canData: "hi")

// cfg.Level.SetLevel(zap.DebugLevel)
// logger.Info("info logging enabled")

// 	cfg.Level.SetLevel(zap.ErrorLevel)
// logger.Info("info logging disabled")
// }
