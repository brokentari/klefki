package session

import (
	"sync"
	"time"

	"github.com/alexedwards/scs/v2"
)

type SessionManager struct {
	*scs.SessionManager
}

var (
	session *SessionManager
	once    sync.Once
)

func GetSession() *SessionManager {
	once.Do(func() {
		sessionManager := scs.New()
		sessionManager.Lifetime = 24 * time.Hour

		session = &SessionManager{sessionManager}
	})

	return session
}
