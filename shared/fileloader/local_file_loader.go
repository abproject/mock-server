package sharedfileloader

import (
	"io/ioutil"
	"path/filepath"
)

type LocalFileLoader struct{}

func NewLocalFileLoader() FileLoader {
	return &LocalFileLoader{}
}

func (fileLoader *LocalFileLoader) Load(path string, filename string) []byte {
	fullPath, _ := filepath.Abs(path)

	file, err := ioutil.ReadFile(filepath.Join(fullPath, filename))
	if err != nil {
		panic(err)
	}
	return file
}
