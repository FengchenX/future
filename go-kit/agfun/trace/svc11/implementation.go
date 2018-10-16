package svc11

import (
	"context"
	opentracing "github.com/opentracing/opentracing-go"
)

type svc11 struct {
	//todo
}

func (s *svc11) Concat(ctx context.Context, a, b string) (string, error) {
	if len(a)+len(b) > StrMaxSize {
		span := opentracing.SpanFromContext(ctx)
		span.SetTag("error", ErrMaxSize.Error())
		return "", ErrMaxSize
	}
	return a + b, nil
}

func (s *svc11) Sum(ctx context.Context, a, b int64) (int64, error) {
	panic("todo")
}

func NewService() Service {
	return &svc11{
		
	}
}
