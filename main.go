package lib

import (
	"io/ioutil"
	"path/filepath"
	"reflect"
	"time"
)

func ReadFile(filename string) ([]byte, error) {
	file, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadFile(file)
}

func GetStructTags(item interface{}, tagName string, strict interface{}) []string {
	isStrict := true
	if strict != nil {
		isStrict = strict.(bool)
	}
	var tags []string
	st := reflect.TypeOf(item)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		fieldTag := field.Tag.Get(tagName)
		if fieldTag != "" {
			tags = append(tags, fieldTag)
		} else if isStrict == false {
			tags = append(tags, field.Name)
		}
	}
	return tags
}

func GetISOTimemillisecons() string {
	return time.Now().Format("2006-01-02T15:03:04.999")
}
func GetISOTimeSeconds() string {
	return time.Now().Format("2006-01-02T15:03:04")
}
