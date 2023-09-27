package worker

import "context"

type Process interface {
	Produce(context.Context) error
	Consume(context.Context) (Data, error)
}
