//
// Author João Nuno.
// 
// joaonrb@gmail.com
//
package sessions

import (
	"testing"
	"io/ioutil"
	"os"
	"lily"
)

const TMP_LOCATION = "/tmp/lily_sessions_test_parse.yaml"

const YAML_SETTINGS_EXAMPLE = `
apps:
  cache:
    type: memory
`

func TestSessionGet(t *testing.T)  {
	defer os.Remove(TMP_LOCATION)
	err := ioutil.WriteFile(TMP_LOCATION, []byte(YAML_SETTINGS_EXAMPLE), 0644)
	if err != nil {
		t.Fatalf("Tmp file couldn't be writen becauser error %s", err.Error())
	}

	// Starting test
	err = lily.Init(TMP_LOCATION)
	if err != nil {
		t.Fatalf("Couldn't init configuration because error %s", err.Error())
	}

	LoadCache(lily.Configuration)
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
