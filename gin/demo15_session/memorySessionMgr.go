package session

import (
	"errors"
	"sync"

	uuid "github.com/satori/go.uuid"
)

//Memory Session管理中间件，实现sessionMgr
type MemorySessionMgr struct {
	SessionMap map[string]Session
	rwlock     sync.RWMutex
}

//构造函数
func NewMemorySessionMgr() *MemorySessionMgr {
	sm := &MemorySessionMgr{
		SessionMap: make(map[string]Session, 1024),
	}
	return sm
}

func (mm *MemorySessionMgr) Init(addr string, options ...string) (err error) {
	return
}

//创建Memory session
func (mm *MemorySessionMgr) CreateSession() (session Session, err error) {
	mm.rwlock.Lock()
	defer mm.rwlock.Unlock()
	//sessionId -> go get github.com/satori/go.uuid
	id := uuid.NewV4()
	sessionId := id.String()
	//创建单个Memory session
	session = NewMemorySession(sessionId)
	//加入mgr大map
	mm.SessionMap[sessionId] = session
	return
}

func (mm *MemorySessionMgr) GetSession(sessionId string) (session Session, err error) {
	mm.rwlock.Lock()
	defer mm.rwlock.Unlock()
	session, ok := mm.SessionMap[sessionId]
	if !ok {
		err = errors.New("session not exist")
		return
	}
	return
}
