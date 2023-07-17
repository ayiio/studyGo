package session

//sessionMgr接口
//sessionID-session
type SessionMgr interface {
	Init(addr string, options ...string) error
	CreateSession() (session Session, err error)
	GetSession(sessionId string) (session Session, err error)
}
