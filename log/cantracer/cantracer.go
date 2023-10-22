package cantracer

import (
	"context"
	"fmt"
	"time"
	"strings"
	"strconv"

	"go.einride.tech/can/pkg/candevice"
	"go.einride.tech/can/pkg/socketcan"
	"go.uber.org/zap"
)

type CanData struct {
	rx *socketcan.Receiver
}

func NewCanData(canInterface string) (*CanData, error) {
	setup(canInterface)
	conn, err := socketcan.DialContext(context.Background(), "can", canInterface)
	if err != nil {
		return &CanData{}, err
	}

	return &CanData{
		rx: socketcan.NewReceiver(conn),
	}, nil
}

func setup(canInterface string) {
	fmt.Println("Setting up " + canInterface)
	d, _ := candevice.New(canInterface)
	_ = d.SetBitrate(250000)
	_ = d.SetUp()
	defer d.SetDown()
	fmt.Println("Done " + canInterface + " setup")
}

type Tracer struct {
	l            *zap.Logger
	canInterface string
	stop         chan struct{}
	samplePeriod time.Duration
	can          *CanData
}

func NewTracer(samplePeriod time.Duration, l *zap.Logger, canInterface string) *Tracer {
	return &Tracer{
		l:            l.Named("can_tracer"),
		samplePeriod: samplePeriod,
		canInterface: canInterface,
	}
}

func (t *Tracer) StartTrace(ctx context.Context) error {
	// Setup code here.
	t.can, _ = NewCanData(t.canInterface)
	t.l.Info("received start command, beginning trace")

	t.stop = make(chan struct{})
	go t.trace(ctx)

	return nil
}

func (t *Tracer) StopTrace() {
	// Send the stop signal
	defer t.l.Sync() // flushes buffer, if any
	t.l.Info("sending stop signal")

	close(t.stop)
}

func (t *Tracer) trace(ctx context.Context) {
	// ticker := time.NewTicker(t.samplePeriod)
	// defer ticker.Stop()

	i := 1 // just for printing, remove me
	for t.can.rx.Receive() {
		frame := t.can.rx.Frame()

		select {
		case <-t.stop:
			t.l.Info("received stop command, exiting tracer")
			return
		case <-ctx.Done():
			t.l.Error("received context deadline", zap.Error(ctx.Err()))
			return
		// case <-ticker.C:
		default:
		}
		// tracer logic here

		// t.l.Info("Tracing", zap.Int("i", i))

		// Create a builder for efficient string concatenation
    var builder strings.Builder

		// add time
		builder.WriteString(time.Now().Format("15:04:05.0000"))

		// add can
		builder.WriteString(" "+strings.TrimPrefix(t.canInterface, "can"))

		// add frame id
		builder.WriteString(" "+strconv.FormatUint(uint64(frame.ID), 10))

		// add Rx
		builder.WriteString(" Rx")

		// add byte length
		builder.WriteString(" "+strconv.FormatUint(uint64(frame.Length), 10))

		// add frame data
    for _, v := range frame.Data {
			builder.WriteString(" "+strconv.FormatUint(uint64(v), 16))
    }

    result := builder.String()
		// fmt.Println(result)
		t.l.Info(result)

		i += 1
	}
}
