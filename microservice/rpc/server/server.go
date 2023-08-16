package main

import (
	"log"
	"net/http"
	"net/rpc"
)

//空的结构体，用于注册
type Area struct {
}

//请求参数
type ParamRequest struct {
	A, B int
}

//计算周长
func (a *Area) MathZ(req *ParamRequest, resp *int) (err error) {
	*resp = 2 * (req.A + req.B)
	return nil
}
//计算面积
func (a *Area) MathM(req *ParamRequest, resp *int) (err error ) {
	*resp = req.A * req.B
	return nil
}

func main() {
	//注册
	rpc.Register(new(Area))
	rpc.HandleHTTP() //绑定HTTP协议
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
