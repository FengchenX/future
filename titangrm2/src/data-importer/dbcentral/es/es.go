package es

import (
	"grm-service/dbcentral/es"
)

type MetaEs struct {
	es.DynamicES
}

// 更新数据data url
func (e MetaEs) UpdateDataUrl(data, url string) error {
	e.Connect()
	_, err := e.Cli.Update().Index(es.EsIndex).Type(es.EsType).Id(data).
		Doc(map[string]interface{}{"data_url": url}).
		Do(e.Ctx)
	if err != nil {
		return err
	}
	return nil
}
