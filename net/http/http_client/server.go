package main

import (
	"fmt"
	"net/http"
)

func f2(w http.ResponseWriter, r *http.Request) {
	//获取客户端请求参数
	fmt.Println("Header=", r.Header)
	fmt.Println("Method=", r.Method)
	fmt.Println("URL=", r.URL)
	fmt.Println("Body=", r.Body) // GET请求空值
	//解析URL中的参数
	var name, age string
	name = r.URL.Query().Get("name")
	age = r.URL.Query().Get("age")
	fmt.Printf("请求参数 name=%v, age=%v\n", name, age)
	w.Write([]byte("success"))
}

func main() {

	http.HandleFunc("/xxx", f2)

	// 默认监听方式
	http.ListenAndServe("127.0.0.1:9009", nil)
	// //自定义Server
	// myServer := &http.Server{
	// 	Addr:           "http://127.0.0.1:9009",
	// 	Handler:        nil,
	// 	ReadTimeout:    2000 * time.Second,
	// 	WriteTimeout:   2000 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	// myServer.ListenAndServe()

}
