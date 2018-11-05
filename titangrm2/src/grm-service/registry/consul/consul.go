package consul

import (
	"grm-service/registry"
)

func NewRegistry(opts ...registry.Option) registry.Registry {
	return registry.NewRegistry(opts...)
}
