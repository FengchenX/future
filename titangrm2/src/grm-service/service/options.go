package service

import (
	"context"
	"time"
)

type Options struct {
	Name      string
	Version   string
	Id        string
	Metadata  map[string]string
	Address   string
	Advertise string
	Namespace string

	RegistryAddr     string
	RegisterTTL      time.Duration
	RegisterInterval time.Duration

	// Alternative Options
	Context context.Context

	BeforeStart []func() error
	BeforeStop  []func() error
	AfterStart  []func() error
	AfterStop   []func() error
}

func newOptions(name, version string) Options {
	opt := Options{
		Name:             name,
		Version:          version,
		Id:               DefaultId,
		Address:          DefaultAddress,
		RegisterTTL:      DefaultRegisterTTL,
		RegisterInterval: DefaultRegisterInterval,
		Context:          context.TODO(),
	}
	return opt
}
