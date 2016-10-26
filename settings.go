// Author João Nuno.
//
// joaonrb@gmail.com
//
package lily

import (
	"gopkg.in/yaml.v2"
)

var Settings = map[string]interface{}{}

func SetSettings(path string) error {
	err := yaml.Unmarshal([]byte(path), Settings)
	if err != nil {
		return err
	}
	return nil
}
