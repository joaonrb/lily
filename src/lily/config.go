//
// Copyright (c) Telefonica I+D. All rights reserved.
//
package lily

var(
	Settings *Settings
)

type Settings struct {
	Loggers   []LogSettings    `json:"loggers,omitempty"`
	AccessLog AccessLogSetings `json:"accesslog,omitempty"`
}

type LogSettings struct {
	Type   string `json:"type,omitempty"`
	Layout string `json:"layout,omitempty"`
	Path   string `json:"path,omitempty"`
	Level  string `json:"level,omitempty"`
}

type AccessLogSetings struct {
	Path string `json:"path,omitempty"`
}
