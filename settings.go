// Author João Nuno.
//
// joaonrb@gmail.com
//
package lily

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

const (
	SETTINGS_FILE       = "SERVER_SETTINGS"
	defaultSettingsPath = "settings.yaml"
)

var Settings = map[string]interface{}{}

func init() {
	path := os.Getenv(SETTINGS_FILE)
	if path == "" {
		path = defaultSettingsPath
	}

	// Open file
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error %s", err.Error())
		return
	}
	defer file.Close()

	// Read file
	setting, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Error %s", err.Error())
		return
	}

	// Read yaml to map
	err = yaml.Unmarshal(setting, Settings)
	if err != nil {
		fmt.Printf("Error %s", err.Error())
		return
	}
}
