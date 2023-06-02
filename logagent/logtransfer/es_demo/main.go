package main

import (
	"context"
	"fmt"

	"encoding/json"

	"github.com/olivere/elastic/v7"
)

// ES demo
// doc: https://pkg.go.dev/github.com/olivere/elastic

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func main() {
	//初始化连接，获取一个client
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetBasicAuth("user", "password"))
	if err != nil {
		panic(err)
	}
	fmt.Println("connect to es success")
	p1 := Person{Name: "test", Age: 21, Married: false}
	//链式操作
	//存入
	put1, err := client.Index().Index("student").Id("go").BodyJson(p1).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed user '%s' to index '%s', id '%s'\n", put1.Id, put1.Index, put1.Type)

	//获取
	resp, err := client.Get().Index("student").Id("go").Do(context.Background())
	if err != nil {
		panic(err)
	}
	if resp.Found {
		fmt.Printf("get id='%s' from index='%s', version='%d'\n", resp.Id, resp.Index, resp.Version)
	}
	
	//解析
	var np Person
	data, err := resp.Source.MarshalJSON()
	if err != nil {
		panic(err)
	}
	json.Unmarshal(data, &np)
	fmt.Println(np.Name, np.Age, np.Married)
}

