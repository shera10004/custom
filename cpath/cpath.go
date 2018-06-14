package cpath

import (
	"os"
	"path/filepath"
)

func GetCurrent() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	return exPath
}
