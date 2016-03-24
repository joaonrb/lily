//
// Author João Nuno.
// 
// joaonrb@gmail.com
//
package sessions

import "time"

var sessionStore ISessionStore

type ISessionStore interface {
	GetSession(string) (Session, bool)
	SetSession(string, time.Time, Session)
	UpdateSession(string, Session)
}
