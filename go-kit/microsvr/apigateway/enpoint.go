package apigateway

import (
	"github.com/go-kit/kit/endpoint"
)

func makeGetAccountEndpoint(s GatewayService) endpoint.Endpoint {
	return s.GetAccount()
} 