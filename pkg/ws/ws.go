package ws

import "gopkg.in/olahol/melody.v1"

var m *melody.Melody
var sessionMaster *SessionM

func GetMelody() *melody.Melody {
	return m
}
func GetSessionMaster() *SessionM {
	return sessionMaster
}

func init() {
	m = melody.New()
	sessionMaster = NewSessonM()
}
