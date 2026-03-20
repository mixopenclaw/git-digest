package main

import (
	"encoding/json"
	"fmt"
)

func RenderJSON(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func RenderHuman(v interface{}) string {
	return fmt.Sprintf("%v", v)
}
