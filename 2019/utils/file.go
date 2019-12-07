package utils

import (
	"bufio"
	"os"
)

func openFile(path string) *os.File {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return f
}
func NewFileReader(path string) *bufio.Reader {
	return bufio.NewReader(openFile(path))
}
