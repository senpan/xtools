package confx

import (
	"log"
	"strings"

	"github.com/spf13/cast"
	"gopkg.in/yaml.v2"
)

// YamlFile yaml struct
// support yaml file parse
// implemented Config interface
type YamlFile struct {
	data map[string]interface{} // Section -> key : value
}

// load yaml file
func loadYamlFile(content []byte) (cfg Config, err error) {
	data := make(map[string]interface{}, 0)
	// convert bytes to map
	err = yaml.Unmarshal(content, &data)
	if err != nil {
		log.Printf("loadYamlFile error: %v", err)
		return nil, err
	}
	yamlFile := new(YamlFile)
	yamlFile.data = data
	return yamlFile, nil
}

// MustValue function implemented
func (yf *YamlFile) MustValue(section, key string, defaultVal ...string) string {
	defaultValue := ""
	if len(defaultVal) > 0 {
		defaultValue = defaultVal[0]
	}
	if val, ok := yf.data[section]; !ok {
		return defaultValue
	} else {
		// match the key
		if data, ok := val.(map[interface{}]interface{}); ok {
			for k, v := range data {
				if cast.ToString(k) == key {
					return cast.ToString(v)
				}
			}
		} else {
			return defaultValue
		}
	}
	return defaultValue
}

// MustValueArray implemented
// split value by delimiter,like ",","-"
func (yf *YamlFile) MustValueArray(section, key, delimiter string) []string {
	val := yf.MustValue(section, key, "")
	if val != "" {
		return strings.Split(val, delimiter)
	}
	return nil
}

// GetKeyList implemented
// get all keys
func (yf *YamlFile) GetKeyList(section string) []string {
	if val, err := yf.GetSection(section); err != nil {
		return nil
	} else {
		data := make([]string, len(val))
		i := 0
		for k := range val {
			data[i] = k
			i++
		}
		return data
	}
}

// GetSectionList empty function
func (yf *YamlFile) GetSectionList() []string {
	return nil
}

// GetSection implemented
func (yf *YamlFile) GetSection(section string) (map[string]string, error) {
	if val, ok := yf.data[section]; !ok {
		return nil, nil
	} else {
		// math the type
		if data, ok := val.(map[interface{}]interface{}); ok {
			ret := make(map[string]string, len(data))
			for k, v := range data {
				ret[cast.ToString(k)] = cast.ToString(v)
			}
			return ret, nil
		}
	}
	return nil, nil
}

// GetSectionToObject implemented
// object must be a pointer
func (yf *YamlFile) GetSectionToObject(section string, obj interface{}) error {
	// no section,use all data
	if section == "" {
		byt, err := yaml.Marshal(yf.data)
		if err != nil {
			return err
		}
		// format bytes to object
		err = yaml.Unmarshal(byt, obj)
		if err != nil {
			return err
		}
	} else if val, ok := yf.data[section]; !ok {
		// not hit value
		return nil
	} else {
		// convert value to object by section
		byt, err := yaml.Marshal(val)
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(byt, obj)
		if err != nil {
			return err
		}
	}
	return nil
}
