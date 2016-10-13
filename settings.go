// Package lily
// Author João Nuno.
// 
// joaonrb@gmail.com
//
package lily


type SLogger struct {
	Type   string `yaml:"type,omitempty"`
	Layout string `yaml:"layout,omitempty"`
	Path   string `yaml:"path,omitempty"`
	Level  string `yaml:"level,omitempty"`
}