//
// Author João Nuno.
//
// joaonrb@gmail.com
//
package lily

import (
	yalm "gopkg.in/yaml.v2"
	"io/ioutil"
)

var(
	Configuration *Settings
)

type Settings struct {
	Bind            string                  `yaml:"bind,omitempty"`
	Port            int                     `yaml:"port,omitempty"`
	ReadTimeout     int                     `yaml:"read_timeout,omitempty"`   // In Millis
	WriteTimeout    int                     `yaml:"write_timeout,omitempty"`  // In Millis
	Https           bool                    `yaml:"https,omitempty"`
	SSLCertificate  string                  `yaml:"ssl_certificate,omitempty"`
	SSLKey          string                  `yaml:"ssl_key,omitempty"`
	StaticFiles     string                  `yaml:"static_files,omitempty"`
	Loggers         map[string]LogSettings  `yaml:"loggers,omitempty"`
	AccessLog       AccessLogSettings       `yaml:"accesslog,omitempty"`
	Middleware      []string                `yaml:"middleware,omitempty"`
}

type LogSettings struct {
	Type   string `yaml:"type,omitempty"`
	Layout string `yaml:"layout,omitempty"`
	Path   string `yaml:"path,omitempty"`
	Level  string `yaml:"level,omitempty"`
}

type AccessLogSettings struct {
	Type string `yaml:"type,omitempty"`
	Path string `yaml:"path,omitempty"`
}


func Init(path string) error {
	setting, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	Configuration = &Settings{}
	return yalm.Unmarshal(setting, Configuration)
}