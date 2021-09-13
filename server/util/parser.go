package util

import (
	"encoding/json"
	"io/ioutil"
)

func LoadMapping(file string) (facultyMap map[string]string, err error) {
	body, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	if err = json.Unmarshal(body, &facultyMap); err != nil {
		return
	}
	return facultyMap, nil
}
