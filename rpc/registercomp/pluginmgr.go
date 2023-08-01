package registry

import (
	"context"
	"fmt"
	"sync"
)

// 插件管理类
// 大map管理，key字符串，value为Register接口对象
// 支持用户自定义调用，自定义插件
// 实现注册中心的初始化，供系统使用

//声明管理者结构体
type PluginMgr struct {
	//map维护所有插件
	plugins map[string]Registry
	lock    sync.Mutex
}

var (
	pluginMgr = &PluginMgr{
		plugins: make(map[string]Registry),
	}
)

//插件注册
func RegisterPlugin(registry Registry) (err error) {
	return pluginMgr.registerPlugin(registry)
}

func (p *PluginMgr) registerPlugin(plugin Registry) (err error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	//先检查是否已注册
	_, ok := p.plugins[plugin.Name()]
	if ok {
		err = fmt.Errorf("registry plugin %s exists", plugin.Name())
		return
	}
	p.plugins[plugin.Name()] = plugin
	return
}

//初始化注册中心
func InitRegistry(ctx context.Context, name string, opts ...Option) (registry Registry, err error) {
	return pluginMgr.initRegistry(ctx, name, opts...)
}

func (p *PluginMgr) initRegistry(ctx context.Context, name string, opts ...Option) (registry Registry, err error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	//初始化前先检查服务是否已存在
	plugin, ok := p.plugins[name]
	if !ok {
		err = fmt.Errorf("plugin %s not exist", name)
		return
	}
	//存在则返回
	registry = plugin
	//插件初始化
	plugin.Init(ctx, opts...)
	return
}
