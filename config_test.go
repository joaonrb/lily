package lily

//
// Author João Nuno.
//
// joaonrb@gmail.com
//

import (
	"testing"
)

func TestParse(t *testing.T) {

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
