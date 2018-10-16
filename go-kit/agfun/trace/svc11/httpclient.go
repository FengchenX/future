package svc11

import (
	"context"
	"fmt"
	"github.com/feng/future/go-kit/agfun/trace/middleware"
	opentracing "github.com/opentracing/opentracing-go"
	"io/ioutil"
	"net/http"
	"net/url"
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

	req = c.traceRequest(req.WithContext(ctx))
	resp, err := c.httpClient.Do(req)
	if err != nil {
		span.SetTag("error", err.Error())
		return "", err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		span.SetTag("error", err.Error())
		return "", err
	}
	return string(data), nil
}

func (c *client) Sum(ctx context.Context, a, b int64) (int64, error) {
	panic("todo")
}

func NewHTTPClient(tracer opentracing.Tracer, baseURL string) Service {
	return &client{
		baseURL:      baseURL,
		httpClient:   &http.Client{},
		tracer:       tracer,
		traceRequest: middleware.ToHTTPRequest(tracer),
	}
}
