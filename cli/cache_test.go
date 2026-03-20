package main

import (
	"os"
	"testing"
)

func TestCacheSaveLoad(t *testing.T) {
	path := "./.testcache.json"
	defer os.Remove(path)
	type data struct{ V string }
	if err := saveCache(path, data{"x"}); err != nil {
		t.Fatalf("saveCache failed: %v", err)
	}
	var d data
	if err := loadCache(path, &d); err != nil {
		t.Fatalf("loadCache failed: %v", err)
	}
	if d.V != "x" { t.Fatalf("unexpected value") }
}
