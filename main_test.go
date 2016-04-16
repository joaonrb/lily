package lily

//
// Author João Nuno.
//
// joaonrb@gmail.com
//

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

const testSettingsLocation = "/tmp/lily_settings.yaml"

const testSettings = `
loggers:
  default:
    type:   console
    layout: "%{level:.4s} %{time:2006-01-02 15:04:05.000} %{shortfile} %{message}"
    level:  debug

accesslog:
  type: console

apps:
  cache:
    type: memory
`

func TestMain(m *testing.M) {
	defer os.Remove(testSettingsLocation)
	err := ioutil.WriteFile(testSettingsLocation, []byte(testSettings), 0644)
	if err != nil {
		fmt.Print("Tmp file couldn't be writen becauser error %s", err.Error())
		os.Exit(1)
	}

	// Starting test
	err = Init(testSettingsLocation)
	if err != nil {
		fmt.Print("Couldn't init configuration because error %s", err.Error())
	}
	LoadCache(Configuration)
	os.Exit(m.Run())
}
