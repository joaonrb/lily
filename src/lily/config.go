//
// Copyright (c) João Nuno. All rights reserved.
//
package lily

var(
	Configuration *Settings
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
	Type string `json:"type,omitempty"`
	Path string `json:"path,omitempty"`
}
