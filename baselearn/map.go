package main

import "fmt"

type S struct {
	 name string
}

func main() {
	//map里结构体无法直接寻址，必须取址
	m := map[string]*S{"x":&S{"one"}}
	m["x"].name = "two"
	fmt.Println(m["x"].name)
}
