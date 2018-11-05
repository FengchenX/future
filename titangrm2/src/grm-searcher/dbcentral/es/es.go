package es

import (
	"context"
	//	"encoding/json"
	"fmt"
	"github.com/olivere/elastic"
	"grm-searcher/types"
	//	"grm-service/log"
	"reflect"
	"strconv"
	"strings"
)

type ESUtil struct {
	Client *elastic.Client
}

const META_TYPE = "type"

const mapping = `
{
	"mappings": {
        "doc": {
            "properties": {
                "geometry": {
                    "type": "geo_shape",
                    "tree": "quadtree",
                    "precision": "2m"
                }
            }
        }
    }
}`

func NewClient(url string) (*elastic.Client, error) {
	client, err := elastic.NewSimpleClient(
		elastic.SetURL(url))
	return client, err
}

func (es *ESUtil) CreateIndex(index, _mapping string) error {
	_, err := es.Client.CreateIndex(index).BodyString(_mapping).Do(context.Background())
	return err
}

func (es *ESUtil) QueryPage(query elastic.Query, index, order, limit, sort, offset string) *elastic.SearchService {
	_search := es.Client.Search().Index(index).Query(query)
	if len(sort) > 0 {
		if order == "desc" {
			_search.Sort(sort, false)
		} else {
			_search.Sort(sort, true)
		}
		fmt.Println("sort,order:", sort, order)
	}

	if len(offset) > 0 && len(limit) > 0 {
		_offset, err := strconv.Atoi(offset)
		if err != nil {
			_offset = 0
		}
		_limit, err := strconv.Atoi(limit)
		if err != nil {
			_limit = 10
		}
		_search.From(_offset).Size(_limit)
	}

	return _search
}

func (es *ESUtil) ScrollPage(query elastic.Query, index, order, limit, sort, scrollid string) {
	_search := es.Client.Scroll().Index().Query(query)
	if len(sort) > 0 {
		if order == "desc" {
			_search.Sort(sort, false)
		} else {
			_search.Sort(sort, true)
		}
	}
	if len(scrollid) > 0 && len(limit) > 0 {
		_limit, err := strconv.Atoi(limit)
		if err != nil {
			_limit = 10
		}
		_search.ScrollId(scrollid)
		_search.Size(_limit)
	}
}

func (es *ESUtil) QueryById(index, dataid string) (types.TypeMeta, error) {
	var datameta types.TypeMeta
	idQuery := elastic.NewIdsQuery(META_TYPE).Ids(dataid)

	_search := es.Client.Search().Index(index).Query(idQuery)
	searchResult, err := _search.Do(context.Background())
	if err != nil {
		return datameta, err
	}

	for _, item := range searchResult.Each(reflect.TypeOf(datameta)) {
		if t, ok := item.(types.TypeMeta); ok {
			return t, nil
		}
	}
	return datameta, nil
}

func (es *ESUtil) TypeQuery(index, typeName string, dataids []string, qi types.SearchInfo) (int64, []*types.MetaInfo, error) {
	var datametas []*types.MetaInfo = make([]*types.MetaInfo, 0)

	var querys []elastic.Query = make([]elastic.Query, 0)
	if len(qi.Key) > 0 {
		stringQuery := elastic.NewQueryStringQuery(strings.ToLower(qi.Key))
		querys = append(querys, stringQuery)
	}

	if len(querys) == 0 {
		querys = append(querys, elastic.NewMatchAllQuery())
	}

	boolQuery := elastic.NewBoolQuery().Must(querys...)
	if qi.Geometry != nil {
		//es的空间查询
		//geojson to geometry
		geoQuery := elastic.NewGeoShapeQuery("geometry").Shape(*qi.Geometry).Relation("intersects")
		boolQuery.Filter(geoQuery)
	}

	if len(dataids) > 0 {
		idsqry := elastic.NewIdsQuery("type").Ids(dataids...)
		boolQuery.Filter(idsqry)
	}

	if len(typeName) > 0 {
		stringQuery := elastic.NewTermQuery("data_type", strings.ToLower(typeName))
		boolQuery.Filter(stringQuery)

		statusQuery := elastic.NewTermQuery("status", "normal")
		boolQuery.Filter(statusQuery)
	}

	//TODO
	for _, f := range qi.Attrs {
		if f.IsTime == true {
			//时间
			rangeQuery := elastic.NewRangeQuery(f.Name).Gte(f.MinValue).Lte(f.MaxValue).Format("yyyy-MM-dd")
			boolQuery.Filter(rangeQuery)
		} else if len(f.MinValue) > 0 && len(f.MaxValue) > 0 {
			//范围
			rangeQuery := elastic.NewRangeQuery(f.Name).Gte(f.MinValue).Lte(f.MaxValue)
			boolQuery.Filter(rangeQuery)
		} else if f.Operand == "in" {
			//in
			vs := f.Value.([]interface{})
			inqry := elastic.NewTermsQuery(f.Name, vs...)
			boolQuery.Filter(inqry)
		} else if f.Operand == "=" {
			//==
			eqqry := elastic.NewTermQuery(f.Name, f.Value)
			boolQuery.Filter(eqqry)
		} else if f.Operand == "like" {
			//like
			wldqry := elastic.NewWildcardQuery(f.Name, f.Value.(string))
			boolQuery.Filter(wldqry)
		}
	}

	querySearch := es.QueryPage(boolQuery, index, qi.Order, qi.Limit, qi.Sort, qi.Offset)
	if qi.Only_geom {
		//只返回geometry
		querySearch.FetchSourceContext(elastic.NewFetchSourceContext(true).Include("geometry"))
	}

	searchResult, err := querySearch.Do(context.Background())
	if err != nil {
		return 0, datametas, err
	}

	var datameta types.TypeMeta
	for _, item := range searchResult.Each(reflect.TypeOf(datameta)) {
		if t, ok := item.(types.TypeMeta); ok {
			datametas = append(datametas, es.TypeMeta2MetaInfo(t))
		}
	}
	return searchResult.Hits.TotalHits, datametas, nil
}

