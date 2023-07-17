package session

import (
	"fmt"
	"strings"
)

var (
	sessionMgr SessionMgr
)

//中间件供用户选择使用哪个版本(Memory/Redis)
func Init(provider string, addr string, options ...string) (err error) {
	switch strings.ToLower(provider) {
	case "memory":
		sessionMgr = NewMemorySessionMgr()
	case "redis":
		sessionMgr = NewRedisSessionMgr()
	default:
		err = fmt.Errorf("不支持该种版本: %s", provider)
		return
	}
	err = sessionMgr.Init(addr, options...)
	return
}
