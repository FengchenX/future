package es

import (
	"context"

	"github.com/olivere/elastic"

	"grm-service/log"
)

const (
	EsIndex = "datameta"
	EsType  = "type"
)

type DynamicES struct {
	Endpoints []string
	Cli       *elastic.Client
	Ctx       context.Context
}

func (e *DynamicES) Connect() error {
	if e.Cli != nil {
		return nil
	}

	client, err := elastic.NewSimpleClient(
		elastic.SetURL(e.Endpoints...),
		//设置健康检查
		//elastic.SetHealthcheck(true),
	)
	if err != nil {
		log.Error(err)
		return err
	}
	e.Cli = client
	e.Ctx = context.Background()
	return nil
}
