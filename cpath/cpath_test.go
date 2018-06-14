package cpath

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func Test_GetCurrnetPath(test *testing.T) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	fmt.Println("ex:", ex)
	fmt.Println("exPath:", exPath)
}
