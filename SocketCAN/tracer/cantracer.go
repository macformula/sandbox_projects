package tracer

import (
	"context"
	"time"

	"go.uber.org/zap"
)

type Tracer struct {
	l            *zap.Logger
	canInterface string
	stop         chan struct{}
	samplePeriod time.Duration
	canData      []string
}

func NewTracer(samplePeriod time.Duration, l *zap.Logger) *Tracer {
	return &Tracer{
		l:            l.Named("can_tracer"),
		samplePeriod: samplePeriod,
	}
}

func (t *Tracer) StartTrace(ctx context.Context) error {
	// Setup code here.

	t.stop = make(chan struct{})
	go t.trace(ctx)

	return nil
}

func (t *Tracer) StopTrace() {
	// Send the stop signal
	t.l.Info("sending stop signal")

	close(t.stop)

	// Save to file logic????
	
}

func (t *Tracer) trace(ctx context.Context) {
	ticker := time.NewTicker(t.samplePeriod)
	defer ticker.Stop()

	i := 0 // just for printing, remove me
	for {
		select {
		case <-t.stop:
			t.l.Info("received stop command, exiting tracer")
			return
		case <-ctx.Done():
			t.l.Error("received context deadline", zap.Error(ctx.Err()))
			return
		case <-ticker.C:
		}
		// tracer logic here

		t.l.Info("Tracing", zap.Int("i", i))

		i += 1
	}
}
