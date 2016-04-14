//
// Author João Nuno.
//
// joaonrb@gmail.com
//
package lily

import (
	"io/ioutil"
	"os"
	"testing"
)

const TMP_LOCATION = "/tmp/lily_test_parse.yaml"

const YAML_SETTINGS_EXAMPLE = `
loggers:
  default:
    type:   console
    layout: "%{level:.4s} %{time:2006-01-02 15:04:05.000} %{shortfile} %{message}"
    level:  debug

accesslog:
  type: console
`

func TestParse(t *testing.T) {
	defer os.Remove(TMP_LOCATION)
	err := ioutil.WriteFile(TMP_LOCATION, []byte(YAML_SETTINGS_EXAMPLE), 0644)
	if err != nil {
		t.Fatalf("Tmp file couldn't be writen becauser error %s", err.Error())
	}

	// Starting test
	err = Init(TMP_LOCATION)
	if err != nil {
		t.Fatalf("Couldn't init configuration because error %s", err.Error())
	}

	// Test loggers
	switch {
	case Configuration == nil:
		t.Fatalf("Configuration not inicialized")
	case len(Configuration.Loggers) != 1:
		t.Errorf("Loggers size is %d and 1 is expected", len(Configuration.Loggers))
	case Configuration.Loggers["default"].Type != "console":
		t.Errorf("Logger type is \"%s\" and \"console\" is expected", Configuration.Loggers["default"].Type)
	case Configuration.Loggers["default"].Layout != "%{level:.4s} %{time:2006-01-02 15:04:05.000} %{shortfile} "+
		"%{message}":
		t.Errorf("Logger layout is \"%s\" and \"%{level:.4s} %{time:2006-01-02 15:04:05.000} %{shortfile} %{message}\""+
			"is expected", Configuration.Loggers["default"].Layout)
	case Configuration.Loggers["default"].Level != "debug":
		t.Error("logger level is \"%s\" and \"debug\" is expected", Configuration.Loggers["default"].Level)
	}

	// Test access log
	if Configuration.AccessLog.Type != "console" {
		t.Errorf("Access log type is \"%s\" and \"console\" is expected", Configuration.AccessLog.Type)
	}

}
