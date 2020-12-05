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
	"io/ioutil"
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
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return Config{}, errors.Wrap(err, "ioutil error")
	}

	return ParseConfig(string(raw))
}


// Default (most comfortable) way of creating a Config object.
func ParseConfig(str string) (Config, error) {
	conf := Config{}
	items := make([]ConfigItem, 0)

	lines := strings.Split(str, "\n")

	index := 0
	item := ConfigItem{}
	var comment string

	for {
		if index >= len(lines) {
			break
		}

		line := strings.TrimSpace(lines[index])

		if strings.HasPrefix(line, "//") {
			comment += line + "\n"
		} else if line != "" {
			if comment != "" {
				item.Comment = comment
				comment = ""
			}
			lineParts := strings.Split(line, ":")
			if len(lineParts) != 2 {
				return conf, errors.New("Each config line (not a comment line) must only contain one ':'.")
			}
			item.Name = strings.ToLower(strings.TrimSpace(lineParts[0]))
			item.Value = strings.TrimSpace(lineParts[1])
			items = append(items, item)
		}

		index += 1
	}
	conf.Items = items
	return conf, nil
}


// Update does not add a new entry. To do that you need to add it to the template you used to create
// the Config object.
func (conf *Config) Update(items map[string]string) {
	for k, v := range items {
		for i, item := range conf.Items {
			if k == item.Name {
				conf.Items[i] = ConfigItem{item.Name, item.Comment, v}
			}
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
	err := ioutil.WriteFile(path, []byte(outString), 0777)
	if err != nil {
		return errors.Wrap(err, "ioutil error")
	}
	return nil
}
