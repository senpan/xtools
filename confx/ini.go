package confx

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	goini "gopkg.in/ini.v1"
)

// IniFile ini struct
// implement the Config interface
type IniFile struct {
	*goini.File
}

// load function
func loadIniFile(path string) (cfg Config, err error) {
	// load file
	content, ie := ioutil.ReadFile(path)
	if ie != nil {
		err = ie
		log.Printf("loadIniFile error: %v", ie)
		return nil, err
	}
	// customizes the config to use environment variables.
	content = []byte(os.ExpandEnv(string(content)))

	config := new(IniFile)
	file, err := goini.Load(content)
	config.File = file
	return config, err
}

// GetSectionToObject implemented
// obj must a pointer
func (ifl *IniFile) GetSectionToObject(section string, obj interface{}) error {
	return ifl.File.Section(section).MapTo(obj)
}

func (ifl *IniFile) MustValue(section, key string, defaultVal ...string) string {
	defVal := ""
	if len(defaultVal) > 0 {
		defVal = defaultVal[0]
	}
	return ifl.File.Section(section).Key(key).MustString(defVal)
}

func (ifl *IniFile) MustValueArray(section, key, delimiter string) []string {
	val := ifl.File.Section(section).Key(key).Value()
	if len(val) == 0 {
		return []string{}
	}

	vales := strings.Split(val, delimiter)
	for i := range vales {
		vales[i] = strings.TrimSpace(vales[i])
	}
	return vales
}

func (ifl *IniFile) GetKeyList(section string) []string {
	return ifl.File.Section(section).KeyStrings()
}
func (ifl *IniFile) GetSection(section string) (map[string]string, error) {
	s, e := ifl.File.GetSection(section)
	if s == nil {
		return map[string]string{}, nil
	}
	return s.KeysHash(), e
}

func (ifl *IniFile) GetSectionList() []string {
	return ifl.File.SectionStrings()
}
