package api

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	kithttp "github.com/go-kit/kit/transport/http"

	grm "grm-service/service"
)

type Service struct {
	Path     string
	Method   string
	Factory  sd.Factory
	IsPublic bool
}

func MakeAPIRouter(ctx context.Context) map[string][]Service {
	services := map[string][]Service{
		grm.DataManagerService: {
			{Path: "/datas/{data-id}", Method: "DELETE", IsPublic: true, Factory: httpFactory(ctx, "DELETE", "/datas/{data-id}")},
			{Path: "/datas/{data-id}/comments", Method: "GET", IsPublic: true, Factory: httpFactory(ctx, "GET", "/datas/{data-id}/comments")},
			{Path: "/datas/{data-id}/comments", Method: "POST", IsPublic: true, Factory: httpFactory(ctx, "POST", "/datas/{data-id}/comments")},
			{Path: "/datas/{data-id}/comments/{comment-id}", Method: "DELETE", IsPublic: true, Factory: httpFactory(ctx, "DELETE", "/datas/{data-id}/comments/{comment-id}")},
			{Path: "/datas/{data-id}/content", Method: "GET", IsPublic: true, Factory: httpFactory(ctx, "GET", "/datas/{data-id}/content")},
			{Path: "/datas/{data-id}/info", Method: "GET", IsPublic: true, Factory: httpFactory(ctx, "GET", "/datas/{data-id}/info")},
			{Path: "/datas/{data-id}/info", Method: "PUT", IsPublic: true, Factory: httpFactory(ctx, "PUT", "/datas/{data-id}/info")},
			{Path: "/datas/{data-id}/metas", Method: "GET", IsPublic: true, Factory: httpFactory(ctx, "GET", "/datas/{data-id}/metas")},
			{Path: "/datas/{data-id}/metas", Method: "PUT", IsPublic: true, Factory: httpFactory(ctx, "PUT", "/datas/{data-id}/metas")},
			{Path: "/datas/{data-id}/snapshot", Method: "POST", IsPublic: true, Factory: httpFactory(ctx, "POST", "/datas/{data-id}/snapshot")},

			{Path: "/datasets", Method: "GET", IsPublic: true, Factory: httpFactory(ctx, "GET", "/datasets")},
			{Path: "/datasets", Method: "POST", IsPublic: true, Factory: httpFactory(ctx, "POST", "/datasets")},
			{Path: "/datasets/{dataset-id}", Method: "PUT", IsPublic: true, Factory: httpFactory(ctx, "PUT", "/datasets/{dataset-id}")},
			{Path: "/datasets/{dataset-id}", Method: "DELETE", IsPublic: true, Factory: httpFactory(ctx, "DELETE", "/datasets/{dataset-id}")},

			{Path: "/datasets/{dataset-id}/data", Method: "POST", IsPublic: true, Factory: httpFactory(ctx, "POST", "/datasets/{dataset-id}/data")},
			{Path: "/datasets/{dataset-id}/data", Method: "PUT", IsPublic: true, Factory: httpFactory(ctx, "PUT", "/datasets/{dataset-id}/data")},

			{Path: "/datasets/{dataset-id}/datas", Method: "DELETE", IsPublic: true, Factory: httpFactory(ctx, "DELETE", "/datasets/{dataset-id}/datas")},
			{Path: "/datasets/{dataset-id}/datas/{src-dataset}", Method: "PUT", IsPublic: true, Factory: httpFactory(ctx, "PUT", "/datasets/{dataset-id}/datas/{src-dataset}")},

			{Path: "/explorer", Method: "GET", IsPublic: true, Factory: httpFactory(ctx, "GET", "/explorer")},
			{Path: "/explorer", Method: "PUT", IsPublic: true, Factory: httpFactory(ctx, "PUT", "/explorer")},

			{Path: "/layers/{data-id}", Method: "GET", IsPublic: true, Factory: httpFactory(ctx, "GET", "/layers/{data-id}")},
			{Path: "/layers/{data-id}", Method: "POST", IsPublic: true, Factory: httpFactory(ctx, "POST", "/layers/{data-id}")},
			{Path: "/layers/{data-id}/{layer-id}", Method: "GET", IsPublic: true, Factory: httpFactory(ctx, "GET", "/layers/{data-id}/{layer-id}")},
			{Path: "/layers/{data-id}/{layer-id}", Method: "PUT", IsPublic: true, Factory: httpFactory(ctx, "PUT", "/layers/{data-id}/{layer-id}")},
			{Path: "/layers/{data-id}/{layer-id}", Method: "DELETE", IsPublic: true, Factory: httpFactory(ctx, "DELETE", "/layers/{data-id}/{layer-id}")},
			{Path: "/layers/{data-id}/{layer-id}/snapshot", Method: "POST", IsPublic: true, Factory: httpFactory(ctx, "POST", "/layers/{data-id}/{layer-id}/snapshot")},

			{Path: "/styles", Method: "GET", IsPublic: true, Factory: httpFactory(ctx, "GET", "/styles")},
			{Path: "/styles", Method: "POST", IsPublic: true, Factory: httpFactory(ctx, "POST", "/styles")},
			{Path: "/styles/{style-id}", Method: "GET", IsPublic: true, Factory: httpFactory(ctx, "GET", "/styles/{style-id}")},
			{Path: "/styles/{style-id}", Method: "PUT", IsPublic: true, Factory: httpFactory(ctx, "PUT", "/styles/{style-id}")},
			{Path: "/styles/{style-id}", Method: "DELETE", IsPublic: true, Factory: httpFactory(ctx, "DELETE", "/styles/{style-id}")},

			{Path: "/types", Method: "GET", IsPublic: true, Factory: httpFactory(ctx, "GET", "/types")},
			{Path: "/types/meta/{type-name}", Method: "GET", IsPublic: true, Factory: httpFactory(ctx, "GET", "/types/meta/{type-name}")},
			{Path: "/types/meta/{type-name}", Method: "PUT", IsPublic: true, Factory: httpFactory(ctx, "PUT", "/types/meta/{type-name}")},
			{Path: "/types/meta/{type-name}", Method: "POST", IsPublic: true, Factory: httpFactory(ctx, "POST", "/types/meta/{type-name}")},
			{Path: "/types/meta/{type-name}/{group}/{field}", Method: "DELETE", IsPublic: true, Factory: httpFactory(ctx, "DELETE", "/types/meta/{type-name}/{group}/{field}")},
			{Path: "/types/{type-name}", Method: "PUT", IsPublic: true, Factory: httpFactory(ctx, "PUT", "/types/{type-name}")},
		},
		grm.TitanAuthService: {
			{Path: "/captcha", Method: "POST", IsPublic: true, Factory: httpFactory(ctx, "POST", "/captcha")},
			{Path: "/login", Method: "PUT", IsPublic: true, Factory: httpFactory(ctx, "PUT", "/login")},
			{Path: "/logout", Method: "DELETE", IsPublic: true, Factory: httpFactory(ctx, "DELETE", "/logout")},
			{Path: "/users", Method: "POST", IsPublic: true, Factory: httpFactory(ctx, "POST", "/users")},
			{Path: "/users", Method: "GET", IsPublic: true, Factory: httpFactory(ctx, "GET", "/users")},
			{Path: "/users/{user-id}/status", Method: "PUT", IsPublic: true, Factory: httpFactory(ctx, "PUT", "/users/{user-id}/status")},

			{Path: "/groups", Method: "GET", IsPublic: true, Factory: httpFactory(ctx, "GET", "/groups")},
			{Path: "/groups", Method: "POST", IsPublic: true, Factory: httpFactory(ctx, "POST", "/groups")},
			{Path: "/groups/{group-id}", Method: "PUT", IsPublic: true, Factory: httpFactory(ctx, "PUT", "/groups/{group-id}")},
			{Path: "/groups/{group-id}", Method: "DELETE", IsPublic: true, Factory: httpFactory(ctx, "DELETE", "/groups/{group-id}")},
			{Path: "/groups/{group-id}/users", Method: "GET", IsPublic: true, Factory: httpFactory(ctx, "GET", "/groups/{group-id}/users")},
			{Path: "/groups/{group-id}/users", Method: "POST", IsPublic: true, Factory: httpFactory(ctx, "POST", "/groups/{group-id}/users")},
			{Path: "/groups/{group-id}/users/{user-id}", Method: "DELETE", IsPublic: true, Factory: httpFactory(ctx, "DELETE", "/groups/{group-id}/users/{user-id}")},
		},
		grm.DataImporterService: {
			{Path: "/scan", Method: "POST", IsPublic: true, Factory: httpFactory(ctx, "POST", "/scan")},
			{Path: "/scan/{task-id}", Method: "GET", IsPublic: true, Factory: httpFactory(ctx, "GET", "/scan/{task-id}")},
			{Path: "/load", Method: "POST", IsPublic: true, Factory: httpFactory(ctx, "POST", "/load")},
			{Path: "/sysdomain", Method: "POST", IsPublic: true, Factory: httpFactory(ctx, "POST", "/sysdomain")},

			{Path: "/tasks/{task-id}", Method: "DELETE", IsPublic: true, Factory: httpFactory(ctx, "DELETE", "/tasks/{task-id}")},
			{Path: "/tasks/{task-id}/logs/{log-type}", Method: "GET", IsPublic: true, Factory: httpFactory(ctx, "GET", "/tasks/{task-id}/logs/{log-type}")},
			{Path: "/tasks/{task-id}/status", Method: "PUT", IsPublic: true, Factory: httpFactory(ctx, "PUT", "/tasks/{task-id}/status")},
			{Path: "/tasks/{task-type}", Method: "GET", IsPublic: true, Factory: httpFactory(ctx, "GET", "/tasks/{task-type}")},
		},
		grm.StorageManagerService: {
			{Path: "/devices", Method: "POST", IsPublic: true, Factory: httpFactory(ctx, "POST", "/devices")},
			{Path: "/devices", Method: "GET", IsPublic: true, Factory: httpFactory(ctx, "GET", "/devices")},
			{Path: "/devices/{device-id}", Method: "PUT", IsPublic: true, Factory: httpFactory(ctx, "PUT", "/devices/{device-id}")},
			{Path: "/devices/{device-id}", Method: "DELETE", IsPublic: true, Factory: httpFactory(ctx, "DELETE", "/devices/{device-id}")},
		},
	}
	return services
}

func httpFactory(ctx context.Context, method, path string) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		var e endpoint.Endpoint
		if !strings.HasPrefix(instance, "http") {
			instance = "http://" + instance
		}
		u, err := url.Parse(instance)
		if err != nil {
			return nil, nil, err
		}
		u.Path = path
		e = kithttp.NewClient(method, u, passVarEncode, passDecode).Endpoint()
		return e, nil, nil
	}
}

func passVarEncode(_ context.Context, r *http.Request, request interface{}) error {
	req := request.(*http.Request)

	r.Body = req.Body
	r.Header = req.Header
	r.ContentLength = req.ContentLength
	r.MultipartForm = req.MultipartForm

	firstIndex := strings.Index(r.URL.Path, "{")
	if firstIndex >= 0 {
		//path中包含{}
		firstStr := r.URL.Path[:firstIndex]
		needIndex := strings.Index(req.URL.Path, firstStr)
		needStr := req.URL.Path[needIndex:]
		r.URL.Path = needStr
	}
	r.URL.RawQuery = req.URL.RawQuery
	return nil
}

func passDecode(_ context.Context, r *http.Response) (interface{}, error) {
	return ioutil.ReadAll(r.Body)
}
