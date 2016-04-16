package lily

//
// Author João Nuno.
//
// joaonrb@gmail.com
//

import (
	"io/ioutil"
	"os"
	"testing"
)

const sessionTestSettingsLocation = "/tmp/lily_sessions_test_parse.yaml"

const sessionTestSettings = `
apps:
  cache:
    type: memory
`

func TestSessionGet(t *testing.T) {
	defer os.Remove(sessionTestSettingsLocation)
	err := ioutil.WriteFile(sessionTestSettingsLocation, []byte(sessionTestSettings), 0644)
	if err != nil {
		t.Fatalf("Tmp file couldn't be writen becauser error %s", err.Error())
	}

	// Starting test
	err = Init(sessionTestSettingsLocation)
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
