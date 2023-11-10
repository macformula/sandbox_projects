package main

import (
	"context"
)

type Tracer interface {
	StartTrace(ctx context.Context) error
	StopTrace()
}