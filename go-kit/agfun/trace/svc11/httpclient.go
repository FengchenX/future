package svc11

import (
	"fmt"
	"context"
	"net/http"
	opentracing "github.com/opentracing/opentracing-go"
	"net/url"
	"github.com/feng/future/go-kit/agfun/trace/middleware"
)

type client struct {
	baseURL      string
	httpClient   *http.Client
	tracer       opentracing.Tracer
	traceRequest middleware.RequestFunc
}

func (c *client) Concat(ctx context.Context, a, b string) (string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Concat")
	defer span.Finish()

	url := fmt.Sprintf(
		"%s/concat/?a=%s&b=%s", c.baseURL, url.QueryEscape(a), url.QueryEscape(b),
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	fmt.Println(req, err)
	panic("todo")
}