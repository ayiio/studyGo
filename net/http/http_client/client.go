package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// 简单GET请求
	resp, err := http.Get("http://127.0.0.1:9009/xxx?name=xxx&age=20")
	if err != nil {
		fmt.Printf("http GET request failed, err=%v\n", err)
		return
	}
	// 从resp中读取服务端返回-方式1
	var data [1024]byte
	n, _ := resp.Body.Read(data[:])
	resp.Body.Close() // 使用完response后必须关闭回复的主体
	fmt.Println("resp.Body.Read=", string(data[:n]))

	// 自组装请求参数，使用encode
	urlObj, _ := url.Parse("http://127.0.0.1:9009/xxx")
	urldata := url.Values{}
	urldata.Set("name", "自定义1")
	urldata.Set("age", "自定义2")
	urldata.Set("gender", "female")
	urldata.Set("hobby", "basketball")
	queryStr := urldata.Encode()
	fmt.Println("编码后的请求参数=", queryStr)
	urlObj.RawQuery = queryStr
	fmt.Println("编码后的请求URL=", urlObj.String())
	// 创建新请求
	req, err := http.NewRequest("GET", urlObj.String(), nil)
	if err != nil {
		fmt.Printf("new request failed, err=%v\n", err)
		return
	}
	// DefaultClient.Do() 发起请求
	// resp2, err := http.DefaultClient.Do(req)
	// 可自定义请求的Client，禁用keepAlives，创建短连接
	// 用于偶发或临时连接的需求
	tr := &http.Transport{
		DisableKeepAlives: true,
	}
	client := http.Client{
		Transport: tr,
	}
	resp2, err := client.Do(req)
	if err != nil {
		fmt.Printf("defaultClient do failed, err=%v\n", err)
		return
	}
	defer resp2.Body.Close() // 使用完response后必须关闭回复的主体
	fmt.Println("服务端返回=", resp2)
	// 从resp中读取服务端返回-方式2
	b, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		fmt.Printf("read body by ioutil failed, err=%v\n", err)
		return
	}
	fmt.Println("ioutil.read.resp.body=", string(b))

	// 简单POST请求  &buf: bufio对象
	// http.Post("http://example.com/upload", "image/jpeg", &buf)

	// http.PostForm("http://example.com/form",
	// 	url.Values{"key": {"value"}, "id": {"123"}})

	// 需要管理代理、TLS配置、keep-alive、压缩或其他设置时，可以通过创建一个Transport
	// client和transport类型都可以安全的被多个goroutine同时使用，一次建立，多次重用，提高效率
	// 可以定义全局变量共同一个client, 全局变量保证DisableKeepAlives为false，用于请求比较频繁的场景
	// tr := &http.Transport{
	// 	TLSClientConfig:    &tls.Config{RootCAs: pool},
	// 	DisableCompression: true,
	// }
	// client := &http.Client{Transport: tr}
	// resp, err := client.Get("http://example.com")
}
