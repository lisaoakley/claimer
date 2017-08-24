package translate

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
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

func T(yamlPath string) string {
	splitPath := strings.Split(yamlPath, ".")
	innerTranslations := translations
	for i, key := range splitPath {
		value, ok := innerTranslations[key]
		if !ok {
			return yamlPath
		}

		stringValue, ok := value.(string)
		if ok && i == (len(splitPath)-1) {
			return stringValue
		}

		mapValue, ok := innerTranslations[key].(map[interface{}]interface{})
		if !ok {
			return yamlPath
		}

		innerTranslations = mapValue
	}
	return yamlPath
}
