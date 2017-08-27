package lib

import (
	"io/ioutil"
	"path/filepath"
)

func ReadFile(filename string) ([]byte, error) {
	file, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadFile(file)
}
