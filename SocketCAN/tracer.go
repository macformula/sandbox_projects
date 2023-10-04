package main

type Tracer interface {
	StartTrace(ctx context.Context) error
	StopTrace()
}