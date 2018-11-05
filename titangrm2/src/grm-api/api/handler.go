package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

func AllowOrigin(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers,auth-session,Accept, Content-Type, Content-Length, Accept-Encoding,X-CSRF-Token,Authorization,X-Requested-With")
}

func AllowSecurity(w http.ResponseWriter) {
	//	w.Header().Set("Content-Security-Policy", "")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-XSS-Protection", "1; mode=block")
}

func Index(w http.ResponseWriter, r *http.Request) {
	AllowOrigin(w)
	w.WriteHeader(200)
}

func MakeAPIHandler(ctx context.Context, endpoint endpoint.Endpoint, logger log.Logger,
	svcName string, ispublic bool,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		done := true
		defer func(begin time.Time) {
			go func(begin time.Time) {
				if !done {
					return
				}
				//日志
				logger.Log(
					"path", r.RequestURI,
					"method", r.Method,
					"took", time.Since(begin),
				)

				//r.Body.Close()
			}(begin)
		}(time.Now())

		defer r.Body.Close()
		fmt.Println("[info] RemoteAddr:", r.RemoteAddr)

		// 跨域问题
		AllowOrigin(w)
		AllowSecurity(w)
		r.Body = http.MaxBytesReader(w, r.Body, 100<<30)
		w.Header().Set("Content-Type", "application/json")

		resp, err := endpoint(ctx, r)
		if err != nil {
			logger.Log("method", "endpoint", "error", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// 这里将json形式的回复信息进行转化
		b, ok := resp.([]byte)
		if !ok {
			//logger.Log("error", base.ErrNotBytes.Error())
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		// 这里将回复信息写入
		_, err = w.Write(b)
		if err != nil {
			logger.Log("method", "Write", "error", err.Error())
			return
		}
	}
}
