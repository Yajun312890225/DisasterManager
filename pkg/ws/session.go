package ws

import (
	"DisasterManager/model"
	"log"
	"sync"

	"gopkg.in/olahol/melody.v1"
)

//Session 一个session代表一个连接
type Session struct {
	CID      string
	Msession *melody.Session
	lock     sync.Mutex
}

//NewSession 新建连接
func NewSession(cid string, s *melody.Session) *Session {
	return &Session{
		CID:      cid,
		Msession: s,
	}
}

func (s *Session) write(msg string) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.Msession.Write([]byte(msg))
}

func (s *Session) close() {
	s.Msession.Close()
}

//SessionM SESSION管理类
type SessionM struct {
	sessions  sync.Map
	onlineCID sync.Map
}

//NewSessonM 新建管理类
func NewSessonM() *SessionM {
	return &SessionM{}
}

//GetSessionByCID 获取session
func (s *SessionM) GetSessionByCID(cid string) *Session {
	tem, exit := s.sessions.Load(cid)
	if exit {
		if sess, ok := tem.(*Session); ok {
			return sess
		}
	}
	return nil
}

//SetSession 设置session
func (s *SessionM) SetSession(cid string, ms *melody.Session) {
	sess := NewSession(cid, ms)
	s.sessions.Store(cid, sess)
}

//DelSessionByCID 关闭连接并删除
func (s *SessionM) DelSessionByCID(cid string) {
	tem, exit := s.sessions.Load(cid)
	if exit {
		if sess, ok := tem.(*Session); ok {
			sess.close()
		}
	}
	s.onlineCID.Delete(cid)
	s.sessions.Delete(cid)
}

//WriteToWeb 向Web推送位置信息
func (s *SessionM) WriteToWeb(msg []byte) {
	s.sessions.Range(func(key, val interface{}) bool {
		if val, ok := val.(*Session); ok {
			if t, exist := val.Msession.Get("type"); exist {
				if t == model.LoginFromWeb {
					if err := val.write(string(msg)); err != nil {
						s.DelSessionByCID(key.(string))
						log.Println(err)
					}
				}
			}

		}
		return true
	})
}

//WriteByCID 向单个客户端发送信息
func (s *SessionM) WriteByCID(cid string, msg []byte) bool {
	tem, exit := s.sessions.Load(cid)
	if exit {
		if sess, ok := tem.(*Session); ok {
			if err := sess.Msession.Write(msg); err == nil {
				return true
			}
		}
	}
	s.DelSessionByCID(cid)
	return false
}
