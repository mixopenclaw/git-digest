package storage

import (
	"io/ioutil"
	"os"
)

type LocalStorage struct {
	Root string
}

func (l *LocalStorage) SaveArtifact(path string, data []byte) error {
	full := l.Root + "/" + path
	return ioutil.WriteFile(full, data, 0644)
}

func (l *LocalStorage) LoadArtifact(path string) ([]byte, error) {
	full := l.Root + "/" + path
	return ioutil.ReadFile(full)
}
