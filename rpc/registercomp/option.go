package registry

import "time"

type Options struct {
	//地址
	Addrs []string
	//超时时间
	TimeOut time.Duration
	//心跳时间
	HeartBeat int64
	//注册地址  /a/b/c/d/xxx/10.xx.xx.1有前缀
	RegistryPath string
}

//定义函数类型的变量
type Option func(opts *Options)

func SetAddr(addrs []string) Option {
	return func(opts *Options) {
		opts.Addrs = addrs
	}
}

func SetTimeOut(tm time.Duration) Option {
	return func(opts *Options) {
		opts.TimeOut = tm
	}
}

func SetHeartBeat(ht int64) Option {
	return func(opts *Options) {
		opts.HeartBeat = ht
	}
}

func SetRegistryPath(rgpath string) Option {
	return func(opts *Options) {
		opts.RegistryPath = rgpath
	}
}
