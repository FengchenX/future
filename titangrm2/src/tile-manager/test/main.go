package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"log"
	"os"
	"reflect"
)

func init() {
	errlog := log.New(os.Stdout, "APP", log.LstdFlags)
	var err error
	client, err = elastic.NewClient(elastic.SetErrorLog(errlog), elastic.SetURL(host))
	if err != nil {
		panic(err)
	}
	info, code, err := client.Ping(host).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	esvesion, err := client.ElasticsearchVersion(host)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esvesion)
}

func main() {
	//create()
	query()
}

var client *elastic.Client
var host = "http://127.0.0.1:9200/"

type Employee struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Age       int      `json:"age"`
	About     string   `json:"about"`
	Interests []string `json:"interests"`
}

/*crud*/
//创建
func create() {
	e1 := Employee{
		FirstName: "Jane",
		LastName:  "Smith",
		Age:       32,
		About:     "I like to collect rock albums",
		Interests: []string{"music"},
	}
	put1, err := client.Index().Index("megacorp").
		Type("employee").
		Id("1").BodyJson(e1).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put1.Id, put1.Index, put1.Type)

	e2 := `{"first_name":"John","last_name":"Smith","age":25,"about":"I love to go rock climbing","interests":["sports","music"]}`
	put2, err := client.Index().
		Index("megacorp").
		Type("employee").
		Id("2").BodyJson(e2).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to indexs%s, type %s\n", put2.Id, put2.Index, put2.Type)

	e3 := `{"first_name":"Douglas","last_name":"Fir","age":35,"about":"I like to build cabinets","interests":["forestry"]}`
	put3, err := client.Index().
		Index("megacorp").
		Type("employee").
		Id("3").
		BodyJson(e3).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put3.Id, put3.Index, put3.Type)

}

//删除
func delete() {

	res, err := client.Delete().Index("megacorp").
		Type("employee").
		Id("1").
		Do(context.Background())
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("delete result %s\n", res.Result)
}

//修改
func update() {
	res, err := client.Update().
		Index("megacorp").
		Type("employee").
		Id("2").
		Doc(map[string]interface{}{"age": 88}).
		Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	fmt.Printf("update age %s\n", res.Result)

}

// 搜索
func query() {
	var res *elastic.SearchResult
	var err error
	// 取所有
	res, err = client.Search("megacorp").Type("employee").Do(context.Background())
	printEmployee(res, err)

	//字段相等
	//q := elastic.NewQueryStringQuery("last_name:Smith")
	//res, err = client.Search("megacorp")

}

func printEmployee(res *elastic.SearchResult, err error) {
	if err != nil {
		print(err.Error())
		return
	}
	var typ Employee
	for _, item := range res.Each(reflect.TypeOf(typ)) {
		t := item.(Employee)
		fmt.Printf("%#v\n", t)
	}
}
