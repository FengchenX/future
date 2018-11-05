package es

import (
	"encoding/json"
	"fmt"

	"data-manager/types"

	"github.com/olivere/elastic"

	"grm-service/dbcentral/es"
)

type MetaEs struct {
	es.DynamicES
}

// 更新数据快视图
func (e *MetaEs) UpdateDataSnap(id, path string) error {
	fmt.Println(id, path)
	e.Connect()
	_, err := e.Cli.Update().Index(es.EsIndex).Type(es.EsType).Id(id).
		Doc(map[string]interface{}{"snapshot": path}).
		Do(e.Ctx)
	if err != nil {
		return err
	}
	return nil
}

// 更新查看次数
func (e *MetaEs) UpdateDataViewCnt(id string, cnt int) error {
	fmt.Println(id, cnt)
	e.Connect()
	_, err := e.Cli.Update().Index(es.EsIndex).Type(es.EsType).Id(id).
		Doc(map[string]interface{}{"view_cnt": cnt}).
		Do(e.Ctx)
	if err != nil {
		return err
	}
	return nil
}

// 更新tags
func (e *MetaEs) UpdateDataTags(id, tags string) error {
	fmt.Println(id, tags)
	e.Connect()
	_, err := e.Cli.Update().Index(es.EsIndex).Type(es.EsType).Id(id).
		Doc(map[string]interface{}{"tags": tags}).
		Do(e.Ctx)
	if err != nil {
		return err
	}
	return nil
}

// 获取数据元数据信息
func (e MetaEs) GetDataMeta(dataId string) (map[string]interface{}, error) {
	e.Connect()
	var data map[string]interface{}
	idQuery := elastic.NewIdsQuery(es.EsType).Ids(dataId)
	searchResult, err := e.Cli.Search().Index(es.EsIndex).
		Query(idQuery).
		Do(e.Ctx)
	if err != nil {
		return nil, err
	}
	if searchResult.Hits.TotalHits > 0 {
		for _, hit := range searchResult.Hits.Hits {
			err := json.Unmarshal(*hit.Source, &data)
			if err != nil {
				return nil, err
			}
		}
	} else {
		return nil, nil
	}
	return data, nil
}

type updateMeta struct {
	Value      interface{} `json:"value"`
	IsModified bool        `json:"is_modified"`
}

type updateMetas struct {
	Metas map[string]updateMeta `json:"metas"`
}

// 更新数据元数据信息
func (e MetaEs) UpdateDataMeta(data string, meta types.UpdateDataMetaReq) error {
	e.Connect()
	doc := make(map[string]interface{})
	for _, group := range meta.Metas {
		var metas updateMetas
		metas.Metas = make(map[string]updateMeta)
		for _, meta := range group.Metas {
			metas.Metas[meta.Field] = updateMeta{Value: meta.Value, IsModified: true}
		}
		doc[group.Group] = metas
	}
	ret, _ := json.Marshal(doc)
	fmt.Println("doc:", string(ret))
	_, err := e.Cli.Update().Index(es.EsIndex).Type(es.EsType).Id(data).
		Doc(doc).
		Do(e.Ctx)
	if err != nil {
		return err
	}
	return nil
}
