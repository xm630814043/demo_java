package requetsq

import (
	"fmt"
	"github.com/jackdanger/collectlinks"
	"net/http"
	"net/url"
	"time"
)

//使用 urllib.parse.urljoin将相对的一个地址组合成一个url，对于输入没有限制，开头必须是http://或者https://，否则将不组合前面的部分
func urlJoin(href, base string) string {
	uri, err := url.Parse(href)
	if err != nil {
		return " "
	}
	baseUrl, err := url.Parse(base)
	if err != nil {
		return " "
	}
	return baseUrl.ResolveReference(uri).String()
}

func Downloads(url string, queue chan string) {
	var visited = make(map[string]bool)

	visited[url] = true
	timeout := time.Duration(5 * time.Second)

	client := &http.Client{
		Timeout:timeout,
	}
	req, _ := http.NewRequest("GET", url, nil)
	// 自定义Header
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http get error", err)
		return
	}
	//函数结束后关闭相关链接
	defer resp.Body.Close()

	links := collectlinks.All(resp.Body)
	for _, link := range links {
		absolute := urlJoin(link, url)
		if url != " " {
			if !visited[absolute] {
				fmt.Println("parse url", absolute)
				go func() {
					queue <- absolute
				}()
			}
		}
	}
}
