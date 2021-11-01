package main

import (
	"os"
	"path/filepath"
)

func ensureDir(fileName string) error {
	dirName := filepath.Dir(fileName)
	if _, err := os.Stat(dirName); err != nil {
		err := os.MkdirAll(dirName, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func create(name string) (*os.File, error) {
	_, e := os.Stat(name)
	if e == nil {
		return nil, os.ErrExist
	}
	return os.Create(name)
}

func main() {
	path := os.Args[1]
	err := ensureDir(path)
	if err != nil {
		os.Exit(1)
	}

	f, err := create(path)
	if err != nil {
		os.Exit(1)
	}

	if os.IsExist(err) {
		os.Exit(1)
	}

	f.Close()

	os.Exit(0)
}
