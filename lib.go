/*
zazabul: A configuration file format


Sample

			// email is used for communication
			// email must follow the format for email ie username@host.ext
			// email is compulsory
			email: banker@banban.com

			// region should be gotten from google cloud documentation
			region: us-central1

			// zone should be gotten from google cloud documentation
			// zone usually derived from the regions and ending with -a or -b or -c
			zone: us-central1-a


*/
package zazabul

import (
	"strings"
	"github.com/pkg/errors"
	"os"
	"fmt"
)


type ConfigItem struct {
	Name string
	Comment string
	Value string
}


type Config struct {
	Items []ConfigItem
}


// Returns a config written in the zazabul format.
func LoadConfigFile(path string) (Config, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return Config{}, errors.Wrap(err, "os error")
	}

	return ParseConfig(string(raw))
}


// Default (most comfortable) way of creating a Config object.
func ParseConfig(str string) (Config, error) {
	conf := Config{}
	items := make([]ConfigItem, 0)

	lines := strings.Split(str, "\n")

	item := ConfigItem{}

	var comment string
	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "//") {
			comment += line + "\n"
		} else if line != "" {
			colonIndex := 0

			for i, ch := range line {
				if fmt.Sprintf("%c", ch) == ":" {
					colonIndex = i
					break
				}
			}

			if colonIndex == 0 {
				continue
			}

			if comment != "" {
				item.Comment = comment
				comment = ""
			}

			item.Name = strings.ToLower(strings.TrimSpace(line[0: colonIndex]))
			item.Value = strings.TrimSpace(line[colonIndex + 1 : ])
			items = append(items, item)
		}
	}
	conf.Items = items
	return conf, nil
}


// Update does not add a new entry. To do that you need to add it to the template you used to create
// the Config object.
func (conf *Config) Update(items map[string]string) {
	for k, v := range items {
		done := false
		for i, item := range conf.Items {
			if k == item.Name {
				conf.Items[i] = ConfigItem{item.Name, item.Comment, v}
				done = true
				break
			}
		}

		if done == false {
			conf.Items = append(conf.Items, ConfigItem{k, "", v})
		}
	}
}


func (conf *Config) Get(configName string) string {
	for _, item := range conf.Items {
		if configName == item.Name {
			return item.Value
		}
	}
	return ""
}


func doesPathExists(p string) bool {
	if _, err := os.Stat(p); os.IsNotExist(err) {
		return false
	}
	return true
}


// Writes the Config to a file whose path was given as path
func (conf *Config) Write(path string) error {
	var outString string
	for _, item := range conf.Items {
		outString += item.Comment + item.Name + ": " + item.Value + "\n\n\n"
	}
	err := os.WriteFile(path, []byte(outString), 0777)
	if err != nil {
		return errors.Wrap(err, "os error")
	}
	return nil
}
