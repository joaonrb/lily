package lily

//
// Author João Nuno.
//
// joaonrb@gmail.com
//

import (
	"testing"
)

func TestSessionLazinessOnGet(t *testing.T) {
	sessionDummy := NewSession("LazinessOnGetDummy")
	if sessionDummy.session != nil {
		t.Error("Session attribute in dummy should be nil.")
	}

	value := sessionDummy.Get("gummy bears")
	if value != nil {
		t.Error("Gummy bears should not be a key to no object.")
	}

	if sessionDummy.session == nil {
		t.Error("Session attribute in dummy should not be nil anymore.")
	}
}

func TestSessionLazinessOnSet(t *testing.T) {
	sessionDummy := NewSession("LazinessOnSetDummy")
	if sessionDummy.session != nil {
		t.Error("Session attribute in dummy should be nil.")
	}

	sessionDummy.Set("Pinokio", "lie")

	if sessionDummy.session == nil {
		t.Error("Session attribute in dummy should not be nil anymore.")
	}
}

func TestSessionSetGetFlowForString(t *testing.T) {
	sessionDummy := NewSession("SetGetFlowForStringDummy")
	sessionDummy.Set("Message", "F you Steve")
	SaveSession(sessionDummy.Cookie, sessionDummy.session)

	sessionSteve := NewSession("SetGetFlowForStringDummy")
	message := sessionSteve.Get("Message")
	if message != "F you Steve" {
		t.Error("Message from cache not the one expected. Expected \"F you Steve\" Got \"%s\" instead", message)
	}
}

func TestSessionSetGetFlowForInt(t *testing.T) {
	sessionDummy := NewSession("SetGetFlowForIntDummy")
	sessionDummy.Set("age", 29)
	SaveSession(sessionDummy.Cookie, sessionDummy.session)

	sessionSteve := NewSession("SetGetFlowForIntDummy")
	age := sessionSteve.Get("age")
	if age != 29 {
		t.Error("Message from cache not the one expected. Expected 29 Got %d instead", age)
	}
}

func TestSessionSetGetFlowForStruct(t *testing.T) {
	type dummy struct {
		name string
		age  int
	}
	sessionDummy := NewSession("SetGetFlowForStructDummy")
	sessionDummy.Set("person", &dummy{"João", 29})
	SaveSession(sessionDummy.Cookie, sessionDummy.session)

	sessionSteve := NewSession("SetGetFlowForStructDummy")
	me := sessionSteve.Get("person").(*dummy)
	if me.name != "João" {
		t.Error("Dummy name not the one expected. Expected \"João\" Got \"%s\" instead", me.name)
	}
	if me.age != 29 {
		t.Error("Dummy age not the one expected. Expected 29 Got %d instead", me.age)
	}
}