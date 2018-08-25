package transport

import (
	"context"
	"net/http"
	"encoding/json"
	"github.com/feng/future/go-kit/agfun/app-server/endpoint"
)

func decodeGetEthBalanceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoint.ReqGetEthBalance
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}