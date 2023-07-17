package session

import (
	"errors"
	"sync"
)

//单个Memory session，实现session接口
type MemorySession struct {
	sessionId string
	// kv map
	data   map[string]interface{}
	rwlock sync.RWMutex
}

//构造函数, memory session
func NewMemorySession(id string) *MemorySession {
	s := &MemorySession{
		sessionId: id,
		data:      make(map[string]interface{}, 16),
	}
	return s
}

//set memory session
func (m *MemorySession) Set(key string, value interface{}) (err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	m.data[key] = value
	return
}

//get memory session
func (m *MemorySession) Get(key string) (value interface{}, err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	value, ok := m.data[key]
	if !ok {
		err = errors.New("key not exists in session")
		return
	}
	return
}

//delete memory session
func (m *MemorySession) Del(key string) (err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	delete(m.data, key)
	return
}

//Memory，save无操作
func (m *MemorySession) Save() (err error) {
	return
}
