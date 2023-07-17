package session

import (
	"encoding/json"
	"sync"

	"github.com/gomodule/redigo/redis"
)

type RedisSession struct {
	sessionId string
	pool      *redis.Pool
	//设置session，可以先放在内存的map中
	//批量导入redis，提升性能
	sessionMap map[string]interface{}
	//读写锁
	rwlock sync.RWMutex
	//记录内存中的map是否被操作
	flag int
}

//常亮状态
const (
	//内存数据无变化
	SessionFlagNone = iota
	//有变化
	SessionFlagModify
)

//构造函数
func NewRedisSession(id string, pool *redis.Pool) *RedisSession {
	rds := &RedisSession{
		sessionId:  id,
		pool:       pool,
		sessionMap: make(map[string]interface{}, 16),
		flag:       SessionFlagNone,
	}
	return rds
}

//保存在内存中的Map
func (r *RedisSession) Set(key string, value interface{}) (err error) {
	//加锁
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	//设置值
	r.sessionMap[key] = value
	//标记记录
	r.flag = SessionFlagModify
	return
}

//获取session
func (r *RedisSession) Get(key string) (session interface{}, err error) {
	//加锁
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	//先判断内存是否有
	session, ok := r.sessionMap[key]
	if !ok {
		//没有再从redis中取
		conn := r.pool.Get()
		reply, err1 := conn.Do("GET", key)
		if err1 != nil {
			return
		}
		//转字符串
		data, err2 := redis.String(reply, err)
		if err2 != nil {
			return
		}
		//反序列化到内存map
		err = json.Unmarshal([]byte(data), &r.sessionMap)
		if err != nil {
			return
		}
	}
	return
}

//删除session
func (r *RedisSession) Del(key string) (err error) {
	//加锁
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	r.flag = SessionFlagModify
	delete(r.sessionMap, key)
	return
}

//从内存转存到redis
func (r *RedisSession) Save() (err error) {
	//加锁
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	//如果数据没变不需要后续操作
	if r.flag != SessionFlagModify {
		return
	}
	//内存中的sessionMap进行序列号
	data, err := json.Marshal(r.sessionMap)
	if err != nil {
		return
	}
	//获取redis连接
	conn := r.pool.Get()
	//保存k-v
	_, err = conn.Do("SET", r.sessionId, string(data))
	//恢复状态
	r.flag = SessionFlagNone
	if err != nil {
		return
	}
	return
}
