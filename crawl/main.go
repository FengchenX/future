package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

/*
([/w-]+/.)+[/w-]+.([^a-z])(/[/w-: ./?%&=]*)?|[a-zA-Z/-/.][/w-]+.([^a-z])(/[/w-: ./?%&=]*)?
(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]

<img class="BDE_I.*?src="(.*?)".*?size="(.*?)".*?width="(.*?)".*?height="(.*?)"
*/

var reUrl = regexp.MustCompile(`(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]`)

//提取内容正则表达式
var reText = regexp.MustCompile(`<img class="BDE_I.*?src="(.*?)".*?size="(.*?)".*?width="(.*?)".*?height="(.*?)"`)

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

var lockx = make(chan int, 1)

func SafeRun(f func()) {
	<-lockx
	f()
	lockx <- 1
}

var visited map[string]bool = make(map[string]bool)

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher, quit chan int) {
	// TODO: 并行的抓取 URL。
	// TODO: 不重复抓取页面。
	// 下面并没有实现上面两种情况：
	defer func() {
		quit <- 1
	}()
	if depth <= 0 || visited[url] {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	SafeRun(func() {
		visited[url] = true
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	sub_quit := make(chan int, len(urls))
	for _, u := range urls {
		Crawl(u, depth-1, fetcher, sub_quit)
	}
	for i := 0; i < len(urls); i++ {
		<-sub_quit
	}
	return
}

func main() {
	lockx <- 1
	quit := make(chan int, 1)
	Crawl("https://studygolang.com/", 4, fetcher, quit)
	fmt.Println(<-quit)
}

// fakeFetcher 是返回若干结果的 Fetcher。
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	str := string(buf)
	urls := reUrl.FindAllStringSubmatch(str, -1)
	texts := reText.FindAllStringSubmatch(str, -1)
	var result fakeResult
	result.body = ""
	for _, item := range urls {
		result.urls = append(result.urls, item[0])
		fmt.Println(item[0])
	}
	for _, tex := range texts {
		result.body = result.body + tex[0]
		fmt.Println(tex[0])
	}
	return "", result.urls, nil
}

// fetcher 是填充后的 fakeFetcher。
var fetcher = make(fakeFetcher)
