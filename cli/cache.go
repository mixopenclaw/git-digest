package main

import (
	"encoding/json"
	"io/ioutil"
)

func saveCache(path string, v interface{}) error {
	b, _ := json.Marshal(v)
	return ioutil.WriteFile(path, b, 0644)
}

func loadCache(path string, v interface{}) error {
	b, err := ioutil.ReadFile(path)
	if err != nil { return err }
	return json.Unmarshal(b, v)
}
