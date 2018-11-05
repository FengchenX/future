package util

import (
	"fmt"
	"math/rand"
	"net/http"
	"reflect"
	"regexp"
	"strings"

	"grm-service/common"

	"github.com/emicklei/go-restful"
	"github.com/pborman/uuid"

	errors "grm-service/errors"
	log "grm-service/log"
)

type nullRet struct {
	ret string `json:"ret,omitempty"`
}

func isNil(i interface{}) bool {
	if i == nil {
		return true
	}

	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	log.Warn("Response type should be a pointer")
	return vi.IsValid()
}

//  http response
func ResWriteError(res *restful.Response, err error) error {
	log.Error(err)
	res.Header().Add("Content-Type", "application/json")
	return res.WriteError(http.StatusOK, errors.New(err.Error(), http.StatusInternalServerError))
}

func ResWriteHeaderEntity(res *restful.Response, value interface{}) error {
	res.Header().Add("Content-Type", "application/json")
	if isNil(value) {
		return res.WriteEntity(nullRet{})
	}
	return res.WriteHeaderAndEntity(http.StatusOK, value)
}

func ResWriteEntity(res *restful.Response, value interface{}) error {
	res.Header().Add("Content-Type", "application/json")
	if isNil(value) {
		return res.WriteEntity(nullRet{})
	}
	return res.WriteEntity(value)
}

// 解析分页参数
func ParserPageArgs(req *restful.Request) *common.PageFilter {
	args := common.PageFilter{
		Limit:  req.QueryParameter("limit"),
		Offset: req.QueryParameter("offset"),
		Order:  req.QueryParameter("order"),
		Sort:   req.QueryParameter("sort"),
	}
	return &args
}

// 拼接分页sql
func PageFilterSql(sql, keyCol string, page *common.PageFilter) string {
	if len(sql) == 0 || page == nil {
		return sql
	}

	if len(page.Sort) > 0 && len(page.Order) == 0 {
		page.Order = "desc"
	}

	if len(page.Sort) > 0 && page.Sort != keyCol {
		sql = fmt.Sprintf(`%s order by %s %s`, sql, page.Sort, page.Order)
		if len(keyCol) > 0 {
			sql = fmt.Sprintf(`%s,%s %s`, sql, keyCol, page.Order)
		}
	} else if len(keyCol) > 0 {
		sql = fmt.Sprintf(`%s order by %s %s`, sql, keyCol, page.Order)
	}

	if len(page.Offset) > 0 && len(page.Limit) > 0 {
		sql = fmt.Sprintf(`%s limit %s offset %s`, sql, page.Limit, page.Offset)
	}
	return sql
}

// 获取uuid
func NewUUID() string {
	uuid := uuid.NewUUID().String()
	return strings.Replace(uuid, "-", "", -1)
}

// 获取切片中随机元素
func GetRandomStr(list []string) string {
	return list[rand.Intn(len(list))]
}

func GetRandomItem(list []interface{}) interface{} {
	return list[rand.Intn(len(list))]
}

var re = regexp.MustCompile("^[0-9]*$")

func IsNum(num string) bool {
	num = strings.TrimSpace(" ")
	if num == "" {
		return false
	}
	return re.Match([]byte(num))
}
