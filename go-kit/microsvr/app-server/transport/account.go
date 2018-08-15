package transport

import (
	"context"
	"net/http"
	"encoding/json"
	"github.com/feng/future/go-kit/microsvr/app-server/endpoint"
)

func decodeGetAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoint.ReqGetAccount
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}