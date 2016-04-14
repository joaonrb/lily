//
// Author João Nuno.
// 
// joaonrb@gmail.com
//
package lily

import (
	"testing"
	"io/ioutil"
	"os"
)

const SESSION_LOCATION = "/tmp/lily_sessions_test_parse.yaml"

const SESSION_SETTINGS_EXAMPLE = `
apps:
  cache:
    type: memory
`

func TestSessionGet(t *testing.T)  {
	defer os.Remove(SESSION_LOCATION)
	err := ioutil.WriteFile(SESSION_LOCATION, []byte(SESSION_SETTINGS_EXAMPLE), 0644)
	if err != nil {
		t.Fatalf("Tmp file couldn't be writen becauser error %s", err.Error())
	}

	// Starting test
	err = Init(SESSION_LOCATION)
	if err != nil {
		t.Fatalf("Couldn't init configuration because error %s", err.Error())
	}

	LoadCache(Configuration)
	sessionDummy := NewSession("dummy")
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
