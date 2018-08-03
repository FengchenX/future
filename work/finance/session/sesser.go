package session

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"sync"
	"time"
)

//Manager session管理者
type Manager struct {
	//cookieName string
	tokenName   string
	lock        sync.Mutex
	provider    Provider
	maxLifeTime int64
}

//Provider 生成Session
type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64)
}

//Session Seession接口
type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	SessionID() string
}

//NewSessionManager 创建Manager
func NewSessionManager(provideName, tokenName string, maxLifeTime int64) (*Manager, error) {
	provide, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?", provideName)
	}
	return &Manager{tokenName: tokenName, provider: provide, maxLifeTime: maxLifeTime}, nil
}

var provides = make(map[string]Provider)

//Register 注册
func Register(name string, provide Provider) {
	if provide == nil {
		panic("session: Register provide is nil")
	}
	if _, dup := provides[name]; dup {
		panic("session: Register called twice for provide" + name)
	}
	provides[name] = provide
}

func (manager *Manager) sessionID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

//SessionStart Session开始运行
func (manager *Manager) SessionStart(c *gin.Context) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	/*cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		sid := manager.sessionId()
		session, _ = manager.provider.SessionInit(sid)
		cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.maxLifeTime)}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid)
	}
	*/
	//判断是否有token
	token := c.Request.Header.Get(manager.tokenName)
	if token == "" {
		sid := manager.sessionID()
		session, _ = manager.provider.SessionInit(sid)
		//c.JSON(http.StatusOK, lib.Result.Success("登录成功", sid))
	}
	return session
}

//Session 从session池中读取session
func (manager *Manager) Session(c *gin.Context) (session Session) {
	//判断是否有session
	token := c.Request.Header.Get(manager.tokenName)
	if token != "" {
		sid := token
		session, _ = manager.provider.SessionRead(sid)
	}
	return session
}

//SessionDestroy 清理session
func (manager *Manager) SessionDestroy(c *gin.Context) {
	token := c.Request.Header.Get(manager.tokenName)
	if token == "" {
		return
	} else {
		manager.lock.Lock()
		defer manager.lock.Unlock()
		manager.provider.SessionDestroy(token)
	}

	/*
		cookie, err := r.Cookie(manager.cookieName)
		if err != nil || cookie.Value == "" {
			return
		} else {
			manager.lock.Lock()
			defer manager.lock.Unlock()
			manager.provider.SessionDestroy(cookie.Value)
			expiration := time.Now()
			cookie := http.Cookie{Name: manager.cookieName, Path: "/",
			HttpOnly: true, Expires: expiration, MaxAge: -1}
			http.SetCookie(w, &cookie)
		}
	*/
}

//GC 清理过期session
func (manager *Manager) GC() {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.provider.SessionGC(manager.maxLifeTime)
	time.AfterFunc(time.Duration(manager.maxLifeTime), func() { manager.GC() })
}
