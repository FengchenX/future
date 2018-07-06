package bridge

import (
	"fmt"
	"net/http"
)

type Request interface {
	HttpRequest() (*http.Request, error)
}

type Client struct {
	Client *http.Client
}

func (c *Client) Query(req Request) (resp *http.Response, err error) {
	httpreq, _ := req.HttpRequest()
	resp, err = c.Client.Do(httpreq)
	return
}

type CdnRequest struct {
}

func (cdn *CdnRequest) HttpRequest() (*http.Request, error) {
	return http.NewRequest("GET", "/cdn", nil)
}

type LiveRequest struct {
}

func (cdn *LiveRequest) HttpRequest() (*http.Request, error) {
	return http.NewRequest("GET", "/live", nil)
}

func TestBridge() {
	client := &Client{http.DefaultClient}

	cdnReq := &CdnRequest{}
	fmt.Println(client.Query(cdnReq))

	liveReq := &LiveRequest{}
	fmt.Println(client.Query(liveReq))
}

type HandsetSoft interface {
	Run()
}

type HandsetGame struct {
}

func (hg HandsetGame) Run() {
	fmt.Println("运行手机游戏")
}

type HandsetAddrList struct {
}

func (hal HandsetAddrList) Run() {
	fmt.Println("运行手机通讯录")
}

type HandsetBrand interface {
	SetHandsetSoft(soft HandsetSoft)
	Run()
}

type HandsetBrandN struct {
	soft HandsetSoft
}

func (hn *HandsetBrandN) SetHandsetSoft(soft HandsetSoft) {
	hn.soft = soft
}

func (hn *HandsetBrandN) Run() {
	hn.soft.Run()
}

type HandsetBrandM struct {
	soft HandsetSoft
}

func (hm *HandsetBrandM) SetHandsetSoft(soft HandsetSoft) {
	hm.soft = soft
}

func (hm *HandsetBrandM) Run() {
	hm.soft.Run()
}