func (es *ESUtil) QueryByDataset(dataids []string, typeName, key, index, order, limit, sort, offset string) (int64, []*types.MetaInfo, error) {
	var datametas []*types.MetaInfo = make([]*types.MetaInfo, 0)

	var querys []elastic.Query = make([]elastic.Query, 0)
	if len(key) > 0 {
		stringQuery := elastic.NewQueryStringQuery(strings.ToLower(key))
		querys = append(querys, stringQuery)
	}
	if len(querys) == 0 {
		querys = append(querys, elastic.NewMatchAllQuery())
	}

	boolQuery := elastic.NewBoolQuery().Must(querys...)
	statusQuery := elastic.NewTermQuery("status", "normal")
	boolQuery.Filter(statusQuery)

	if len(dataids) > 0 {
		idsqry := elastic.NewIdsQuery("type").Ids(dataids...)
		boolQuery.Filter(idsqry)
	}

	if len(typeName) > 0 {
		stringQuery := elastic.NewTermQuery("data_type", strings.ToLower(typeName))
		boolQuery.Filter(stringQuery)
	}

	querySearch := es.QueryPage(boolQuery, index, order, limit, sort, offset)
	searchResult, err := querySearch.Do(context.Background())
	if err != nil {
		return 0, datametas, err
	}

	var datameta types.TypeMeta
	for _, item := range searchResult.Each(reflect.TypeOf(datameta)) {
		if t, ok := item.(types.TypeMeta); ok {
			datametas = append(datametas, es.TypeMeta2MetaInfo(t))
		}
	}
	return searchResult.Hits.TotalHits, datametas, nil
}

func (es *ESUtil) QueryByKey(index, key, order, limit, sort, offset string) (int64, []*types.MetaInfo, error) {
	var datametas []*types.MetaInfo = make([]*types.MetaInfo, 0)

	// Search with a term query
	stringQuery := elastic.NewQueryStringQuery(key)
	_search := es.QueryPage(stringQuery, index, order, limit, sort, offset)

	searchResult, err := _search.Do(context.Background())
	if err != nil {
		return 0, datametas, err
	}

	var datameta types.TypeMeta
	for _, item := range searchResult.Each(reflect.TypeOf(datameta)) {
		if t, ok := item.(types.TypeMeta); ok {
			datametas = append(datametas, es.TypeMeta2MetaInfo(t))
		}
	}
	return searchResult.Hits.TotalHits, datametas, err
}

func (es *ESUtil) TypeMeta2MetaInfo(datameta types.TypeMeta) *types.MetaInfo {
	var mi types.MetaInfo = types.MetaInfo{}

	mi.LoadUser = datameta.User
	mi.DataType = datameta.DataType
	mi.Status = datameta.Status
	mi.SubType = datameta.SubType
	mi.TypeLabel = datameta.Label
	mi.UUID = datameta.DataId
	mi.EnvelopeGeoJson = datameta.Geometry

	mi.SnapPath = datameta.SnapPath
	mi.ThumbPath = datameta.ThumbPath
	mi.Tags = datameta.Tags
	mi.ViewCount = datameta.ViewCount

	mi.CreateTime = datameta.CreateTime
	mi.Detail = datameta.Detail
	mi.DataUrl = datameta.DataUrl

	if datameta.BasicInformation != nil {
		if datameta.BasicInformation.Metas["data_time"].Value != nil {
			mi.DataTime = datameta.BasicInformation.Metas["data_time"].Value.(string)
		}
		if datameta.BasicInformation.Metas["file_size"].Value != nil {
			mi.FileSize = datameta.BasicInformation.Metas["file_size"].Value.(float64)
		}
		if datameta.BasicInformation.Metas["modify_time"].Value != nil {
			mi.ModifyTime = datameta.BasicInformation.Metas["modify_time"].Value.(string)
		}
		if datameta.BasicInformation.Metas["name"].Value != nil {
			mi.Name = datameta.BasicInformation.Metas["name"].Value.(string)
		}
		if datameta.BasicInformation.Metas["path"].Value != nil {
			mi.Path = datameta.BasicInformation.Metas["path"].Value.(string)
		}
		if datameta.BasicInformation.Metas["shp_type"].Value != nil {
			mi.ShpType = datameta.BasicInformation.Metas["shp_type"].Value.(string)
		}
		mi.CentryPoint = datameta.BasicInformation.Metas["centry_point"].Value
	}

	if datameta.Content != nil && datameta.Content.Metas != nil {
		mi.Band = datameta.Content.Metas["band"].Value.(string)
	}
	return &mi
}
