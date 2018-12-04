package types

import (
	"time"
)

type UserSession struct {
	Channel         string `json:"channel,omitempty"`
	User            string `json:"user,omitempty"`
	Timestamp       string `json:"ts,omitempty"`
	ThreadTimestamp string `json:"thread_ts,omitempty"`
	Intent          string
	DbType          string
	Artifact        string
	Cluster         string
	Sentiment       string
	Updated         time.Time
}

// Sessions is a global variable to store bot sessions
var Sessions = newSessionManager()

func newSessionManager() *SessionManager {
	return &SessionManager{
		Sessions: make(map[string]UserSession),
		Threads:  make(map[string]UserSession),
	}
}

type SessionManager struct {
	Sessions map[string]UserSession
	Threads  map[string]UserSession
}

func (sm *SessionManager) AddSession(us UserSession) {
	sm.Sessions[us.User] = us
	sm.Threads[us.ThreadTimestamp] = us
}

func (sm *SessionManager) GetSession(user string) UserSession {
	return sm.Sessions[user]

}

func (sm *SessionManager) GetThread(user string) UserSession {
	return sm.Sessions[user]

}
