package translate

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"strings"
	"fmt"
)

var translations map[interface{}]interface{}

func LoadTranslationFile(path string) error {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read file: %s", path)
	}

	if err := yaml.Unmarshal([]byte(contents), &translations); err != nil {
		return fmt.Errorf("failed to parse YAML: %s", contents)
	}

	return nil
}

func T(yamlPath string) (string, error) {
	splitPath := strings.Split(yamlPath, ".")
	finalValue := ""
	m := translations
	for i, key := range splitPath {
		value, ok := m[key]
		if !ok {
			return "", fmt.Errorf("could not find key: %s", strings.Join(splitPath[0:i+1], "."))
		}

		stringValue, ok := value.(string)
		if ok {
			if i != len(splitPath) - 1 {
				return "", fmt.Errorf("found string instead of map for key: %s", strings.Join(splitPath[0:i+1], "."))
			}
			finalValue = stringValue
			break
		}

		mapValue, ok := m[key].(map[interface{}]interface{})
		if !ok {
			return "", fmt.Errorf("could not convert value to map for key: %s", strings.Join(splitPath[0:i+1], "."))
		}

		m = mapValue
	}
	return finalValue, nil
}
