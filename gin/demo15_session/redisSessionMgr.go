package session

import (
	"errors"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
	uuid "github.com/satori/go.uuid"
)

type RedisSessionMgr struct {
	//redis地址
	addr string
	//密码
	password string
	//连接池
	pool *redis.Pool
	//锁
	rwlock sync.RWMutex
	//管理者map
	sessionMap map[string]Session
}

//构造函数
func NewRedisSessionMgr() SessionMgr {
	rsm := &RedisSessionMgr{
		sessionMap: make(map[string]Session, 32),
	}
	return rsm
}

//初始化redisSessionMgr
func (rm *RedisSessionMgr) Init(addr string, options ...string) (err error) {
	rm.addr = addr
	//若有其他参数
	if len(options) > 0 {
		rm.password = options[0]
	}
	//创建连接池
	rm.pool = &redis.Pool{
		MaxIdle:     64,
		MaxActive:   1000,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", rm.addr)
			if err != nil {
				return nil, err
			}
			//判断密码
			if _, err := conn.Do("AUTH", rm.password); err != nil {
				conn.Close()
				return nil, err
			}
			return conn, err
		},
		//连接测试, 开发环境
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			_, err := conn.Do("PING")
			return err
		},
	}
	return
}

//创建RedisSession
func (rm *RedisSessionMgr) CreateSession() (session Session, err error) {
	rm.rwlock.Lock()
	defer rm.rwlock.Unlock()
	//生成sessionID
	id := uuid.NewV4()
	sessionID := id.String()
	//创建redisSession
	session = NewRedisSession(sessionID, rm.pool)
	//加入到大map
	rm.sessionMap[sessionID] = session
	return
}

//根据sessionID获取session
func (rm *RedisSessionMgr) GetSession(sessionId string) (session Session, err error) {
	rm.rwlock.Lock()
	defer rm.rwlock.Unlock()
	session, ok := rm.sessionMap[sessionId]
	if !ok {
		err = errors.New("session not exist")
		return
	}
	return
}
