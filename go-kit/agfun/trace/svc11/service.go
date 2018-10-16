package svc11

import (
	"context"
	"errors"
)

const (
	StrMaxSize = 1024
)

var (
	ErrMaxSize = errors.New("maximum size of 1024 bytes exceeded")
)

type Service interface {
	Concat(ctx context.Context, a, b string) (string, error)
	Sum(ctx context.Context, a, b int64) (int64, error)
}