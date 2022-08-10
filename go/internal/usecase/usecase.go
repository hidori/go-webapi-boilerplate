package usecase

import (
	"context"
)

// Usecase は、ユースケースのインターフェースです。
type Usecase[TInputPort any, TOutputPort any] interface {
	Execute(ctx context.Context, input TInputPort) (TOutputPort, error)
}
