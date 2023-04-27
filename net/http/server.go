package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("./index.html")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("无法解析文件,遇到错误:%v", err)))
	}
	w.Write(b)
}

func main() {
	http.HandleFunc("/post", f1)

	err := http.ListenAndServe("127.0.0.1:9009", nil)
	if err != nil {
		fmt.Printf("http listen and serve failed, err=%v\n", err)
		return
	}
}
