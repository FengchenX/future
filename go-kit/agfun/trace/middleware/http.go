package middleware

import (
	"fmt"
	"net"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"

	"github.com/openzipkin-contrib/zipkin-go-opentracing/thrift/gen-go/zipkincore"
)

type RequestFunc func(req *http.Request) *http.Request

func ToHTTPRequest(tracer opentracing.Tracer) RequestFunc {
	return func(req *http.Request) *http.Request {
		// Retrieve the Span from context.
		if span := opentracing.SpanFromContext(req.Context()); span != nil {

			// We are going to use this span in a client request, so mark as such.
			ext.SpanKindRPCClient.Set(span)

			// Add some standard OpenTracing tags, useful in an HTTP request.
			ext.HTTPMethod.Set(span, req.Method)
			span.SetTag(zipkincore.HTTP_HOST, req.URL.Host)
			span.SetTag(zipkincore.HTTP_PATH, req.URL.Path)
			ext.HTTPUrl.Set(
				span,
				fmt.Sprintf("%s://%s%s", req.URL.Scheme, req.URL.Host, req.URL.Path),
			)

			// Add information on the peer service we're about to contact.
			if host, portString, err := net.SplitHostPort(req.URL.Host); err == nil {
				ext.PeerHostname.Set(span, host)
				if port, err := strconv.Atoi(portString); err != nil {
					ext.PeerPort.Set(span, uint16(port))
				}
			} else {
				ext.PeerHostname.Set(span, req.URL.Host)
			}

			// Inject the Span context into the outgoing HTTP Request.
			if err := tracer.Inject(
				span.Context(),
				opentracing.TextMap,
				opentracing.HTTPHeadersCarrier(req.Header),
			); err != nil {
				fmt.Printf("error encountered while trying to inject span: %+v\n", err)
			}
		}
		return req
	}
}

type HandlerFunc func(next gin.HandlerFunc) gin.HandlerFunc

func FromHTTPRequest(tracer opentracing.Tracer, operationName string,
) HandlerFunc {
	return func(next gin.HandlerFunc) gin.HandlerFunc {
		return func(c *gin.Context) {
			wireContext, err := tracer.Extract(
				opentracing.TextMap,
				opentracing.HTTPHeadersCarrier(c.Request.Header),
			)
			if err != nil {
				fmt.Sprintf("error encountered while trying to extract span: %+v\n", err)
			}
			span := tracer.StartSpan(operationName, ext.RPCServerOption(wireContext))
			defer span.Finish()
			ctx := opentracing.ContextWithSpan(c.Request.Context(), span)
			c.Request = c.Request.WithContext(ctx)
		}
	}
}