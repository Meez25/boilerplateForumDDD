package utils

import (
	"io/fs"
	"net/http"
)

type CustomFileSystem struct {
	Fs http.FileSystem
}

func (cfs CustomFileSystem) Open(name string) (http.File, error) {
	file, err := cfs.Fs.Open(name)

	if err != nil {
		return nil, err
	}

	stat, err := file.Stat()

	if err != nil {
		return nil, err
	}

	if stat.IsDir() {
		closeErr := file.Close()
		if closeErr != nil {
			return nil, closeErr
		}
		return nil, fs.ErrNotExist
	}

	return file, nil
}
